## grafana

### docker

```sh
docker run -d --name=grafana -p 3000:3000 grafana/grafana
```

```sh
$ docker run \
  -d \
  -p 3000:3000 \
  --name=grafana \
  -e "GF_SERVER_ROOT_URL=http://grafana.server.name" \
  -e "GF_SECURITY_ADMIN_PASSWORD=secret" \
  grafana/grafana
```

### reference
- https://grafana.com/
- https://grafana.com/dashboards
- https://grafana.com/plugins
- https://grafana.com/docs/installation/docker/