package main

import (
	"github.com/gin-gonic/gin"

	"SimpleFizzBuzz/server"
)

func main() {
	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(server.Logger())
	srv := server.NewServer(r)
	srv.RegisterRoutes()

	srv.Run()
}
