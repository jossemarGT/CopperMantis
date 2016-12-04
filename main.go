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
	"github.com/CopperMantis/CopperMantis/models"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
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
	db, err := getDB(viper.Sub("database"), "postgres")
	if err != nil {
		service.LogError("Couldn't open database connection", "err", err)
	}

	// Mount "contest" controller
	c := controllers.NewContestController(service, db)
	app.MountContestController(service, c)
	// Mount "document" controller
	c2 := controllers.NewDocumentController(service, db)
	app.MountDocumentController(service, c2)
	// Mount "monitoring" controller
	c3 := controllers.NewMonitoringController(service)
	app.MountMonitoringController(service, c3)
	// Mount "profile" controller
	c4 := controllers.NewProfileController(service, db)
	app.MountProfileController(service, c4)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}

func getDB(v *viper.Viper, backend string) (*gorm.DB, error) {
	v.SetDefault("sslmode", "disabled")
	v.SetDefault("max_connections", 50)
	v.SetDefault("log_mode", false)
	v.SetDefault("migrate", false)

	url := fmt.Sprintf("dbname=%s user=%s password=%s sslmode=%s port=%d host=%s",
		v.GetString("schema"),
		v.GetString("user"),
		v.GetString("password"),
		v.GetString("sslmode"),
		v.GetInt("port"),
		v.GetString("host"))

	db, err := gorm.Open(backend, url)
	if err != nil {
		return nil, err
	}
	db.LogMode(v.GetBool("log_mode"))
	db.DB().SetMaxOpenConns(v.GetInt("max_connections"))

	if v.GetBool("migrate") {
		migrateDB(db)
	}

	return db, nil
}

func migrateDB(db *gorm.DB) {
	// TODO: Find a fancy way to determine the model list
	db.AutoMigrate(models.Contest{}, models.Document{}, models.Profile{})
}
