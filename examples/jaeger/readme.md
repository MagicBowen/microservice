## jaeger

### quick start

#### docker

Run jaeger in storage of memory:

```sh
docker run -d -p6831:6831/udp -p16686:16686 --name jaeger jaegertracing/all-in-one:latest
```

#### example

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

### manual setup

#### Jaeger Collector

```sh
docker run \
  -e SPAN_STORAGE_TYPE=elasticsearch \
  jaegertracing/jaeger-collector:1.11 \
  --help
```

```sh
docker run \
  -e SPAN_STORAGE_TYPE=elasticsearch \
  -e ES_SERVER_URLS=<...> \
  jaegertracing/jaeger-collector:1.11

```

#### Jaeger UI

```sh
docker run -d --rm \
  -p 16686:16686 \
  -p 16687:16687 \
  -e SPAN_STORAGE_TYPE=elasticsearch \
  -e ES_SERVER_URLS=http://<ES_SERVER_IP>:<ES_SERVER_PORT> \
  jaegertracing/jaeger-query:1.11
```

#### Jaeger Agent

```sh
docker run \
  --rm \
  -p5775:5775/udp \
  -p6831:6831/udp \
  -p6832:6832/udp \
  -p5778:5778/tcp \
  jaegertracing/jaeger-agent:1.11 \
  --reporter.grpc.host-port=jaeger-collector.jaeger-infra.svc:14250
```

#### Jaeger Dependence DAG

```sh
docker run --env STORAGE=elasticsearch --env ES_NODES=http://elasticsearch:9200 --env ES_USERNAME=elastic --env ES_PASSWORD=changeme jaegertracing/spark-dependencies
```

### reference
- https://www.jaegertracing.io/docs/1.11/getting-started/
- https://medium.com/opentracing/take-opentracing-for-a-hotrod-ride-f6e3141f7941
- https://github.com/jaegertracing/jaeger/tree/master/examples/hotrod