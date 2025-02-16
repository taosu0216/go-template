package interfaces

import (
	"context"
	"encoding/json"
	"fmt"
	"go-template/data"
	"go-template/data/cons"
	"go-template/logger"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

type DataRepo struct {
	*data.ToolsCtx

	*logger.ZapLogger
}

func NewDataRepo(logger *logger.ZapLogger) *DataRepo {
	repo := &DataRepo{data.Tools, logger}
	go repo.StartKafkaConsumer()
	return repo
}

func (dr *DataRepo) GetSrvAddr() string { return dr.Cfg.RouterCfg.Addr }
func (dr *DataRepo) GetSrvPort() string { return dr.Cfg.RouterCfg.Port }
func (dr *DataRepo) GetSrvExitWait() time.Duration {
	exitWait, err := strconv.Atoi(dr.Cfg.RouterCfg.ExitWait)
	if err != nil {
		return time.Second * 1
	}
	return time.Millisecond * time.Duration(exitWait)
}
func (dr *DataRepo) GetJwtKey() string { return dr.Cfg.JwtCfg.JwtSecret }

func (dr *DataRepo) GetMailInfo() (string, string, string, string) {
	return dr.Cfg.MailCfg.Addr, dr.Cfg.MailCfg.Host, dr.Cfg.MailCfg.Password, dr.Cfg.MailCfg.From
}

func (dr *DataRepo) GetVideoStoragePath() string {
	return dr.Cfg.VideoCfg.StoragePath
}

func (dr *DataRepo) SendKafkaMessage(topic string, message []byte) error {
	writer, exists := dr.KafkaWritersMaps[topic]
	if !exists {
		return fmt.Errorf("kafka writer for topic %s does not exist", topic)
	}
	// 创建Kafka消息
	msg := kafka.Message{
		Value: message,
	}

	// 向Kafka发送消息
	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		dr.Errorf("Failed to write messages to Kafka: %v", err)
		return err
	}

	dr.Infof("Message sent to Kafka topic %s", topic)
	return nil
}

func (dr *DataRepo) StartKafkaConsumer() {
	for t, reader := range dr.KafkaReadersMaps {
		go func(r *kafka.Reader, topic string) {
			for {
				msg, err := r.ReadMessage(context.Background())
				if err != nil {
					dr.ErrorfInDataKafka("Failed to fetch message: %v", err)
					continue
				}

				// 在这里处理接收到的消息
				// 可以根据实际需求进行解码或者其他处理
				dr.InfofInDataKafka("Received message: %s", string(msg.Value))

				// 根据topic进行处理
				switch topic {
				case cons.EmailRegisterNotificationTopic:
					go func() {
						var message cons.SendEmailModel
						err = json.Unmarshal(msg.Value, &message)
						if err != nil {
							// 处理错误
							dr.ErrorfInDataKafka("SendEmail [Unmarshal] err is: %v", err)
							return
						}
						userKey := fmt.Sprintf("user_%s", message.Email)
						err = dr.Cache.HSet(context.Background(), cons.UserMapKey, userKey, cons.UserStatusWaitToRegister).Err()
						if err != nil {
							dr.ErrorfInDataKafka("SendRegisterEmail [Cache HSet] err: %v", err)
						} else {
							dr.InfofInDataKafka("SendRegisterEmail [Cache HSet] success")
						}
					}()
					dr.SendEmail(msg.Value)
				case cons.EmailPasswordResetNotificationTopic:
					dr.SendEmail(msg.Value)
				}

			}
		}(reader, t)
	}
}
