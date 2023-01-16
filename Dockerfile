FROM golang:1.19.5-alpine3.16@sha256:65060885c8882f0119d77edcb42b414a0b96f6a8d466e9bbd782311ae51bf9f1 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk -q --no-progress add git make; \
	make build

FROM alpine:3.17.0@sha256:8914eb54f968791faf6a8638949e480fef81e697984fba772b3976835194c6d4
COPY --from=builder /app/inframap /app/
ENTRYPOINT ["/app/inframap"]
