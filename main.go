package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	static "github.com/gin-contrib/static"
	gin "github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"

	application "./application"
	config "./config"
	database "./database"
)

func main() {
	// init app directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	configDir := flag.String("config_dir", dir, "Find your configuration file location")
	if err := config.AppConfig(*configDir); err != nil {
		log.Fatal(err)
	}

	mode := os.Getenv("MODE")
	fmt.Printf("-------------- APPLICATION MODE: %s --------------\n\n", mode)
	if mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// save api request log to file
	if os.Getenv("DEBUG_LOG") == "active" {
		debugLog, err := os.Create(fmt.Sprintf("%s/log/debugLog-%d.log", dir, time.Now().Unix()))
		if err != nil {
			log.Fatal(err)
		}
		gin.DefaultWriter = io.MultiWriter(debugLog, os.Stdout)
		log.SetOutput(gin.DefaultWriter)
	}

	// init app
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	// set CORS
	app.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, POST, PUT, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, X-Requested-With, Accept",
		Credentials:     true,
		ValidateHeaders: false,
	}))

	// Database initialization
	if err := database.Init(); err != nil {
		log.Fatal("Error init database connection.", err)
	}

	// init application module
	application.InitAppModule(app)

	// serve static html file
	app.Use(static.Serve("/", static.LocalFile(fmt.Sprintf("%s/public", os.Getenv("APP_DIR")), true)))

	// run server
	port := ":" + os.Getenv("PORT")
	if err := app.Run(port); err != nil {
		log.Fatal(err)
	}
}
