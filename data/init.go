package data

import (
	"context"
	"fmt"
	"go-template/data/conf"
	"go-template/data/cons"
	"go-template/data/db/ent"
	"go-template/logger"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"

	_ "github.com/lib/pq"
)

type ToolsCtx struct {
	Cfg conf.Config `json:"cfg"`

	DB    *ent.Client   `json:"db"`
	Cache *redis.Client `json:"cache"`

	Logger *logger.ZapLogger

	KafkaWritersMaps map[string]*kafka.Writer
	KafkaReadersMaps map[string]*kafka.Reader
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

	kafkaWriters := initKafkaWriter(cfg, logger)
	kafkaReaders := initKafkaReader(cfg, logger)

	logger.InitInfof("kafka client")

	Tools = &ToolsCtx{
		Cfg:    cfg,
		DB:     db,
		Cache:  cache,
		Logger: logger,

		KafkaWritersMaps: kafkaWriters,
		KafkaReadersMaps: kafkaReaders,
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

func initKafkaWriter(cfgs conf.Config, logger *logger.ZapLogger) map[string]*kafka.Writer {
	writers := make(map[string]*kafka.Writer)
	for _, topic := range cons.KafkaTopics {
		writer := kafka.NewWriter(kafka.WriterConfig{
			Brokers: cfgs.KafkaCfg.Brokers,
			Topic:   topic,
			Dialer: &kafka.Dialer{
				Timeout:   10 * time.Second,
				DualStack: true,
				LocalAddr: &net.TCPAddr{IP: net.ParseIP("0.0.0.0")},
			},
		})
		// 你可以在这里添加更多的初始化逻辑，比如测试连接
		writers[topic] = writer
	}
	return writers
}

func initKafkaReader(cfgs conf.Config, logger *logger.ZapLogger) map[string]*kafka.Reader {
	readers := make(map[string]*kafka.Reader)
	for _, topic := range cons.KafkaTopics {
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:        cfgs.KafkaCfg.Brokers,
			Topic:          topic,
			GroupID:        cfgs.KafkaCfg.GroupID, // 假设你有一个GroupID
			StartOffset:    kafka.LastOffset,
			CommitInterval: time.Second,
		})
		// 你可以在这里添加更多的初始化逻辑，比如测试连接
		readers[topic] = reader
	}
	logger.InfofInData("init kafka reader %v", readers)
	return readers
}
