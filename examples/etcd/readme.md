## How to setup ETCD cluster

According `docker-compose.yml`

## add traefik dynamic configurations to etcd
```bash
docker exec -it -e ETCDCTL_API=3 etcd1 etcdctl get --prefix /traefik
docker exec -it -e ETCDCTL_API=3 etcd1 etcdctl put /traefik/frontends/service/entrypoints/0 web
docker exec -it -e ETCDCTL_API=3 etcd1 etcdctl put /traefik/frontends/service/backend service
docker exec -it -e ETCDCTL_API=3 etcd1 etcdctl put /traefik/frontends/service/routes/test_1/rule Path:/
docker exec -it -e ETCDCTL_API=3 etcd1 etcdctl put /traefik/backends/service/servers/server1/url http://172.19.0.2:8866
```

## etcdkeeper

support V2&V3

https://github.com/evildecay/etcdkeeper

## etcd-viewer

just support V2

https://github.com/nikfoundas/etcd-viewer

```yml
  etcd-viewer:
    image: nikfoundas/etcd-viewer
    container_name: etcd-viewer
    ports:
      - "8090:8080"
    depends_on:
      - etcd1      
    networks:
      - microservices
```