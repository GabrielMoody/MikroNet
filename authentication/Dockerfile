FROM golang:alpine
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
COPY go.mod .
COPY go.sum .
EXPOSE 8010
RUN go mod tidy
RUN go build -o binary ./cmd/main.go
ENTRYPOINT ["/app/binary"]
