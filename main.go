//go:generate go run build/goa_gen.go
// Package main is where all the API pieces like models, controllers, services,
// middlewares and configuration handlers glue together. The only expected file
// in this Package is "main.go", this enforces to move all the autogenrated
// controllers (check README.md file for generated code details) to their
// respective package.
//
// Note: For CopperMantis CLI tool refer to tool directory.

package main

import (
	"fmt"
	"strings"

	"github.com/CopperMantis/CopperMantis/app"
	"github.com/CopperMantis/CopperMantis/controllers"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"

	"github.com/spf13/viper"
)

func main() {

	// Create service
	service := goa.New("CopperMantis")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Load Configurations
	viper.SetConfigName(".copper-mantis")                  // name of config file (without extension)
	viper.AddConfigPath("$HOME")                           // adding user's home directory as first search path
	viper.AddConfigPath(".")                               // adding cwd as second search path
	viper.SetEnvPrefix("cms")                              // adding CMS environment variable prefix
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Add env variable support for nested configs
	viper.AutomaticEnv()                                   // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		service.LogInfo("Configuration", "file", viper.ConfigFileUsed())
	}

	// Open db connection
	pgconf := fmt.Sprintf("dbname=%s user=%s password=%s sslmode=%s port=%d host=%s",
		viper.GetString("database.schema"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.sslmode"),
		viper.GetInt("database.port"),
		viper.GetString("database.host"))
	fmt.Println(pgconf)

	// Mount "contest" controller
	c := controllers.NewContestController(service)
	app.MountContestController(service, c)
	// Mount "document" controller
	c2 := controllers.NewDocumentController(service)
	app.MountDocumentController(service, c2)
	// Mount "monitoring" controller
	c3 := controllers.NewMonitoringController(service)
	app.MountMonitoringController(service, c3)
	// Mount "profile" controller
	c4 := controllers.NewProfileController(service)
	app.MountProfileController(service, c4)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
