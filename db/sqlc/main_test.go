package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/datmaithanh/simplebank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries

var testDB *sql.DB
var err error

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Can not load config", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)

	code := m.Run()
	os.Exit(code)
}
