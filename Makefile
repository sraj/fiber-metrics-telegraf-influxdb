BIN=bin/fiber-metrics-telegraf-influxdb

.PHONY: build run up down dc-build logs test clean

build:
	go build -ldflags="-s -w" -o $(BIN) .

run:
	air -c .air.toml

up:
	docker-compose up -d

down:
	docker-compose down

dc-build:
	docker-compose up -d --build

logs:
	docker-compose logs -f

test:
	go test ./... -v

clean:
	rm -rf bin
	docker-compose down -v --remove-orphans
