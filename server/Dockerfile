FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go build

FROM alpine

COPY --from=builder /build/server /server

EXPOSE 80

ENTRYPOINT [ "/server" ]