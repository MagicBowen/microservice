## jaeger

### docker

Run jaeger in storage of memory:

```sh
docker run -d -p6831:6831/udp -p16686:16686 --name jaeger jaegertracing/all-in-one:latest
```

### example

```sh
docker run \
  --rm \
  --link jaeger \
  --env JAEGER_AGENT_HOST=jaeger \
  --env JAEGER_AGENT_PORT=6831 \
  -p8080-8083:8080-8083 \
  jaegertracing/example-hotrod:latest \
  all
```

### reference
- https://www.jaegertracing.io/docs/1.11/getting-started/
- https://medium.com/opentracing/take-opentracing-for-a-hotrod-ride-f6e3141f7941
- https://github.com/jaegertracing/jaeger/tree/master/examples/hotrod