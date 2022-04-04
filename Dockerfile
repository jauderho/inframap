FROM golang:1.18.0-alpine3.15@sha256:fcb74726937b96b4cc5dc489dad1f528922ba55604d37ceb01c98333bcca014f as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk -q --no-progress add git make; \
	make build

FROM alpine:3.15.3@sha256:f22945d45ee2eb4dd463ed5a431d9f04fcd80ca768bb1acf898d91ce51f7bf04
COPY --from=builder /app/inframap /app/
ENTRYPOINT ["/app/inframap"]
