package admin

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ratheeshkumar/user_api_gateway/v2/pkg/admin/handler"
	adminpb "github.com/ratheeshkumar/user_api_gateway/v2/pkg/admin/pb"
)

type AdminRoutes struct {
	client adminpb.AdminServiceClient
}

func NewAdminRoute(c *gin.Engine) {
	client, err := ClientDial()
	if err != nil {
		log.Fatalf("error Not connected with gRPC server, %v", err.Error())
	}

	adminHandler := AdminRoutes{
		client: client,
	}

	apiAdmin := c.Group("/api/admin")
	{
		apiAdmin.POST("/login", adminHandler.Login)
		apiAdmin.POST("/create/user", adminHandler.CreateUser)
		apiAdmin.PUT("/edit/user", adminHandler.EditUser)
		apiAdmin.DELETE("delete/user", adminHandler.DeleteUser)
		apiAdmin.GET("/users", adminHandler.FindAllUsers)
		apiAdmin.GET("/id/user", adminHandler.FindUserByID)
		apiAdmin.GET("/email/user", adminHandler.FindUserByEmail)
	}
}

func (a *AdminRoutes) Login(c *gin.Context) {
	handler.AdminLoginHandler(c, a.client)
}

func (a *AdminRoutes) CreateUser(c *gin.Context) {
	handler.CreateUserHandler(c, a.client)
}

func (a *AdminRoutes) EditUser(c *gin.Context) {
	handler.EditUserHandler(c, a.client)
}

func (a *AdminRoutes) DeleteUser(c *gin.Context) {
	handler.DeleteUserHandler(c, a.client)
}

func (a *AdminRoutes) FindAllUsers(c *gin.Context) {
	handler.FindAllUsersHandler(c, a.client)
}

func (a *AdminRoutes) FindUserByEmail(c *gin.Context) {
	handler.FindUserByEmailHandler(c, a.client)
}

func (a *AdminRoutes) FindUserByID(c *gin.Context) {
	handler.FindUserByIDHandler(c, a.client)
}
