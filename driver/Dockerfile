FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /go/bin/app ./cmd/

FROM alpine
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8020
CMD ["app"]