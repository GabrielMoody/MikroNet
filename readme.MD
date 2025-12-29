<p align="center">
  <img src="https://github.com/GabrielMoody/MikroNet/actions/workflows/go_test.yml/badge.svg" alt="CI" />
  <img src="https://img.shields.io/github/license/GabrielMoody/MikroNet" alt="License" />
</p>

# Ride Hailing Microservices

MikroNet is a backend systems for real-time driver tracking, order management and notification services to server mobile applications both for the users and the drivers.

## Feature
- Real-time drivers location
- Geospatial query with PostGis for accurate point and redis for faster access
- User order to nearest driver available
- Push notification
- Event-driven using RabbitMQ
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

## How to Run

### Prerequisite
- Docker
- Docker Compose v2.1+

On linux/git bash:
```bash
make
```