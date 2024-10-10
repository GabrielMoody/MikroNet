FROM golang:alpine
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
EXPOSE 8013
RUN go mod download
RUN go build -o binary ./cmd/main.go
ENTRYPOINT ["/app/binary"]