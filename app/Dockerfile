FROM golang:alpine AS builder

RUN apk add --no-cache build-base

COPY . /tmp/source

WORKDIR /tmp/source

RUN go mod tidy
RUN go mod download

RUN go build -o x-tech cmd/main.go

RUN mkdir /app
RUN cp /tmp/source/x-tech /app/x-tech
RUN rm -rf /tmp/*

FROM alpine

COPY --from=builder /app/x-tech /app/x-tech

ENTRYPOINT ["/app/x-tech"]
