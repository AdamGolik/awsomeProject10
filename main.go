package main

import (
	"awesomeProject10/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}
func main() {
	r := gin.Default()

	r.Run(":8080")
}
