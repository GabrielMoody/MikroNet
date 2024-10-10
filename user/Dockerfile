FROM golang:alpine
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
EXPOSE 8014
RUN go mod tidy
RUN go build -o binary ./cmd/main.go
ENTRYPOINT ["/app/binary"]