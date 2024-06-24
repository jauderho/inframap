FROM golang:1.20.2-alpine3.16@sha256:0848e987c1be480a253637c82e4e08f3c6589b644804320d14e7a2321326f98f as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk -q --no-progress add git make; \
	make build

FROM alpine:3.20.1@sha256:b89d9c93e9ed3597455c90a0b88a8bbb5cb7188438f70953fede212a0c4394e0
COPY --from=builder /app/inframap /app/
ENTRYPOINT ["/app/inframap"]
