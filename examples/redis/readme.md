## how to run redis by docker

```sh
docker run -p 6379:6379 -v $PWD/data:/data -d redis redis-server --appendonly yes
```