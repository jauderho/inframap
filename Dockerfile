FROM golang:1.19.3-alpine3.16@sha256:3607071679bd7702e5461ff72ad0886760a03cb00a09163be3020d8f1cda5299 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk -q --no-progress add git make; \
	make build

FROM alpine:3.16.3@sha256:b95359c2505145f16c6aa384f9cc74eeff78eb36d308ca4fd902eeeb0a0b161b
COPY --from=builder /app/inframap /app/
ENTRYPOINT ["/app/inframap"]
