package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type conn struct {
	db *sql.DB
}

func Open(ctx context.Context, logger *logrus.Logger) (*conn, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	logger.Info("Successfully connected to the database!")
	return &conn{db: db}, nil
}
