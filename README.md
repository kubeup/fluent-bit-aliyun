# Fluent Bit Aliyun

[![CircleCI](https://circleci.com/gh/kubeup/fluent-bit-aliyun/tree/master.svg?style=shield)](https://circleci.com/gh/kubeup/fluent-bit-aliyun)

Fluent Bit output plugin which sends logs to [Aliyun Simple Log Service(SLS)](https://cn.aliyun.com/product/sls).
You can do realtime analysis and archiving using other Aliyun products. Check Aliyun
documents for detailed usage.

## Installing

Currently we only provide Docker image for Linux machines at [Docker Hub](https://hub.docker.com/r/kubeup/fluent-bit-aliyun/).
If you're not using Docker, you can compile the plugin from source.

### Docker

Docker supports fluentd as a log driver. You can launch `fluent-bit-aliyun` in a container
and ask Docker to send logs to it.

```
$ docker run -d --network host -e ALIYUN_ACCESS_KEY=YOUR_ACCESS_KEY -e ALIYUN_ACCESS_KEY_SECRET=YOUR_ACCESS_KEY_SECRET -e ALIYUN_SLS_PROJECT=YOUR_PROJECT -e ALIYUN_SLS_LOGSTORE=YOUR_LOGSTORE -e ALIYUN_SLS_ENDPOINT=cn-hangzhou.log.aliyuncs.com kubeup/fluent-bit-aliyun:master /fluent-bit/bin/fluent-bit -c /fluent-bit/etc/fluent-bit-forwarder.conf -e /fluent-bit/out_sls.so
$ docker run --log-driver=fluentd -d nginx
```

You can also set fluentd as your default log driver by adding `--log-driver=fluentd`
to Docker daemon.

### Kubernetes

We use `DaemonSet` to deploy `fluent-bit-aliyun` on all nodes which will collect Docker
logs from all pods on each machine. The `kubernetes` filter will add tags from pod metadata.

First we create a `Secret` containing all the configurations.

```
$ kubectl create secret generic fluent-bit-config --namespace=kube-system --from-literal=ALIYUN_ACCESS_KEY=YOUR_ACCESS_KEY --from-literal=ALIYUN_ACCESS_KEY_SECRET=YOUR_ACCESS_KEY_SECRET --from-literal=ALIYUN_SLS_PROJECT=YOUR_PROJECT --from-literal=ALIYUN_SLS_LOGSTORE=YOUR_LOGSTORE --from-literal=ALIYUN_SLS_ENDPOINT=cn-hangzhou.log.aliyuncs.com
```

Then we deploy the `DaemonSet`.

```
$ kubectl create -f https://raw.githubusercontent.com/kubeup/fluent-bit-aliyun/master/fluent-bit-daemonset.yaml
```

We could check the status.

```
$ kubectl get pods --namespace=kube-system
```

## Configurations

You can config the plugin using environment variables. We support these configs:

- ALIYUN_ACCESS_KEY
- ALIYUN_ACCESS_KEY_SECRET 
- ALIYUN_SLS_PROJECT
- ALIYUN_SLS_LOGSTORE
- ALIYUN_SLS_ENDPOINT

You should consider creating a dedicated account in [RAM](https://cn.aliyun.com/product/ram)
just for logging for better security protection.
