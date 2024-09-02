FROM golang:alpine
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
COPY config/auth-config.env /app/profiles/config/
EXPOSE 8011
RUN go mod tidy
RUN go build -o binary ./cmd/main.go
ENTRYPOINT ["/app/binary"]