package conf

type VideoCfg struct {
	// 视频存储路径
	StoragePath string `json:"storagePath"`
	// 允许的域名列表
	AllowedDomains []string `json:"allowedDomains"`
	// 链接有效期(秒)
	TokenExpire int64 `json:"tokenExpire"`
	// 密钥
	SecretKey string `json:"secretKey"`
}

func (c *Config) GetVideoStoragePath() string {
	return c.VideoCfg.StoragePath
}

func (c *Config) GetVideoAllowedDomains() []string {
	return c.VideoCfg.AllowedDomains
}

func (c *Config) GetVideoTokenExpire() int64 {
	if c.VideoCfg.TokenExpire <= 0 {
		return 7200 // 默认2小时
	}
	return c.VideoCfg.TokenExpire
}

func (c *Config) GetVideoSecretKey() string {
	return c.VideoCfg.SecretKey
}
