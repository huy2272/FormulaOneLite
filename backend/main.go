package backend

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.Run(":8080") // listening on localhost:8080
}
