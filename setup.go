package gframe

import (
	"context"

	"github.com/abaole/gframe/db"
	"github.com/abaole/gframe/redis"
	"github.com/spf13/viper"
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
	Vp                *viper.Viper
	ctx, stop         = context.WithCancel(context.Background())
	defEnabledOptions = &featureEnabledOptions{false, false, false}
)

func Init(appId, path string) error {

	var err error

	Vp = InitConfig(path)
	if err != nil {
		return err
	}

	if err = Vp.Sub(glibConfigEnablesKey).Unmarshal(defEnabledOptions); err != nil {
		return release(err)
	}

	// init database
	if defEnabledOptions.Db {
		dbConfig := db.Options{}
		if err = Vp.Sub(glibConfigDb).Unmarshal(&dbConfig); err != nil {
			return release(err)
		}

		initDBManger(&dbConfig)
	}

	if defEnabledOptions.Redis {
		rdsConfig := redis.Options{}
		if err = Vp.Sub(glibConfigRedis).Unmarshal(&rdsConfig); err != nil {
			return release(err)
		}
		if err = runRedisManger(&rdsConfig); err != nil {
			return release(err)
		}
	}

	// init tracer
	if defEnabledOptions.Tracer {
		tCfg := TracerConfig{}
		v := Vp.Sub(glibConfigTracer)
		if err = v.Unmarshal(&tCfg); err != nil {
			return release(err)
		}
		if err = InitTracing(&tCfg); err != nil {
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
