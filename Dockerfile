FROM golang:1.17-alpine as builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./cmd ./cmd
COPY pkg ./pkg

RUN CGO_ENABLED=1 go build -o bin/port-service ./cmd/portservice/main.go

FROM alpine:latest
MAINTAINER Artur Obrzut <artek.obrzut@gmail.com>

ENV URL ""
ENV PORT 50051
WORKDIR /app
COPY --from=builder /app/bin/port-service ./port-service

RUN addgroup -S portserv \
&& adduser -D -S  -s /sbin/nologin -G portserv -u 10001 portserv \
&& chown -R portserv:portserv /app/port-service

USER 10001
EXPOSE 50051
CMD [ "/app/port-service" ]