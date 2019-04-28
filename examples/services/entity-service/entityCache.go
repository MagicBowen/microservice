package main

import (
	"errors"
)

type entityCache interface {
	set(key string, value string) error
	get(key string) (string, error)
}

type noneCache struct{}

func (nc *noneCache) set(key string, value string) error {
	return nil
}

func (nc *noneCache) get(key string) (string, error) {
	return "", errors.New("None cache for storing")
}
