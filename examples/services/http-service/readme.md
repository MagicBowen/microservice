## Test

```sh
curl localhost:8866/users/1
curl -l -H "Content-type: application/json" -X POST -d '{"id" : 1, "name" : "bowen"}' localhost:8866/users
curl -l -H "Content-type: application/json" -X PUT -d '{"name" : "wangbo"}' localhost:8866/users/1
curl -X DELETE localhost:8866/users/1
```

```sh
curl --header "jaeger-baggage:microservice-trace-id=1" 127.0.0.1/api/users/1
curl 127.0.0.1/api/users/1
curl -l -H "Content-type: application/json" -X POST -d '{"id" : 1, "name" : "bowen"}' 127.0.0.1/api/users
curl -l -H "Content-type: application/json" -X PUT -d '{"name" : "wangbo"}' 127.0.0.1/api/users/1
curl -X DELETE 127.0.0.1/api/users/1
```