package di

import (
	"github.com/ratheeshkumar/user_api_gateway/v2/pkg/admin"
	"github.com/ratheeshkumar/user_api_gateway/v2/pkg/server"
	"github.com/ratheeshkumar/user_api_gateway/v2/pkg/user"
)

// func Init() {
// 	Server := server.Server()
// 	user.NewUserRoute(Server.R)
// 	admin.NewAdminRoute(Server.R)
// 	server.Server().StartServer()
// }

func Init() {
	// Create a single instance of the server
	s := server.Server()

	// Set up the routes using the same instance
	user.NewUserRoute(s.R)
	admin.NewAdminRoute(s.R)

	// Start the server using the same instance
	s.StartServer()
}
