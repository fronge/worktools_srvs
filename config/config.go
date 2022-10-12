package config

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	Name     string `json:"user_name"`
	Password string `json:"password"`
	DB       string `json:"db"`
}

type ConsulConfig struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}

type ServerConfig struct {
	MysqlInfo  MysqlConfig  `json:"mysql"`
	ConsulInfo ConsulConfig `json:"consul"`
	Name       string       `json:"name"`
	Host       string       `json:"host"`
	Port       int64        `json:"port"`
}
