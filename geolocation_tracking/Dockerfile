FROM golang:alpine
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
#COPY config/auth-config.env /app/geolocation_tracking/config/
EXPOSE 8016
RUN go mod tidy
RUN go build -o binary ./cmd/main.go
ENTRYPOINT ["/app/binary"]