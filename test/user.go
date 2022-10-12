package main

import (
	"context"
	"fmt"
	"worktools_srvs/proto"

	"google.golang.org/grpc"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("114.116.50.177:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}

func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{Pn: 1, PSize: 20})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.NickName)
		checkRsp, err := userClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			Password:          "addmain",
			EncryptedPassword: user.PassWord,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkRsp.Success)
	}
}

func TestCreateUser() {
	for i := 10; i < 20; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			NickName: fmt.Sprintf("Speike%d", i),
			Mobile:   fmt.Sprintf("136143345%d", i),
			PassWord: "addmain",
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(rsp.Id)
	}
}

func main() {
	Init()
	TestGetUserList()
	// TestCreateUser()
	conn.Close()

}
