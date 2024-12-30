FROM golang:alpine
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
#COPY config/auth-config.env /app/business_owner/config/
EXPOSE 8013
RUN go mod tidy
RUN go build -o binary ./cmd/main.go
ENTRYPOINT ["/app/binary"]