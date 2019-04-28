package main

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

type userRepo struct {
	db    entityDB
	cache entityCache
}

func (repo *userRepo) createUser(user *userEntity) error {
	return repo.db.insertOne(bson.M{"id": user.ID, "name": user.Name})
}

func (repo *userRepo) getUserByID(id int) *userEntity {
	var user userEntity
	if err := repo.db.findOne(bson.M{"id": id}, &user); err != nil {
		log.Printf("find user by id(%d) failed: %v", id, err)
		return nil
	}
	return &user
}

func (repo *userRepo) updateUser(user *userEntity) error {
	return repo.db.updateOne(bson.M{"id": user.ID}, bson.M{"$set": bson.M{"name": user.Name}})
}

func (repo *userRepo) deleteUser(id int) error {
	return repo.db.delete(bson.M{"id": id})
}

func createUserRepo(db entityDB, cache entityCache) *userRepo {
	return &userRepo{db: db, cache: cache}
}
