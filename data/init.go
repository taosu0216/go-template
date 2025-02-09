package data

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-template/data/conf"
	"go-template/data/db/ent"
	"go-template/logger"
	"time"

	_ "github.com/lib/pq"
)

type ToolsCtx struct {
	Cfg conf.Config `json:"cfg"`

	DB    *ent.Client   `json:"db"`
	Cache *redis.Client `json:"cache"`

	Logger *logger.ZapLogger
}

var Tools *ToolsCtx

func (t *ToolsCtx) LogDuration(name string, start time.Time) {
	duration := time.Since(start).Milliseconds()
	str := fmt.Sprintf("func: [%s] | duration=%dms", name, duration)
	t.Logger.LogDurationInData(str)
}

func Init(logger *logger.ZapLogger) {
	cfg := conf.InitCfg()

	db := initDB(cfg, logger)
	//logger.InitInfof("db client")

	cache := initCache(cfg, logger)
	logger.InitInfof("cache client")

	Tools = &ToolsCtx{
		Cfg:    cfg,
		DB:     db,
		Cache:  cache,
		Logger: logger,
	}
}

func initCache(cfgs conf.Config, logger *logger.ZapLogger) *redis.Client {

	opts := &redis.Options{
		Addr:         cfgs.CacheCfg.Addr,
		Username:     cfgs.CacheCfg.User,
		Password:     cfgs.CacheCfg.Password,
		WriteTimeout: time.Duration(cfgs.CacheCfg.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfgs.CacheCfg.ReadTimeout) * time.Second,
		DialTimeout:  time.Duration(5) * time.Second,
	}
	rdb := redis.NewClient(opts)
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logger.InitErrf("[链接cache 失败] err: %v", err)
		panic(err)
	}
	return rdb
}

func initDB(cfgs conf.Config, logger *logger.ZapLogger) *ent.Client {
	cli, err := ent.Open("postgres", cfgs.DBCfg.Source)
	if err != nil {
		logger.InitErrf("[链接db 失败] err: %v", err)
		panic(err)
	}
	return cli
}
