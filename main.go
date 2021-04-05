package main

import (
	"fmt"
	"net/http"
	"time"

	brandRepoCommand "github.com/andrefebrianto/rest-api/src/domains/Brand/repositories/postgres/commands"
	brandRepoQuery "github.com/andrefebrianto/rest-api/src/domains/Brand/repositories/postgres/queries"
	brandUseCase "github.com/andrefebrianto/rest-api/src/domains/Brand/usecases"
	productRepoCommand "github.com/andrefebrianto/rest-api/src/domains/Product/repositories/postgres/commands"
	productRepoQuery "github.com/andrefebrianto/rest-api/src/domains/Product/repositories/postgres/queries"
	productUseCase "github.com/andrefebrianto/rest-api/src/domains/Product/usecases"
	controller "github.com/andrefebrianto/rest-api/src/httpcontroller"
	"github.com/andrefebrianto/rest-api/src/utilities/databases/elasticsearch"
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
	contextTimeout := time.Duration(config.GetInt("context.timeout")) * time.Second

	// Echo instance
	httpServer := echo.New()

	// Middleware
	httpServer.Use(middleware.Logger())
	httpServer.Use(middleware.Recover())
	httpServer.Use(setCorsHeader)

	// Routes
	httpServer.GET("/", healthCheck)

	//Initialize PostgreSQL database
	var postgreSQLConfigs []map[string]interface{}
	err := config.UnmarshalKey("postgreSQL", &postgreSQLConfigs)
	if err != nil {
		panic(err)
	}

	postgresql.InitConnection(postgreSQLConfigs)

	//Initialize MongoDB database
	var mongoDBConfigs []map[string]interface{}
	err = config.UnmarshalKey("mongoDB", &mongoDBConfigs)
	if err != nil {
		panic(err)
	}

	//Initialize ElastocSearch database
	var elasticSearchConfigs []map[string]interface{}
	err = config.UnmarshalKey("elasticsearch", &elasticSearchConfigs)
	if err != nil {
		panic(err)
	}

	elasticsearch.InitConnection(elasticSearchConfigs)

	pgClient, err := postgresql.GetConnection("product-db")
	if err != nil {
		panic(err)
	}

	esClient, err := elasticsearch.GetConnection("product-catalog-db")
	if err != nil {
		panic(err)
	}

	fmt.Println(esClient)
	//Create Repositories
	brandCommand := brandRepoCommand.CreateRepository(pgClient)
	brandQuery := brandRepoQuery.CreateRepository(pgClient)

	productCommand := productRepoCommand.CreateRepository(pgClient)
	productQuery := productRepoQuery.CreateRepository(pgClient)

	//Create Use Cases
	brandUC := brandUseCase.CreateBrandUseCase(brandCommand, brandQuery, contextTimeout)

	productUC := productUseCase.CreateProductUseCase(productCommand, productQuery, contextTimeout)

	controller.CreateBrandHandler(httpServer, brandUC)
	controller.CreateProductHandler(httpServer, productUC)
	controller.CreatePurchaseHandler(httpServer, productUC)

	// Start server
	httpServer.Logger.Fatal(httpServer.Start(config.GetString("server.port")))
}
