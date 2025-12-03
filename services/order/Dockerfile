FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make ca-certificates tzdata
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /go/bin/app ./cmd/
RUN cp /usr/share/zoneinfo/Asia/Singapore /etc/localtime && \
  echo "Asia/Singapore" > /etc/timezone

FROM alpine
WORKDIR /usr/bin
RUN apk add --no-cache tzdata
ENV TZ=Asia/Singapore
COPY --from=build /go/bin .
EXPOSE 8060
CMD ["app"]