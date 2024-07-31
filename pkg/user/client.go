package user

import (
	"log"

	pb "github.com/ratheeshkumar/user_api_gateway/v2/pkg/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial() (pb.UserServiceClient, error) {
	grpc, err := grpc.NewClient(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dailing grpc client :8082")
		return nil, err
	}
	log.Printf("succefully connected to admin server at port :8082")
	return pb.NewUserServiceClient(grpc), nil
}
