module entity-service

require (
	github.com/MagicBowen/microservice/examples/services/utils/registration v0.0.0-20190514145601-46a1a0c1f9af
	github.com/MagicBowen/microservice/examples/services/utils/tracing v0.0.0-20190514145601-46a1a0c1f9af
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-cmp v0.3.0 // indirect
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/magicbowen/microservice v0.0.0-20190514145601-46a1a0c1f9af
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/tidwall/pretty v0.0.0-20190325153808-1166b9ac2b65 // indirect
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.0.1
	google.golang.org/grpc v1.20.1
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.20.0

replace google.golang.org/appengine v1.1.0 => github.com/golang/appengine v1.1.0

replace cloud.google.com/go v0.26.0 => github.com/googleapis/google-cloud-go v0.26.0

replace google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8 => github.com/google/go-genproto v0.0.0-20180817151627-c66870c02cf8

replace golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a => github.com/golang/sys v0.0.0-20190215142949-d0b11bdaac8a

replace golang.org/x/sys v0.0.0-20180909124046-d0be0721c37e => github.com/golang/sys v0.0.0-20180909124046-d0be0721c37e

replace golang.org/x/sys v0.0.0-20181107165924-66b7b1311ac8 => github.com/golang/sys v0.0.0-20181107165924-66b7b1311ac8

replace golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33 => github.com/golang/sys v0.0.0-20180905080454-ebe1bf3edb33

replace golang.org/x/sys v0.0.0-20190412213103-97732733099d => github.com/golang/sys v0.0.0-20190412213103-97732733099d

replace golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223 => github.com/golang/sys v0.0.0-20190222072716-a9d3bda3a223

replace golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f

replace golang.org/x/sync v0.0.0-20181108010431-42b317875d0f => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f

replace golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be => github.com/golang/oauth2 v0.0.0-20180821212333-d2e6202438be

replace golang.org/x/net v0.0.0-20190311183353-d8887717615a => github.com/golang/net v0.0.0-20190311183353-d8887717615a

replace golang.org/x/net v0.0.0-20180906233101-161cd47e91fd => github.com/golang/net v0.0.0-20180906233101-161cd47e91fd

replace golang.org/x/net v0.0.0-20190424112056-4829fb13d2c6 => github.com/golang/net v0.0.0-20190424112056-4829fb13d2c6

replace golang.org/x/net v0.0.0-20181220203305-927f97764cc3 => github.com/golang/net v0.0.0-20181220203305-927f97764cc3

replace golang.org/x/net v0.0.0-20181201002055-351d144fa1fc => github.com/golang/net v0.0.0-20181201002055-351d144fa1fc

replace golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3 => github.com/golang/net v0.0.0-20190404232315-eb5bcb51f2a3

replace golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3 => github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3

replace golang.org/x/tools v0.0.0-20190311212946-11955173bddd => github.com/golang/tools v0.0.0-20190311212946-11955173bddd

replace golang.org/x/tools v0.0.0-20180221164845-07fd8470d635 => github.com/golang/tools v0.0.0-20180221164845-07fd8470d635

replace golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0

replace golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 => github.com/golang/crypto v0.0.0-20190308221718-c2843e01d9a2

replace golang.org/x/crypto v0.0.0-20190426145343-a29dc8fdc734 => github.com/golang/crypto v0.0.0-20190426145343-a29dc8fdc734

replace go.etcd.io/etcd v3.3.12+incompatible => github.com/etcd-io/etcd v3.3.12+incompatible

replace go.etcd.io/bbolt v1.3.2 => github.com/etcd-io/bbolt v1.3.2

replace go.uber.org/zap v1.10.0 => github.com/uber-go/zap v1.10.0

replace go.uber.org/multierr v1.1.0 => github.com/uber-go/multierr v1.1.0

replace gopkg.in/yaml.v2 v2.2.1 => github.com/go-yaml/yaml v0.0.0-20180328195020-5420a8b6744d

replace gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce => github.com/go-mgo/mgo v0.0.0-20180705113604-9856a29383ce

replace go.uber.org/atomic v1.4.0 => github.com/uber-go/atomic v1.3.2
