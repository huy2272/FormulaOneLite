package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/huy2272/FormulaOneLite/internal/controller"
)

func main() {
	server := gin.Default()

	controller.RegisterRoutes(server)
	server.Run(":8080") // listening on localhost:8080
}
