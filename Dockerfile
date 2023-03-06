FROM golang:1.20.1-alpine3.16@sha256:bf83658eca96639b6a27a30b6b43800b632f0994fd9a63212b0fbf9c8e056f5f as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk -q --no-progress add git make; \
	make build

FROM alpine:3.17.2@sha256:69665d02cb32192e52e07644d76bc6f25abeb5410edc1c7a81a10ba3f0efb90a
COPY --from=builder /app/inframap /app/
ENTRYPOINT ["/app/inframap"]
