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

type LoggerConfig struct {
	Dir      string `json:"dir"`
	FileName string `json:"file_name"`
	Level    string `json:"level"`
}

type ServerConfig struct {
	MysqlInfo  MysqlConfig  `json:"mysql"`
	ConsulInfo ConsulConfig `json:"consul"`
	LoggerInfo LoggerConfig `json:"logger"`
	Name       string       `json:"name"`
	Host       string       `json:"host"`
	Port       int64        `json:"port"`
}
