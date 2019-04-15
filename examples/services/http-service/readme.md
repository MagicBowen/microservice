## Test

```sh
curl localhost:8866/users/1
curl -l -H "Content-type: application/json" -X POST -d '{"id" : 1, "name" : "bowen"}' localhost:8866/users
curl -l -H "Content-type: application/json" -X PUT -d '{"name" : "wangbo"}' localhost:8866/users/1
curl -X DELETE localhost:8866/users/1
```