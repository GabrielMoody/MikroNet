# Ride Hailing Microservices Documentation

MikroNet is a backend systems for real-time driver tracking, order management and notification services to server mobile applications both for the users and the drivers.

## Feature
- Real-time drivers location
- Geospatial query with PostGis for accurate point and redis for faster access
- User order to nearest driver available
- Push notification
- Event-driver using RabbitMQ
- JWT Authentication
- API Gateway using Kong

## Tech Stack
- Golang
- PostgreSQL + PostGIS
- Redis
- RabbitMQ
- Docker & Docker Compose
- Kong Gateway
- Promotheus + Grafana + Loki + FluentBit