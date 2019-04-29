package main

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

type userRepo struct {
	db    entityDB
	cache *entityCache
}

func (repo *userRepo) createUser(user *userEntity) error {
	if err := repo.db.insertOne(bson.M{"id": user.ID, "name": user.Name}); err != nil {
		return err
	}
	return repo.cache.set(user.ID, user)
}

func (repo *userRepo) getUserByID(id int) *userEntity {
	var user userEntity
	if err := repo.cache.get(id, &user); err != nil {
		if err := repo.db.findOne(bson.M{"id": id}, &user); err != nil {
			log.Printf("find user by id(%d) failed: %v", id, err)
			return nil
		}
	}
	return &user
}

func (repo *userRepo) updateUser(user *userEntity) error {
	if err := repo.db.updateOne(bson.M{"id": user.ID}, bson.M{"$set": bson.M{"name": user.Name}}); err != nil {
		return nil
	}
	return repo.cache.set(user.ID, user)
}

func (repo *userRepo) deleteUser(id int) error {
	if err := repo.db.delete(bson.M{"id": id}); err != nil {
		return err
	}
	return repo.cache.del(id)
}

func createUserRepo(db entityDB, cache cache) *userRepo {
	return &userRepo{db: db, cache: createEntityCache("entity/user", cache)}
}
