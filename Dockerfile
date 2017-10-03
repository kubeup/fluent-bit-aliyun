FROM golang:1.7.3 as builder
WORKDIR /go/src/github.com/yanshaocong/fluent-bit-aliyun/
COPY ./ .
RUN go build -buildmode=c-shared -o out_sls.so .

FROM fluent/fluent-bit:0.11
USER root
COPY docker-image/etc /fluent-bit/etc/
COPY --from=builder /go/src/github.com/yanshaocong/fluent-bit-aliyun/out_sls.so /fluent-bit/
CMD ["/fluent-bit/bin/fluent-bit", "-c", "/fluent-bit/etc/fluent-bit.conf", "-e", "/fluent-bit/out_sls.so"]
