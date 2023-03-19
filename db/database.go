package db

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/tushargupta98/api-in-go/config"
)

var DB *sqlx.DB

func init() {
	dbConfig := config.GetConfig().Database
	err := initPostgreSQL(dbConfig)
	if err != nil {
		log.WithError(err).Fatal("Error initializing PostgreSQL database")
	}
}

func initPostgreSQL(dbConfig config.DatabaseConfig) error {
	// Construct the data source name (DSN) using the provided configuration
	port, err := strconv.Atoi(dbConfig.Port)
	if err != nil {
		log.Error("Error parsing cache idle_timeout:", err)
	}
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s sslrootcert=%s search_path=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		port,
		dbConfig.Name,
		dbConfig.SSLMode,
		dbConfig.SSLRootCert,
		dbConfig.Schema,
	)

	// Create a new context with a 10-second timeout for the connection attempt
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a new database connection with the given DSN and context
	DB, err = sqlx.ConnectContext(ctx, dbConfig.Driver, dsn)
	if err != nil {
		log.Error("api_in_go DB Connection Error: ", err)
		return err
	}

	// Set connection pool parameters
	DB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	DB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	DB.SetConnMaxLifetime(time.Duration(dbConfig.MaxConnLifetime) * time.Minute)

	// Ping the database to ensure the connection is working
	if err := DB.PingContext(ctx); err != nil {
		log.Error("Ping to api_in_go DB Failed: ", err.Error())
		return err
	}

	log.Info("api_in_go Database Connection Successful.")
	return nil
}
