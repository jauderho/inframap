FROM golang:1.20.2-alpine3.16@sha256:0848e987c1be480a253637c82e4e08f3c6589b644804320d14e7a2321326f98f as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk -q --no-progress add git make; \
	make build

FROM alpine:3.18.0@sha256:02bb6f428431fbc2809c5d1b41eab5a68350194fb508869a33cb1af4444c9b11
COPY --from=builder /app/inframap /app/
ENTRYPOINT ["/app/inframap"]
