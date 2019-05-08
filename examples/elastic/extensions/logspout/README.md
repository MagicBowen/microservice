# Logspout extension

Logspout collects all Docker logs using the Docker logs API, and forwards them to Logstash without any additional configuration.

## Usage

If you want to include the Logspout extension, run Docker Compose from the root of the repository with an additional
command line argument referencing the `logspout-compose.yml` file:

```bash
$ docker-compose -f docker-compose.yml -f extensions/logspout/logspout-compose.yml up
```

In your Logstash pipeline configuration, enable the `udp` input and set the input codec to `json`:

```
input {
  udp {
    port  => 5000
    codec => json
  }
}
```

## Documentation

https://github.com/looplab/logspout-logstash

## debug

```sh
docker run -d --name="logspout" --volume=/var/run/docker.sock:/var/run/docker.sock -p 8070:80 --net=examples_microservices gliderlabs/logspout

curl http://127.0.0.1:8070/logs
```

## docker compose

```yml
  logstash:
    build:
      context: ./elastic/logstash/
      args:
        ELK_VERSION: $ELK_VERSION
    volumes:
      - ./elastic/logstash/config/logstash.yml:/usr/share/logstash/config/logstash.yml:ro
      - ./elastic/logstash/pipeline:/usr/share/logstash/pipeline:ro
    ports:
      - "5000:5000/udp"
      - "5000:5000/tcp"
    environment:
      LS_JAVA_OPTS: "-Xmx256m -Xms256m"
      LOGSPOUT: ignore
    networks:
      - microservices
    depends_on:
      - elasticsearch

  logspout:
    image: bekt/logspout-logstash
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
    environment:
      ROUTE_URIS: logstash://logstash:5000
      DEBUG: 1
      LOGSPOUT: ignore
    ports:
      - "8090:80"
    networks:
      - microservices
    depends_on:
      - logstash
```
