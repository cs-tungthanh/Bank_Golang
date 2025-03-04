package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/cs-tungthanh/Bank_Golang/composer"
	db "github.com/cs-tungthanh/Bank_Golang/db/sqlc"
	"github.com/cs-tungthanh/Bank_Golang/util"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

// db
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database management commands",
}

// db migrate
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Simple Bank",
	Long:  `All software has versions. This is Simple Bank's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Simple Bank Application v1.0.0 -- HEAD")
	},
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start service with args:", args)

		config, err := util.LoadConfig(".")
		if err != nil {
			log.Fatal("cannot load config: ", err)
		}

		conn, err := sql.Open(config.DBDriver, config.DBSource)
		if err != nil {
			log.Fatal("cannot connect to db: ", err)
		}

		store := db.NewStore(conn)
		router := gin.Default()
		SetupRoutes(config, store, router)

		err = router.Run(config.ServerAddress)
		if err != nil {
			log.Fatal("Cannot start server:", err)
		}
	},
}

func SetupRoutes(cfg util.Config, store db.Store, router *gin.Engine) {
	apiService, err := composer.ComposeAPIService(cfg, store)
	if err != nil {
		log.Fatal("Cannot create API service:", err)
	}

	userGroup := router.Group("/users")
	{
		userGroup.POST("/", apiService.UserAPI.CreateUser)
		userGroup.POST("login", apiService.UserAPI.LoginUser)
	}

	accountGroup := router.Group("/accounts")
	{
		accountGroup.POST("/", apiService.AccountAPI.CreateAccount)
		accountGroup.POST("/transfers", apiService.AccountAPI.CreateTransfer)

		accountGroup.GET("/:id", apiService.AccountAPI.GetAccount)
		accountGroup.GET("/", apiService.AccountAPI.ListAccount)
	}
}

func Execute() {
	dbCmd.AddCommand(migrateCmd)

	rootCmd.AddCommand(dbCmd)
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("[rootCmd.Execute] error: %v\n", err)
		os.Exit(1)
	}
}
