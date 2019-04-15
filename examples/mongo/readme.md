## How to run mongodb by docker

```sh
docker run -p 27017:27017 -v $PWD/data/db:/data/db -d mongo
```