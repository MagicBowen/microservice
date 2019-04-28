## How to run mongodb by docker

```sh
docker run -p 27017:27017 -v $PWD/data/db:/data/db -d mongo
```

## mongo shell

tutorial : https://docs.mongodb.com/manual/reference/mongo-shell/

```sh
docker exec -it mongo mongo
```

```sh
# display dbs
db

# switch to db
use <database>

# Print a list of all collections for current database
show collections

# Drops or removes completely the collection
db.mycollection.drop()

# CRUD
db.myCollection.find()
db.myCollection.insertOne( { x: 1 } )
db.mycollection.insertMany()
db.mycollection.updateOne()
db.mycollection.updateMany()
db.mycollection.deleteOne()
db.mycollection.deleteMany()
db.mycollection.createIndex()
```

### sdk

- golang : https://github.com/mongodb/mongo-go-driver