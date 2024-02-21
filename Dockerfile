FROM golang:1.21.7 AS build

WORKDIR /build
COPY . .
RUN go build -o http-server

FROM alpine:latest

WORKDIR /http-server
COPY --from=build /build/http-server .
RUN apk add libc6-compat

CMD ["./http-server"]


