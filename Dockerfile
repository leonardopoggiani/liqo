FROM golang:1.19 as builder
WORKDIR /tmp/builder

COPY go.mod go.sum ./
RUN go mod download

ARG COMPONENT=virtual-kubelet
RUN test -n "$COMPONENT" || ( echo "The COMPONENT argument is unset. Aborting" && false )

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$(go env GOARCH) go build -ldflags="-s -w" ./cmd/$COMPONENT

FROM alpine:3.15

RUN apk update && \
    apk add --no-cache ca-certificates && \
    update-ca-certificates && \
    echo "@testing https://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories && \
    apk update && \
    apk add criu-dev@testing && \
    apk add containerd@testing && \
    rm -rf /var/cache/apk/*

ARG COMPONENT=virtual-kubelet
COPY --from=builder /tmp/builder/$COMPONENT /usr/bin/$COMPONENT
RUN ln -s /usr/bin/$COMPONENT /usr/bin/liqo-component

ENTRYPOINT [ "/usr/bin/liqo-component" ]