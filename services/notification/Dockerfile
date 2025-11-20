FROM golang:alpine
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
#COPY config/auth-config.env /app/notification/config/
EXPOSE 8015
RUN go mod tidy
RUN go build -o binary ./cmd/main.go
ENTRYPOINT ["/app/binary"]