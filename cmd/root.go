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

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello", args)
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
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
