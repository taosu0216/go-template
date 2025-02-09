package conf

type MailCfg struct {
	Addr     string `json:"Addr"`
	Host     string `json:"Host"`
	Password string `json:"Password"`
	From     string `json:"From"`
}
