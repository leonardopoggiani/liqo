FROM golang:1.19 AS builder

ENV GO111MODULE=on

WORKDIR /go/src/liqo.io/

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o discovery .

FROM alpine:3.15

COPY --from=builder /go/src/liqo.io/discovery .

ENTRYPOINT ["./discovery"]

