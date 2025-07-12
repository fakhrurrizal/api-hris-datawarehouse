package main

import (
	router "hris-datawarehouse/app/routers"
	"hris-datawarehouse/config"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gopkg.in/tylerb/graceful.v1"
)

// @title HRIS Data Warehouse
// @version V1.2412.081710
// @description API documentation by Kang Fakhrur

// @securityDefinitions.apikey JwtToken
// @in header
// @name Authorization

func main() {
	app := echo.New()
	router.Init(app)
	config.Database()
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	app.Server.Addr = "127.0.0.1:" + config.LoadConfig().Port
	log.Printf("Server: " + config.LoadConfig().BaseUrl)
	log.Printf("Documentation: " + config.LoadConfig().BaseUrl + "/docs")
	graceful.ListenAndServe(app.Server, 5*time.Second)
}
