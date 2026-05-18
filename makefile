run:
	docker compose up -d

no-cache:
	docker compose build --no-cache && docker compose up -d --force-recreate

down:
	docker compose down

build:
	docker compose build