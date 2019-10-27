package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/Weltloose/grpcServer/grpcForRedis"
	gr "github.com/Weltloose/grpcServer/grpcForRedis"
	"github.com/Weltloose/grpcServer/redis"
	"google.golang.org/grpc"
)

const (
	port = ":10086"
)

type server struct{}

func (s *server) GetAuth(ctx context.Context, t *gr.Tuid) (*gr.TaInfo, error) {
	name, passwd := redis.GetAuth(t.Uid)
	return &grpcForRedis.TaInfo{Name: name, Passwd: passwd}, nil
}

func (s *server) SetAuthInfo(ctx context.Context, req *gr.ItemInfo) (*gr.Tuid, error) {
	return &gr.Tuid{Uid: redis.GenerateAuthCookie(req.Name, req.Passwd, time.Duration(req.Duration))}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to listen: ", err)
	}
	s := grpc.NewServer()
	gr.RegisterRedisOpServer(s, &server{})
	s.Serve(lis)
}
