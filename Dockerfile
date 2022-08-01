FROM golang:1.18.4-alpine3.16@sha256:9b2a2799435823dbf6fef077a8ff7ce83ad26b382ddabf22febfc333926400e4 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk -q --no-progress add git make; \
	make build

FROM alpine:3.16.1@sha256:7580ece7963bfa863801466c0a488f11c86f85d9988051a9f9c68cb27f6b7872
COPY --from=builder /app/inframap /app/
ENTRYPOINT ["/app/inframap"]
