FROM golang:1.18.3-alpine3.15@sha256:f9181168749690bddb6751b004e976bf5d427425e0cfb50522e92c06f761def7 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk -q --no-progress add git make; \
	make build

FROM alpine:3.16.0@sha256:686d8c9dfa6f3ccfc8230bc3178d23f84eeaf7e457f36f271ab1acc53015037c
COPY --from=builder /app/inframap /app/
ENTRYPOINT ["/app/inframap"]
