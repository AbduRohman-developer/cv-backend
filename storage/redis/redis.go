package redis

import (
	r "github.com/go-redis/redis/v8"
)

type repository struct {
	Client r.Client
}

func NewClient() {}
