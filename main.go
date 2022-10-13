package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"worktools_srvs/global"
	"worktools_srvs/handler"
	"worktools_srvs/initialize"
	"worktools_srvs/libraries"
	"worktools_srvs/proto"

	"go.uber.org/zap"
	grpc "google.golang.org/grpc"
)

func main() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitMySQL()

	zap.S().Info(fmt.Sprintf("正在开启服务 %s:%d", global.ServerConfig.Host, global.ServerConfig.Port))

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", global.ServerConfig.Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	// 只在线上使用consul
	if initialize.ENV == initialize.PRO {
		libraries.HealthCheck(server)
		libraries.RegisterConsul()
	}

	// 运行服务
	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic(err)
		}
		zap.S().Infof("正在开启服务 %s:%d", global.ServerConfig.Host, global.ServerConfig.Port)
	}()

	// 停止服务信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 线上的需要注销consul
	if initialize.ENV == initialize.PRO {
		if err = libraries.ConsulServiceDeregister(); err != nil {
			zap.S().Info("注销失败")
		}
	}
	zap.S().Info("注销成功")
}
