FROM golang:1.12.0 AS builder
ARG moduleName=github.com/xieqiaoyu/nats-streaming-cli
COPY . /gomod
WORKDIR /gomod
RUN CGO_ENABLED=0 GO111MODULE=on GOOS=linux   GOARCH=amd64 go build -a -ldflags "-s -w -X '${moduleName}/metadata.Version=`git describe --tags|xargs`' -X '${moduleName}/metadata.OS=linux' -X '${moduleName}/metadata.ARCH=amd64'" -o pkg/nats-streaming-cli .


FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /gomod/pkg/nats-streaming-cli /
ENTRYPOINT ["/nats-streaming-cli"]