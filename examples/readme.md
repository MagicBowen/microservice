## TODO

- [x]: http service implementation
- [x]: register http service to traefik
- [x]: setup etcd cluster
- [x]: traefik fetch http service from etcd | file | docker
- [x]: http serfice use gRPC to fetch entity
- [x]: entity service implementation, using mongo
- [x]: entity service using redis
- [x]: put entity service and http service in docker
- [x]: register and discovery of mongo and redis from etcd
- [x]: register and discovery of entity service from etcd
- [x]: LB to entity service
- [ ]: modify api to graphiz
- [ ]: concurrent node compete to the master using etcd
- [ ]: logstash log to kafka (spike fluentD)
- [ ]: search log from elasticsearch
- [ ]: metrics in promotheus
- [ ]: distributed trace
- [ ]: CD pipeline
- [ ]: data analyzing using kafka
- [ ]: some job using FaaS: https://cloud.tencent.com/developer/article/1365541
- [ ]: move docker to K8S

## network

docker-compose.yml文件中声明了网络`microservices`，最后实际会创建网络`examples_microservices`。
可以使用`docker network ls`查看。

docker deamon针对用户自定义网络，会在docker embedded DNS里面为其中的主机增加DNS解析项。

任意进入一个该网络中的主机，可以看到主机的DNS服务器指向内嵌DNS服务器

```sh
docker exec -it ubuntu sh
cat etc/resolv.conf
```

输入如下：
```
nameserver 127.0.0.11
options ndots:0
```

其中`127.0.0.11`就是内嵌DNS服务器的地址。关于内嵌DNS的详细介绍见[官文](https://docs.docker.com/v17.09/engine/userguide/networking/configure-dns/)

默认情况下，容器内的DNS服务器配置会采用宿主机的配置，关于默认DNS配置参见[官文](https://docs.docker.com/v17.09/engine/userguide/networking/default_network/configure-dns/)。

以下情况下会使用docker deamon的内嵌DNS服务器：
- 容器有一个name
- 容器使用了网络别名net-alias
- 容器使用link
这时容器内的`etc/resolv.conf`则指向内嵌DNS服务器`127.0.0.11`。该DNS服务器会做容器域名到IP地址的DNS解析。

例如我们在自建网络内的容器内执行`dig example_http-service_1 a`就能获得DNS的解析结果，看到对应container name对应的IP地址。

安装`dig`： `apt-get install dnsutils`

如果用docker-compose启动的service，使用`dig service_name a`，也能看到该service下的所有主机的IP地址。由此可知，docker-compose在scale service的时候，会把该service下的每个container的IP加入到内嵌DNS的service name下。

经过测试，内嵌DNS中一个service name对应了多个container的IP的时候，做DNS解析请求获得service name的时候，DNS服务器总是返回最后一个有效的IP地址。

## service registration & discovery

reference:
- https://medium.com/@maxy_ermayank/service-registration-and-discovery-configuration-management-dffb15fc08a7
- https://deaddesk.top/service-discovery-with-etcd/

## Load Balance

- Proxy Model
- Balancing-aware Client
- External Load Balancing Service

reference:
- https://grpc.io/blog/loadbalancing/
- https://github.com/grpc/grpc/blob/master/doc/load-balancing.md
- https://segmentfault.com/a/1190000008672912
- https://www.cnblogs.com/SmartLee/p/5161415.html
- https://github.com/liyue201/grpc-lb