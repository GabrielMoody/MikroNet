run:
	docker compose -f docker-compose.development.yaml up -d

build:
	docker compose -f docker-compose.development.yaml build --no-cache