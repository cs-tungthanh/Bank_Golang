package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/cs-tungthanh/Bank_Golang/api"
	db "github.com/cs-tungthanh/Bank_Golang/db/sqlc"
	"github.com/cs-tungthanh/Bank_Golang/util"
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
		server, err := api.NewServer(config, store)
		if err != nil {
			log.Fatal("cannot create server:", err)
		}

		err = server.Start(config.ServerAddress)
		if err != nil {
			log.Fatal("Cannot start server:", err)
		}
	},
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
