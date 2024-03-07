package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	driverName = "postgres"
	dataSource = "postgres://ed:password@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		log.Fatalln("Could not connect to DB", err)
	}

	testQueries = New(db)

	os.Exit(m.Run())
}
