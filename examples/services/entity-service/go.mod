module entity-service

require (
	github.com/MagicBowen/microservice/examples/services/utils/registration v0.0.0-20190512112104-06dfc5b32c5d
	github.com/MagicBowen/microservice/examples/services/utils/tracing v0.0.0-20190512112104-06dfc5b32c5d
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-cmp v0.3.0 // indirect
	github.com/magicbowen/microservice v0.0.0-20190512112104-06dfc5b32c5d
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

replace golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f

replace golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be => github.com/golang/oauth2 v0.0.0-20180821212333-d2e6202438be

replace golang.org/x/net v0.0.0-20190311183353-d8887717615a => github.com/golang/net v0.0.0-20190311183353-d8887717615a

replace golang.org/x/net v0.0.0-20180906233101-161cd47e91fd => github.com/golang/net v0.0.0-20180906233101-161cd47e91fd

replace golang.org/x/net v0.0.0-20190424112056-4829fb13d2c6 => github.com/golang/net v0.0.0-20190424112056-4829fb13d2c6

replace golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3 => github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3

replace golang.org/x/tools v0.0.0-20190311212946-11955173bddd => github.com/golang/tools v0.0.0-20190311212946-11955173bddd

replace golang.org/x/tools v0.0.0-20180221164845-07fd8470d635 => github.com/golang/tools v0.0.0-20180221164845-07fd8470d635

replace golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0

replace golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 => github.com/golang/crypto v0.0.0-20190308221718-c2843e01d9a2
