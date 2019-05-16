## prometheus exporters

### redis

```sh
docker pull oliver006/redis_exporter
docker run -d --name redis_exporter -p 9121:9121 oliver006/redis_exporter
```

```yml
scrape_configs:

...

- job_name: redis_exporter
  static_configs:
  - targets: ['localhost:9121']

...
```

github: https://github.com/oliver006/redis_exporter

### mongoDB

```sh
docker pull ssalaues/mongodb-exporter
docker run --rm ssalaues/mongodb-exporter -h
-mongodb.uri
```

github: https://github.com/dcu/mongodb_exporter

### elasticsearch

```sh
docker pull justwatch/elasticsearch_exporter:1.0.2
docker run --rm -p 9114:9114 justwatch/elasticsearch_exporter:1.0.2
# es.uri
```

```yml
elasticsearch_exporter:
    image: justwatch/elasticsearch_exporter:1.0.2
    command:
     - '-es.uri=http://elasticsearch:9200'
    restart: always
    ports:
    - "127.0.0.1:9114:9114"
```