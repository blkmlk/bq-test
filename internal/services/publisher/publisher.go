package publisher

import (
	"bq/internal/models"
	"bq/internal/services/config"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/sarulabs/di/v2"
)

type publisher struct {
	config      config.Config
	server      *http.Server
	upgrader    websocket.Upgrader
	connections map[string]*connection
}

type connection struct {
	conn *websocket.Conn
	ch   chan []byte
}

func newConnection(c *websocket.Conn) *connection {
	return &connection{
		conn: c,
		ch:   make(chan []byte, 100),
	}
}

func newPublisher(ctn di.Container) Publisher {
	return &publisher{
		config:      ctn.Get(config.DefinitionName).(config.Config),
		connections: make(map[string]*connection),
	}
}

func (p *publisher) init() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c, err := p.upgrader.Upgrade(writer, request, nil)

		if err != nil {
			return
		}

		conn := newConnection(c)
		p.connections[c.RemoteAddr().String()] = conn

		for {
			select {
			case b := <-conn.ch:
				err = conn.conn.WriteMessage(websocket.TextMessage, b)
				if err != nil {
					return
				}
			}
		}
	})

	p.server = &http.Server{
		Addr:    p.config.Host,
		Handler: mux,
	}

	go func() {
		fmt.Printf("Running publisher %s...\n", p.config.Host)
		p.server.ListenAndServe()
	}()

	return nil
}

func (p *publisher) PublishRecord(ctx context.Context, record *models.Record) error {
	data, err := json.Marshal(record)
	if err != nil {
		return err
	}

	for _, c := range p.connections {
		select {
		case c.ch <- data:
		default:
		}
	}
	return nil
}
