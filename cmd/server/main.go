package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"time"

	"github.com/lmittmann/tint"
	auth "github.com/phrkdll/doomsday/internal/server"
	pb "github.com/phrkdll/doomsday/protos"
	"github.com/phrkdll/must/pkg/must"
	"google.golang.org/grpc"
)

var (
	Host = "localhost"
	Port = "1982"
)

func main() {
	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{TimeFormat: time.RFC3339})))
	slog.Info("Starting doomsday server")

	listener := must.Return(net.Listen("tcp", fmt.Sprintf("%v:%v", Host, Port))).ElsePanic()

	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	pb.RegisterAuthServiceServer(grpcServer, &auth.AuthService{})

	slog.Info("Server listening", "host", Host, "port", Port)

	must.Succeed(grpcServer.Serve(listener)).ElsePanic()
}
