package conf

type CacheCfg struct {
	Addr         string `json:"Addr"`
	User         string `json:"User"`
	Password     string `json:"Password"`
	ReadTimeout  int    `json:"ReadTimeout"`
	WriteTimeout int    `json:"WriteTimeout"`
}
