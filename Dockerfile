FROM golang:1.19.3-alpine3.16@sha256:d171aa333fb386089206252503bc6ab545072670e0286e3d1bbc644362825c6e as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk -q --no-progress add git make; \
	make build

FROM alpine:3.17.0@sha256:8914eb54f968791faf6a8638949e480fef81e697984fba772b3976835194c6d4
COPY --from=builder /app/inframap /app/
ENTRYPOINT ["/app/inframap"]
