package cmd

import (
	"github.com/fadedpez/driver/internal/handlers"
	driverapialpha "github.com/fadedpez/driver/protos"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
)

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("starting server")
		address := "localhost:5000"
		lis, err := net.Listen("tcp", address)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()

		driverapialpha.RegisterDriverAPIServer(s, &handlers.Alpha{})

		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	rootCommand.AddCommand(serverCommand)
}