package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	core "github.com/shoplineapp/captin/core"
	incoming "github.com/shoplineapp/captin/incoming"
	stores "github.com/shoplineapp/captin/internal/stores"
	models "github.com/shoplineapp/captin/models"
)

func main() {
	fmt.Println("Starting in port:", getEnv("CAPTIN_PORT", "3000"))
	port := fmt.Sprintf(":%s", getEnv("CAPTIN_PORT", "3000"))

	// Load webhooks configuration
	pwd, _ := os.Getwd()
	path := os.Args[1:][0]
	absPath := filepath.Join(pwd, path)
	configMapper := models.NewConfigurationMapperFromPath(absPath)
	captin := core.NewCaptin(*configMapper)

	// Set up api server
	router := gin.Default()
	handler := incoming.HttpEventHandler{}
	handler.Setup(*captin)
	handler.SetRoutes(router)

	fmt.Printf("* Binding captin on 0.0.0.0%s\n", port)
	router.Run(port)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
