package main

import (
	fiber_swagger "fiber-swagger"
	"github.com/gofiber/fiber/v2"
	"net"
	"strconv"
)

func main() {
	app := fiber.New()

	listener, _ := net.Listen("tcp", net.JoinHostPort("localhost", strconv.Itoa(8080)))

	swaggerMiddleware := fiber_swagger.NewMiddleware("./docs/swagger_example.json", "/")

	swaggerMiddleware.Register(app)

	_ = GetHandler()

	SetupRoutes(app)

	_ = app.Listener(listener)
}
