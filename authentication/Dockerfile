FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY views /app/views
COPY static /app/static
RUN go build -o /go/bin/app ./cmd/

FROM alpine
WORKDIR /usr/bin
COPY --from=build /go/bin .
COPY --from=build /app/views /usr/bin/views
COPY --from=build /app/static /usr/bin/static
EXPOSE 8050
CMD ["app"]