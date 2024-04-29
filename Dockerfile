FROM golang:1.20.2-alpine3.16@sha256:0848e987c1be480a253637c82e4e08f3c6589b644804320d14e7a2321326f98f as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk -q --no-progress add git make; \
	make build

FROM alpine:3.19.1@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b
COPY --from=builder /app/inframap /app/
ENTRYPOINT ["/app/inframap"]
