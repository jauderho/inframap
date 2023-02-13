FROM golang:1.20.0-alpine3.16@sha256:846aec1428bbc484681625bc066c057f2fc5805eb9ea67bd47e4f86e4ca35d83 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk -q --no-progress add git make; \
	make build

FROM alpine:3.17.1@sha256:f271e74b17ced29b915d351685fd4644785c6d1559dd1f2d4189a5e851ef753a
COPY --from=builder /app/inframap /app/
ENTRYPOINT ["/app/inframap"]
