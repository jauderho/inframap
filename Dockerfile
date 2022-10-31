FROM golang:1.19.2-alpine3.16@sha256:e4dcdac3ed37d8c2b3b8bcef2909573b2ad9c2ab53ba53c608909e8b89ccee36 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk -q --no-progress add git make; \
	make build

FROM alpine:3.16.2@sha256:bc41182d7ef5ffc53a40b044e725193bc10142a1243f395ee852a8d9730fc2ad
COPY --from=builder /app/inframap /app/
ENTRYPOINT ["/app/inframap"]
