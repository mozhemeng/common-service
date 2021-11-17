package cmd

import (
	pb "common_service/internal/proto"
	"common_service/internal/rpc"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func runRpcServer() {
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, rpc.NewUserServer())
	reflection.Register(s)

	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatal(err)
	}

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

var rpcCmd = &cobra.Command{
	Use:   "rpc",
	Short: "run rpc server",
	Run: func(cmd *cobra.Command, args []string) {
		runRpcServer()
	},
}
