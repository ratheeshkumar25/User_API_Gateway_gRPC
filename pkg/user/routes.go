package user

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ratheeshkumar/user_api_gateway/v2/pkg/user/handler"
	pb "github.com/ratheeshkumar/user_api_gateway/v2/pkg/user/pb"
)

type User struct {
	client pb.UserServiceClient
}

func (u *User) Login(c *gin.Context) {
	handler.UserLoginHandler(c, u.client, "user")
}

func (u *User) Signup(c *gin.Context) {
	handler.UserSignupHandler(c, u.client)
}

func NewUserRoute(c *gin.Engine) {
	client, err := ClientDial()
	if err != nil {
		log.Fatalf("error not connected with gRPC server, %v", err.Error())
	}

	userHandler := User{
		client: client,
	}
	apiUser := c.Group("/api/user")
	{
		apiUser.POST("/login", userHandler.Login)
		apiUser.POST("/signup", userHandler.Signup)
	}
}
