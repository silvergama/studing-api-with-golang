package repository

import (
	"database/sql"
	"fmt"
	"sync/atomic"

	"github.com/silvergama/studying-api-with-golang/config"
)

var (
	DB      *sql.DB
	dbReady = int32(0)
)

func setupDB() error {
	if isDBReady() {
		return nil
	}
	config := config.GetConfig()
	dbURI := fmt.Sprintf("%s:%s@%s:%s/%s",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
	)

	db, err := sql.Open(config.DB.Dialect, dbURI)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	DB = db

	dbIsReady()

	return nil
}

func isDBReady() bool {
	return atomic.LoadInt32(&dbReady) == 1
}

func dbIsReady() {
	atomic.StoreInt32(&dbReady, 1)
}
