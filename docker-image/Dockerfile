FROM fluent/fluent-bit:0.11
USER root
COPY etc /fluent-bit/etc/
COPY out_sls.so /fluent-bit/
CMD ["/fluent-bit/bin/fluent-bit", "-c", "/fluent-bit/etc/fluent-bit.conf", "-e", "/fluent-bit/out_sls.so"]
