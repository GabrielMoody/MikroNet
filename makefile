run:
	docker compose -f docker-compose.development.yaml up -d

down:
	docker compose -f docker-compose.development.yaml down

build:
	docker compose -f docker-compose.development.yaml build