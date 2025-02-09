package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	AiCfg     *AiCfg     `json:"AiCfg"`
	DBCfg     *DBCfg     `json:"DBCfg"`
	CacheCfg  *CacheCfg  `json:"CacheCfg"`
	RouterCfg *RouterCfg `json:"RouterCfg"`
	MailCfg   *MailCfg   `json:"MailCfg"`
	JwtCfg    *JwtCfg    `json:"JwtCfg"`
}

func InitCfg() Config {
	cfg := Config{}
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper Read Config failed, err:%v\n", err)
		return Config{}
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err = viper.Unmarshal(&cfg); err != nil {
		fmt.Printf("viper Unmarshal failed, err:%v\n", err)
		return Config{}
	}

	return cfg
}
