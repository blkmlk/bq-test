all: test run

test:
	docker-compose -p bq-test up -d
	sleep 5
	go test -mod=vendor -race -p 1 -count=1 ./internal/... || (sudo docker-compose -p bq-test down; exit 1)

run:
	docker-compose -p bq-test up -d
	sleep 5
	go run -mod=vendor cmd/bq/main.go || (sudo docker-compose -p bq-test down; exit 1)

.PHONY: all