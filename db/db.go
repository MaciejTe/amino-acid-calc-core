package db

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"

	_ "github.com/lib/pq" // postgreSQL driver
	log "github.com/sirupsen/logrus"
)

// NewDbClient returns postgresql database handle client.
func NewDbClient() (*sql.DB, error) {
	log.Debug("Getting database connection parameters from configuration file...")
	dbConfig := viper.GetStringMapString("db")
	log.Debug("Getting database connection parameters from configuration file... DONE")
	connectionInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig["host"], dbConfig["port"], dbConfig["user"], dbConfig["password"], dbConfig["dbname"])

	db, err := sql.Open("postgres", connectionInfo)
	if err != nil {
		log.Fatalf("Failed to connect to database! Details: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database! Details: %v", err)
	}

	log.Debug("Successfully connected to the database!")

	return db, err
}
