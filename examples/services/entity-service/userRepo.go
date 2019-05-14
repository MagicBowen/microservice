package main

import (
	"context"
	"log"

	"gopkg.in/mgo.v2/bson"
)

type userRepo struct {
	db    entityDB
	cache *entityCache
}

func (repo *userRepo) createUser(ctx context.Context, user *userEntity) error {
	if err := repo.db.insertOne(ctx, bson.M{"id": user.ID, "name": user.Name}); err != nil {
		return err
	}
	return repo.cache.set(ctx, user.ID, user)
}

func (repo *userRepo) getUserByID(ctx context.Context, id int) *userEntity {
	var user userEntity
	if err := repo.cache.get(ctx, id, &user); err != nil {
		if err := repo.db.findOne(ctx, bson.M{"id": id}, &user); err != nil {
			log.Printf("find user by id(%d) failed: %v", id, err)
			return nil
		}
	}
	return &user
}

func (repo *userRepo) updateUser(ctx context.Context, user *userEntity) error {
	if err := repo.db.updateOne(ctx, bson.M{"id": user.ID}, bson.M{"$set": bson.M{"name": user.Name}}); err != nil {
		return nil
	}
	return repo.cache.set(ctx, user.ID, user)
}

func (repo *userRepo) deleteUser(ctx context.Context, id int) error {
	if err := repo.db.delete(ctx, bson.M{"id": id}); err != nil {
		return err
	}
	return repo.cache.del(ctx, id)
}

func createUserRepo(db entityDB, cache cache) *userRepo {
	return &userRepo{db: db, cache: createEntityCache("entity/user", cache)}
}
