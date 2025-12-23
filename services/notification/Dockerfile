FROM golang:alpine AS build
RUN apk update && apk add --no-cache git
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /go/bin/app ./cmd/main.go

FROM alpine
WORKDIR /usr/bin
COPY --from=build /go/bin/app .
EXPOSE 9000
CMD ["app"]