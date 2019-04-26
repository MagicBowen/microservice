package main

import api "github.com/magicbowen/microservice/examples/services/api"

type userEntity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func createUserEntity(user *api.UserInfoMsg) *userEntity {
	return &userEntity{ID: int(user.GetId()), Name: user.GetName()}
}
