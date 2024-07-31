package admin

import (
	"log"

	adminpb "github.com/ratheeshkumar/user_api_gateway/v2/pkg/admin/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial() (adminpb.AdminServiceClient, error) {
	grpc, err := grpc.NewClient(":8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dailing grpc client :8083")
		return nil, err
	}
	log.Printf("succefully connected to admin server at port :8083")
	return adminpb.NewAdminServiceClient(grpc), nil
}
