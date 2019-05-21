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
- [x]: logstash log to ELK (spike fluentD)
- [x]: distributed trace
- [x]: dependence of service by health
- [x]: metrics in prometheus
- [x]: route to kibana, jaeger, prometheus by traefik
- [ ]: entity service use mongo to publish entity event
- [ ]: event services compete to deal msgs in kafka
- [ ]: Use graphQL
- [ ]: CD pipeline
- [ ]: data analyzing by spark
- [ ]: stream by flink
- [ ]: some job by FaaS: https://cloud.tencent.com/developer/article/1365541 （spike Knative）
- [ ]: move docker to K8S
- [ ]: API by swagger
- [ ]: Contract Test
- [ ]: Mock Server

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

ETCD vs Redis：
redis首先追求的是性能和可用性，满足AP。ETCD满足CP，有租约和watch机制，数据有版本号，更适合服务注册以及分布式配置中心。

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

syslog：
syslog是网络上各种设备将日志收集到日志服务器的一种数据协议，它几乎被所有的网络设备支持，并且能够记录多种事件类型的日志消息，支持syslog的设备常见的有路由器、交换机、打印机等等，甚至unix-like的服务器也可以支持产生syslog消息，用以记录用户的登录、防火墙事件、apache或者nginx access日志等。
目前有两个syslog RFC标准RFC 3164、RFC 5424，两者的主要区别在于数据的格式不同.Syslog包分为3个部分，PRI, HEADER,以及MSG，总长度不能超过1024个字节.
主流的syslog工具有四种syslogd[RFC 3164]、sysklogd[RFC 5424]、syslog-ng、rsyslog，前两种的是系统组件，后两种是最近兴起的开源工具，syslog-ng、rsyslog可以兼容syslogd、sysklogd发送过来的数据，支持各种syslog协议数据的格式化、过滤等处理，比如可以定义数据的时间格式、按照syslog数据中任意部分过滤数据以及重定向输出到远端等。
syslog设备依据两个重要的文件：/etc/syslogd守护进程和/etc/syslog.conf配置文件。通常情况下，多数syslog信息被写到/var/adm或/var/log目录下的信息文件中(messages.*)。一个典型的syslog记录包括生成程序的名字和一个文本信息。它还包括一个设备和一个优先级范围。

config a syslog server : https://robinle.github.io/docker/2017/01/05/Docker-rsyslog-logstash/

## distributed trace

concept：
- Trace: The description of a transaction as it moves through a distributed system. (TraceID)
- Span: A named, timed operation representing a piece of the workflow. Spans accept key:value tags as well as fine-grained, timestamped, structured logs attached to the particular span instance.
	- An operation name
	- A start timestamp and finish timestamp
	- A set of key:value span Tags
	- A set of key:value span Logs
	A SpanContext

	- References between Spans
		- ChildOf：In a ChildOf reference, the parent Span depends on the child Span in some capacity. 
		- FollowsFrom：Some parent Spans do not depend in any way on the result of their child Spans. 
- SpanContext: The span context contains the trace identifier, span identifier, and any other data that the tracing system needs to propagate to the downstream service
	- An implementation-dependent state to refer to the distinct span within a trace,i.e., the implementing Tracer’s definition of spanID and traceID
	- Any Baggage Items
		- These are key:value pairs that cross process-boundaries.
		- These may be useful to have some data available for access throughout the trace.

examples:

```
Example Span:
    t=0            operation name: db_query               t=x 

     +-----------------------------------------------------+
     | · · · · · · · · · ·    Span     · · · · · · · · · · |
     +-----------------------------------------------------+

Tags:
- db.instance:"jdbc:mysql://127.0.0.1:3306/customers
- db.statement: "SELECT * FROM mytable WHERE foo='bar';"

Logs:
- message:"Can't connect to mysql server on '127.0.0.1'(10061)"

SpanContext:
- trace_id:"abc123"
- span_id:"xyz789"
- Baggage Items:
  - special_id:"vsid1738"
```

speciatin:
- OpenTrace: https://opentracing.io/specification/
- specification: https://github.com/opentracing/specification/
- 规范中文版本：https://github.com/opentracing-contrib/opentracing-specification-zh
- https://github.com/opentracing/specification/blob/master/semantic_conventions.md

tools:
- Google: Dapper
- Twitter: Zipkin , OpenZipkin(scala, special RPC)
- Ubber: Jaeger
- Alibaba: EagleEye
- Elastic APM : https://www.elastic.co/guide/en/apm/agent/index.html 
- gRPC tracing: https://github.com/grpc-ecosystem/grpc-opentracing
- http tracing: https://github.com/opentracing-contrib/go-stdlib

reference：
- Dapper: https://storage.googleapis.com/pub-tools-public-publication-data/pdf/36356.pdf
- Jaeger: https://www.jaegertracing.io/docs/1.11/
- EagleEys: https://zhuanlan.zhihu.com/p/31324360
- diff: https://riboseyim.github.io/2018/05/18/DevOps-OpenTracing/
- plugins: https://github.com/opentracing-contrib/
- tutorial: - https://github.com/yurishkuro/opentracing-tutorial
- https://www.jaegertracing.io/docs/1.11/getting-started/
- https://medium.com/opentracing/take-opentracing-for-a-hotrod-ride-f6e3141f7941
- https://github.com/jaegertracing/jaeger/tree/master/examples/hotrod

