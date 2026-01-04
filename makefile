.PHONY: up down build logs run

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

