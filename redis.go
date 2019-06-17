package framework

import (
	"framework/redis"
)

var CRedis *redis.Cacher

func Redis() *redis.Cacher {
	return CRedis
}

func runRedisManger(opts *redis.Options) error {
	c, err := redis.New(
		redis.Options{
			Addr:     opts.Addr,
			Password: opts.Password,
			Prefix:   "pay_",
		})
	if err != nil {
		return err
	}
	CRedis = c
	return nil
}

func closeRedis() error {
	if CRedis != nil {
		return CRedis.Close()
	}
	return nil
}
