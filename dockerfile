FROM golang:1.15

ENV GO111MODULE=on\
    GOPROXY=https://goproxy.cn,direct\
    CGO_ENABLED=0\
    GOOS=linux\
    GOARCH=amd64

WORKDIR /worktools_srvs

# VOLUME /opt/log

COPY  . /worktools_srvs


STOPSIGNAL SIGINT

EXPOSE 50051

ENTRYPOINT ["./worktools_srvs"]