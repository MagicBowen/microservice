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
- [ ]: Use graphQL
- [x]: logstash log to ELK (spike fluentD)
- [ ]: distributed trace
- [ ]: metrics in promotheus
- [ ]: entity service use mongo to publish entity event
- [ ]: event services compete to deal msgs in kafka
- [ ]: CD pipeline
- [ ]: data analyzing using spark
- [ ]: some job using FaaS: https://cloud.tencent.com/developer/article/1365541
- [ ]: move docker to K8S
- [ ]: API by swagger
- [ ]: Contract test
- [ ]: mock server

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

registration：
- self registration
- by others: eg. registrator

discovery：
- server side
- client side
	- in client
	- out client

reference:
- https://medium.com/@maxy_ermayank/service-registration-and-discovery-configuration-management-dffb15fc08a7
- https://deaddesk.top/service-discovery-with-etcd/

## Load Balance

- Proxy Model
- Balancing-aware Client
- External Load Balancing Service

### LB of gRPC 

```sh
	r, err := d.Resolver(targetServiceName)
	if err != nil {
		log.Fatalf("Discovery initial resolver for gRPC failed: %v", err)
	}
	b := grpc.RoundRobin(r)

	client.cc, err = grpc.Dial("", grpc.WithInsecure(), grpc.WithBalancer(b))
```

```sh
// Deprecated: please use package balancer/roundrobin.
func RoundRobin(r naming.Resolver) Balancer {
	return &roundRobin{r: r}
}
```

```sh
// Deprecated: please use package resolver.
type Operation uint8

const (
	// Add indicates a new address is added.
	Add Operation = iota
	// Delete indicates an existing address is deleted.
	Delete
)

// Update defines a name resolution update. Notice that it is not valid having both
// empty string Addr and nil Metadata in an Update.
//
// Deprecated: please use package resolver.
type Update struct {
	// Op indicates the operation of the update.
	Op Operation
	// Addr is the updated address. It is empty string if there is no address update.
	Addr string
	// Metadata is the updated metadata. It is nil if there is no metadata update.
	// Metadata is not required for a custom naming implementation.
	Metadata interface{}
}
//
// Resolver creates a Watcher for a target to track its resolution changes.
//
// Deprecated: please use package resolver.
type Resolver interface {
	// Resolve creates a Watcher for target.
	Resolve(target string) (Watcher, error)
}

// Watcher watches for the updates on the specified target.
//
// Deprecated: please use package resolver.
type Watcher interface {
	// Next blocks until an update or error happens. It may return one or more
	// updates. The first call should get the full set of the results. It should
	// return an error if and only if Watcher cannot recover.
	Next() ([]*Update, error)
	// Close closes the Watcher.
	Close()
}
```

```sh
type Balancer interface {
	// Start does the initialization work to bootstrap a Balancer. For example,
	// this function may start the name resolution and watch the updates. It will
	// be called when dialing.
	Start(target string, config BalancerConfig) error
	// Up informs the Balancer that gRPC has a connection to the server at
	// addr. It returns down which is called once the connection to addr gets
	// lost or closed.
	// TODO: It is not clear how to construct and take advantage of the meaningful error
	// parameter for down. Need realistic demands to guide.
	Up(addr Address) (down func(error))
	// Get gets the address of a server for the RPC corresponding to ctx.
	// i) If it returns a connected address, gRPC internals issues the RPC on the
	// connection to this address;
	// ii) If it returns an address on which the connection is under construction
	// (initiated by Notify(...)) but not connected, gRPC internals
	//  * fails RPC if the RPC is fail-fast and connection is in the TransientFailure or
	//  Shutdown state;
	//  or
	//  * issues RPC on the connection otherwise.
	// iii) If it returns an address on which the connection does not exist, gRPC
	// internals treats it as an error and will fail the corresponding RPC.
	//
	// Therefore, the following is the recommended rule when writing a custom Balancer.
	// If opts.BlockingWait is true, it should return a connected address or
	// block if there is no connected address. It should respect the timeout or
	// cancellation of ctx when blocking. If opts.BlockingWait is false (for fail-fast
	// RPCs), it should return an address it has notified via Notify(...) immediately
	// instead of blocking.
	//
	// The function returns put which is called once the rpc has completed or failed.
	// put can collect and report RPC stats to a remote load balancer.
	//
	// This function should only return the errors Balancer cannot recover by itself.
	// gRPC internals will fail the RPC if an error is returned.
	Get(ctx context.Context, opts BalancerGetOptions) (addr Address, put func(), err error)
	// Notify returns a channel that is used by gRPC internals to watch the addresses
	// gRPC needs to connect. The addresses might be from a name resolver or remote
	// load balancer. gRPC internals will compare it with the existing connected
	// addresses. If the address Balancer notified is not in the existing connected
	// addresses, gRPC starts to connect the address. If an address in the existing
	// connected addresses is not in the notification list, the corresponding connection
	// is shutdown gracefully. Otherwise, there are no operations to take. Note that
	// the Address slice must be the full list of the Addresses which should be connected.
	// It is NOT delta.
	Notify() <-chan []Address
	// Close shuts down the balancer.
	Close() error
}
```

```sh
type roundRobin struct {
	r      naming.Resolver
	w      naming.Watcher
	addrs  []*addrInfo // all the addresses the client should potentially connect
	mu     sync.Mutex
	addrCh chan []Address // the channel to notify gRPC internals the list of addresses the client should connect to.
	next   int            // index of the next address to return for Get()
	waitCh chan struct{}  // the channel to block when there is no connected address available
	done   bool           // The Balancer is closed.
}
```

gRPC通过和balance交互获得可用链接地址。首先gRPC在dial的时候调用balance.Start，然后使用balance.Notify()获得一个地址更新的chan。每次有地址更新则通过chan获得新的地址全集，和内部持有的已经链接的地址进行对比。新增的地址调用balance.Up将连接状态修改为connected。不再存在的地址调用之前获得down方法修改地址状态为非连接状态。每次gRPC请求使用的具体地址调用balance.Get进行获取。如果balance.Get没有可用连接态的地址，则根据参数BalancerGetOptions决定是否阻塞Get函数。如果阻塞的话，当下次有新地址Up的时候会采用一个chan (waitCh)通知Get取消阻塞。

reference:
- https://grpc.io/blog/loadbalancing/
- https://github.com/grpc/grpc/blob/master/doc/load-balancing.md
- https://segmentfault.com/a/1190000008672912
- https://www.cnblogs.com/SmartLee/p/5161415.html
- https://github.com/liyue201/grpc-lb

## GraphQL

- 标准化了协议的定义
- 按需返回查询数据
- 基于强类型，改善了协作方式和工具（生成文档、编辑器、schema校验、、、）

reference:
- https://graphql.org/
- https://graphql.github.io/graphql-spec/
- https://graphql.org/code/
- https://www.infoq.cn/article/LVQGuC3vQX-T3PpVCkHt

## logs

- elk
- logstash to elasticsearch
- fluentd or logpout to kafka, logstash receive from kafka
- kibana/ grafana for dashboard

## distributed trace


## metrics


## MQ


## distributed lock