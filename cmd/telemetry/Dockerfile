FROM golang:1.19 as builder
WORKDIR /home/poggio/Documenti/GitHub/liqo/cmd/

COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN  go mod download

ARG COMPONENT=telemetry
RUN test -n "$COMPONENT" || ( echo "The COMPONENT argument is unset. Aborting" && false )

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$(go env GOARCH) go build -ldflags="-s -w" .


FROM alpine:3.15

RUN apk update && \
    apk add --no-cache ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

ARG COMPONENT=telemetry
COPY --from=builder /home/poggio/Documenti/GitHub/liqo/cmd/$COMPONENT /usr/bin/$COMPONENT
RUN ln -s /usr/bin/$COMPONENT /usr/bin/liqo-component

ENTRYPOINT [ "/usr/bin/liqo-component" ]
