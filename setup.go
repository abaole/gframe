package gframe

import (
	"context"

	"github.com/abaole/gframe/conf"
	"github.com/abaole/gframe/db"
	"github.com/abaole/gframe/redis"
)

type featureEnabledOptions struct {
	Db     bool `json:"db"`
	Tracer bool `json:"tracer"`
	Redis  bool `json:"redis"`
}

const (
	glibConfigEnablesKey = "supports"
	glibConfigDb         = "db"
	glibConfigTracer     = "tracer"
	glibConfigRedis      = "redis"
)

var (
	ctx, stop         = context.WithCancel(context.Background())
	defEnabledOptions = &featureEnabledOptions{false, false, false}
)

func Init(appId, path string) error {

	var err error

	conf.SetAppID(appId)

	viper := InitConfig(path)
	if err != nil {
		return err
	}

	if err = viper.Sub(glibConfigEnablesKey).Unmarshal(defEnabledOptions); err != nil {
		return release(err)
	}

	// init database
	if defEnabledOptions.Db {
		dbConfig := db.Options{}
		if err = viper.Sub(glibConfigDb).Unmarshal(&dbConfig); err != nil {
			return release(err)
		}

		initDBManger(&dbConfig)
	}

	if defEnabledOptions.Redis {
		rdsConfig := redis.Options{}
		if err = viper.Sub(glibConfigRedis).Unmarshal(&rdsConfig); err != nil {
			return release(err)
		}
		if err = runRedisManger(&rdsConfig); err != nil {
			return release(err)
		}
	}

	// init tracer
	if defEnabledOptions.Tracer {
		tCfg := tracerConfig{}
		if err = viper.Sub(glibConfigTracer).Unmarshal(&tCfg); err != nil {
			return release(err)
		}
		if err = InitTracing(tCfg); err != nil {
			return release(err)
		}
	}

	return nil
}

func release(err error) error {
	stop()
	closeDb()
	closeTracer()
	closeRedis()
	return err
}

// Destroy - 释放资源
func Destroy() error {
	return release(nil)
}
