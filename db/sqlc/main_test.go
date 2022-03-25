package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/cs-tungthanh/Bank_Golang/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

// TestMain func is entry point of all unittests inside a package
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config file:", err)
	}
	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDb)

	// to start running the unit test
	os.Exit(m.Run())
}
