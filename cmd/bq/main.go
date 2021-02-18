package main

import (
	"bq/internal/services/config"
	"bq/internal/services/listener"
	"bq/internal/services/publisher"
	"bq/internal/services/storage"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sarulabs/di/v2"
)

func main() {
	path := flag.String("settings", "settings.yaml", "Settings file path")
	flag.Parse()

	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal(err)
	}

	err = builder.Add(
		di.Def{
			Name: config.DefinitionNameConfigPath,
			Build: func(ctn di.Container) (interface{}, error) {
				return *path, nil
			},
		},
		config.Definition,
		listener.Definition,
		storage.Definition,
		publisher.Definition,
	)

	if err != nil {
		log.Fatal(err)
	}

	ctn := builder.Build()

	l := ctn.Get(listener.DefinitionName).(listener.Listener)

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = l.Start(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Received signal", <-ch)
	cancel()
	wg.Wait()
}
