package initialize

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"

	"worktools_srvs/global"
)

var ENV string

const (
	DEV  = "dev"
	TEST = "test"
	PRO  = "pro"
)

func InitConfig() {
	flag.StringVar(&ENV, "e", "", "set env, e.g dev prod")
	flag.Parse()

	if len(ENV) == 0 {
		panic("Not found ENV")
	}

	path := filepath.Join(".", "config", ENV)
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("load config error %v", err))
	}

	// MYSQL
	viperMysqlCfg := viper.GetStringMap("mysql")
	global.ServerConfig.MysqlInfo.Host = viperMysqlCfg["host"].(string)
	global.ServerConfig.MysqlInfo.Port = viperMysqlCfg["port"].(int64)
	global.ServerConfig.MysqlInfo.Name = viperMysqlCfg["name"].(string)
	global.ServerConfig.MysqlInfo.Password = viperMysqlCfg["password"].(string)
	global.ServerConfig.MysqlInfo.DB = viperMysqlCfg["db"].(string)

	// ServerConfig
	vipeServerCfg := viper.GetStringMap("server")
	global.ServerConfig.Port = vipeServerCfg["port"].(int64)
	global.ServerConfig.Host = vipeServerCfg["host"].(string)
	global.ServerConfig.Name = vipeServerCfg["name"].(string)

	// ConsulConfig
	consulCfg := viper.GetStringMap("consul")
	global.ServerConfig.ConsulInfo.Port = consulCfg["port"].(int64)
	global.ServerConfig.ConsulInfo.Host = consulCfg["host"].(string)
}
