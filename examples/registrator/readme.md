## registrator

Do not support ETCD V3!

### github

https://github.com/gliderlabs/registrator

### setup

```
docker pull gliderlabs/registrator:latest
```

```sh
docker run -d \
    --name=registrator \
    --net=host \
    --volume=/var/run/docker.sock:/tmp/docker.sock \
    gliderlabs/registrator:latest \
      consul://localhost:8500
```

```yml
  registrator:
    image: gliderlabs/registrator
    container_name: registrator
    command: etcd://etcd1:2379
    volumes:
      - /var/run/docker.sock:/tmp/docker.sock
    depends_on:
      - etcd1      
    networks: 
      - microservices  
```