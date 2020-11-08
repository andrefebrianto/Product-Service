package main

import (
	"fmt"
	"net/http"

	"github.com/andrefebrianto/rest-api/src/utilities/databases/postgresql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

// Handler
func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Service is running properly")
}

func setCorsHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		context.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(context)
	}
}

func loadConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigFile(`configs.json`)
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if config.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

	return config
}

func main() {
	config := loadConfig()

	// Echo instance
	httpServer := echo.New()

	// Middleware
	httpServer.Use(middleware.Logger())
	httpServer.Use(middleware.Recover())
	httpServer.Use(setCorsHeader)

	// Routes
	httpServer.GET("/", healthCheck)

	var postgreSQLConfigs []map[string]interface{}
	err := config.UnmarshalKey("postgreSQL", &postgreSQLConfigs)
	if err != nil {
		panic(err)
	}

	// postgreSQLConnectionPool := postgresql.ConnectionPool{}
	postgreSQLConnectionPool := postgresql.InitConnection(postgreSQLConfigs)
	fmt.Println(postgreSQLConnectionPool.Connections)

	var mongoDBConfigs []map[string]interface{}
	err = config.UnmarshalKey("mongoDB", &mongoDBConfigs)
	if err != nil {
		panic(err)
	}

	// Start server
	// httpServer.Logger.Fatal(httpServer.Start(viper.GetString("server.port")))
}
