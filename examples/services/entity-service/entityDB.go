package main

type entityDB interface {
	findOne(filter interface{}, value interface{}) error
	insertOne(value interface{}) error
	updateOne(filter interface{}, update interface{}) error
	delete(filter interface{}) error
}
