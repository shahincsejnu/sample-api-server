package server

import (
	"github.com/gin-gonic/gin"
	"github.com/shahincsejnu/sample-api-server/endpoints"
	"log"
)

func Start() {
	r := gin.Default()

	r.GET("/cluster/health", endpoints.GetServerHealthHandler)
	r.GET("/cluster/services", endpoints.GetAllServicesHandler)

	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	} // localhost:8080
}
