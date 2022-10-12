package libraries

import (
	"fmt"
	"worktools_srvs/global"

	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

var (
	client *api.Client
	srvID  string
)

// 注册consul
func RegisterConsul() {
	var err error

	cfg := api.DefaultConfig()

	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	client, err = api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	srvID = fmt.Sprintf("%s", uuid.NewV4())

	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	registration.ID = srvID
	registration.Port = int(global.ServerConfig.Port)
	registration.Tags = []string{"frange", "user_srv"}
	// registration.Address = global.ServerConfig.Host
	registration.Address = "114.116.50.177"

	registration.Check = &api.AgentServiceCheck{
		Method: "TCP",
		TCP:    fmt.Sprintf("%s:%d", "114.116.50.177", global.ServerConfig.Port),
		// TCP:                            fmt.Sprintf("%s:%d", global.ServerConfig.Host, global.ServerConfig.Port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
}

// 健康检查
func HealthCheck(server *grpc.Server) {
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	zap.S().Info("注册服务健康检查", global.ServerConfig.Host)
}

// 注销consul
func ConsulServiceDeregister() (err error) {
	return client.Agent().ServiceDeregister(srvID)
}
