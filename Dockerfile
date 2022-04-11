FROM golang:1.18.0-alpine3.15@sha256:a2ca4f4c0828b1b426a3153b068bf32a21868911c57a9fc4dccdc5fbb6553b35 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk -q --no-progress add git make; \
	make build

FROM alpine:3.15.3@sha256:f22945d45ee2eb4dd463ed5a431d9f04fcd80ca768bb1acf898d91ce51f7bf04
COPY --from=builder /app/inframap /app/
ENTRYPOINT ["/app/inframap"]
