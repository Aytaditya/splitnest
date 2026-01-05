.PHONY: up down build logs run help

help:
	@echo "Makefile commands:"
	@echo "  up     - Build and start the Docker containers"
	@echo "  down   - Stop and remove the Docker containers"
	@echo "  build  - Build the Docker images"
	@echo "  logs   - Follow the logs of the Docker containers"
	@echo "  run    - Start the Docker containers"

up:
	docker compose up --build

down:
	docker compose down

build:
	docker compose build

logs:
	docker compose logs -f

run: 
	docker compose up 

