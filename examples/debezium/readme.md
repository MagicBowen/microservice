## debezium

### Start the topology as defined in http://debezium.io/docs/tutorial/

```sh
export DEBEZIUM_VERSION=0.8
docker-compose -f docker-compose-mongodb.yaml up
```

### Initialize MongoDB replica set and insert some test data

```sh
docker-compose -f docker-compose-mongodb.yaml exec mongodb bash -c '/usr/local/bin/init-inventory.sh'
```

### Start MongoDB connector

```sh
curl -i -X POST -H "Accept:application/json" -H  "Content-Type:application/json" http://localhost:8083/connectors/ -d @register-mongodb.json
```

### Consume messages from a Debezium topic

```sh
docker-compose -f docker-compose-mongodb.yaml exec kafka /kafka/bin/kafka-console-consumer.sh \
    --bootstrap-server kafka:9092 \
    --from-beginning \
    --property print.key=true \
    --topic dbserver1.inventory.customers
```

In microservice case:

```sh
docker exec examples_kafka_1 /kafka/bin/kafka-console-consumer.sh \
    --bootstrap-server kafka:9092 \
    --from-beginning \
    --property print.key=true \
    --topic mongocollector.microservice-example.user
```

### use kafka watcher to listen topic

```sh
docker run -it --name watcher --rm --net=examples_microservices --link examples_zookeeper_1:zookeeper --link examples_kafka_1:kafka -e ZOOKEEPER_CONNECT=zookeeper:2181 -e KAFKA_BROKER=kafka:9092 debezium/kafka:0.9 watch-topic -a -k mongocollector.microservice-example.user
```

### Modify records in the database via MongoDB client

```sh
docker-compose -f docker-compose-mongodb.yaml exec mongodb bash -c 'mongo -u $MONGODB_USER -p $MONGODB_PASSWORD --authenticationDatabase admin inventory'

db.customers.insert([
    { _id : 1005, first_name : 'Bob', last_name : 'Hopper', email : 'thebob@example.com' }
]);
```

### Shut down the cluster

```sh
docker-compose -f docker-compose-mongodb.yaml down
```

### reference

- https://debezium.io/docs/tutorial/
- https://github.com/debezium/debezium-examples