TODO:
- [x]: Launch Jaeger Agent
- [x]: Launch Jaeger Collector with ES
- [x]: Launch Jaeger UI
- [x]: Trace Traefik
- [x]: Trace Http Service
- [x]: Trace gRPC
- [x]: Trace MongoDB seperate
- [x]: Trace Redis seperate

## metrics

reference:
- Skywaling: https://github.com/apache/skywalking

## MQ

消息顺序
- 严格消息顺序，需要保证消息发往同一个消费者，且获得上一条消息的ACK后再发送后一条消息。
消息去重
- 消费端处理消息的业务逻辑保持幂等性
- 保证每条消息都有唯一编号且保证消息处理成功与去重表的日志同时出现（消费者维护一份成功处理消息ID的日志）
- 去重表需要有一个去重窗口
事务消息
- 将事务分为本地事务和远端事务
- 本地事务和消息发送在一起，保证事务和消息的发送的仍有事务关系
- 由消息队列保证消息不丢失，通知远端。远端处理超时或者失败，继续让远端重试处理
- 远端处理失败一定次数后，上报告警人工干预

kafka的Exactly once：
- producer开启幂等性`enable.idempotence = true`，`ack=-1`，这样可以保证发送到每个分区的消息的exactly once
- consumer需要配置手动提交offset，需要在业务处理完后持久化offset并且同步提交offset。主要业务处理后业务数据的持久化和offset的持久化以及同步提交offset需要在一个事务里面。对业务数据和offset可以事务存储在一个支持ACID的存储里。


reference:
- 将数据库变更自动写到kafka：https://debezium.io/
- https://github.com/debezium/debezium-examples
- dockers of debezium: https://hub.docker.com/u/debezium
- kafka事务：http://www.jasongj.com/kafka/transaction/

## DB

### mysql

半同步复制
binary log
relay log

### mongoDB

oplog:
单节点mongoDB使能oplog，启动配置文件或者命令行使能“--replSet rs0”，然后开始Mongo shell

```sh
use local
rs.initiate()
db.oplog.rs.find()
```

读写分离架构演进
读写分离后的数据一致性问题以及解决手段
mysql的读写分离和Mongo的集群的区别：核心在于mongoDB内置选举以及sharding的能力，sharding可以让数据容量随着机器数的增多水平扩展。mongoDB的集群一般是副本集和sharding结合。

reference:
- https://mp.weixin.qq.com/s/7vUrsY0SDvBTZQewwNkCiA

## cache

cache和DB的数据一致性方案
- 读操作先从缓存读，缓存没有的话，去DB里拿，顺便写入缓存
- 数据库更改，一般先删除缓存（大多数缓存的复杂对象修改代价大），然后写数据库。写数据库成功后不更新缓存，等待读操作自行更新缓存
- 为了避免数据从缓存中删除后，在更改数据库未成功的时候，又有并发读取该数据引起的缓存更新不一致的问题。每个写DB的操作先建立一个写DB缓存，写成功后删除。新的并发进来发现缓存没有数据，还需要检查写DB缓存，正在写DB为成功的数据需要客户反复读。
- 写数据库的操作为了避免出现缓存不一致，还可以有一个binlog的异步消费者，针对任何DB修改，等待一段时间t后，再次删除缓存一次。这样可以再做一次保证，任何不一致最多只能存在t时间。这可以解决读操作从备用数据库读到还未来的及同步过去的数据造成的不一致的问题。

redis主从复制和集群区别很明显，前者是为了提供冗余高可用，后者是为了将数据分散保存，不同的redis节点有不同的数据。

reference:
- https://segmentfault.com/a/1190000015804406


## distributed lock


## circuit breaker

reference:
- https://github.com/afex/hystrix-go


## global ID service

demand:
- 全局唯一
- 趋势递增：由于多数RDBMS使用B-tree的数据结构来存储索引数据，在主键的选择上面我们应该尽量使用有序的主键保证写入性能。
- 单调递增：事务版本号、IM增量消息、排序等特殊需求
- 信息安全：与上一个冲突，在一些应用场景下，会需要ID无规则、不规则。避免泄露MAC地址或者ID规律。
- 高可用：

algorithm:
- UUID
- snowflake
	- 64bit,对ID进行分段：timestamp + workerId + seqID 
- 借助数据库
	- 利用数据生成自增索引
	- 多数据库分库分表，解决高可用问题
	- 每个写库设置不同的 auto_increment 初始值，以及相同的增长步长
	- 借助redis的单进程特点，使用redis生成
- Leaf
	- Leaf-segment （可以利用双buffer优化）
	- Leaf-snowflake

reference:
- https://tech.meituan.com/2017/04/21/mt-leaf.html
- https://gavinlee1.github.io/2017/06/28/%E5%B8%B8%E8%A7%81%E5%88%86%E5%B8%83%E5%BC%8F%E5%85%A8%E5%B1%80%E5%94%AF%E4%B8%80ID%E7%94%9F%E6%88%90%E7%AD%96%E7%95%A5%E5%8F%8A%E7%AE%97%E6%B3%95%E7%9A%84%E5%AF%B9%E6%AF%94/