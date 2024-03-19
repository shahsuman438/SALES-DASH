COMPOSE_FILE=docker-compose.yml

.PHONY: all clean build up down restart logs

all: build up

clean: 
	docker-compose -f $(COMPOSE_FILE) down
build: 
	docker-compose -f $(COMPOSE_FILE) build
up: 
	docker-compose -f $(COMPOSE_FILE) up -d
down: 
	docker-compose -f $(COMPOSE_FILE) down
restart: 
	docker-compose -f $(COMPOSE_FILE) restart
logs: 
	docker-compose -f $(COMPOSE_FILE) logs -f
remove: 
	docker-compose -f $(COMPOSE_FILE) down --rmi all -v 
server-ip:
	docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' leapflow-server
restart-all:
	make down && make build && make up && docker ps
