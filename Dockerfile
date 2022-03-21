FROM golang:1.18.0-alpine3.15@sha256:fcb74726937b96b4cc5dc489dad1f528922ba55604d37ceb01c98333bcca014f as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk -q --no-progress add git make; \
	make build

FROM alpine:3.15.0@sha256:21a3deaa0d32a8057914f36584b5288d2e5ecc984380bc0118285c70fa8c9300
COPY --from=builder /app/inframap /app/
ENTRYPOINT ["/app/inframap"]
