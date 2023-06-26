package main

import (
	"fmt"
	"log"
	"stockingapi/app/configs"
	"stockingapi/app/routes"
	"stockingapi/pkg/db"
	"time"
)

func main() {
	appConfig := configs.LoadConfig()

	redis, err := db.ConnectRedis("localhost:6379", "", "", 0)
	if err != nil {
		log.Fatal("Cannot Connect to Redis.")
		return
	}
	db, err := db.ConnectDB(appConfig.DB_HOST, appConfig.DB_USER, appConfig.DB_PASS, appConfig.DB_NAME, appConfig.DB_PORT)
	if err != nil {
		log.Fatal("Cannot Connect to Database! (ERR_DB_CONNECTION)")
		return
	}

	router := routes.SetupRoutes(db, redis)
	router.SetTrustedProxies(nil)

	if appConfig.GIN_MODE == "release" {
		go func() {
			time.Sleep(3 * time.Second)
			fmt.Printf("Starting HTTP on :%s \n", appConfig.APP_PORT)
		}()
	}
	router.Run(":" + appConfig.APP_PORT)
}
