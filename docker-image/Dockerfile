FROM kubeup/fluent-bit:latest
USER root
COPY fluent-bit.conf /fluent-bit/etc/
COPY parsers.conf /fluent-bit/etc/
COPY out_sls.so /fluent-bit/
CMD ["/fluent-bit/bin/fluent-bit", "-c", "/fluent-bit/etc/fluent-bit.conf", "-e", "/fluent-bit/out_sls.so"]
