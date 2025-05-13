FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o /go/bin/app ./cmd/

FROM alpine
RUN apk add --no-cache tzdata
ENV TZ=Asia/Singapore
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8040
CMD ["app"]