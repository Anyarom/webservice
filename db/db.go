package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"os"
	"sync"
)

type Client struct {
	*sql.DB
	mx  sync.RWMutex
	log *zerolog.Logger
}

func InitSqliteClient(log *zerolog.Logger) (*Client, error) {

	dbPath := "users.sqlite"
	var fileDb *os.File
	_, err := os.Stat(dbPath)
	if err == nil {
		fileDb, err = os.Open(dbPath)
		if err != nil {
			return nil, err
		}
	} else {
		fileDb, err = os.Create(dbPath)
		if err != nil {
			log.Error().Err(err).Send()
			return nil, err
		}
	}

	defer func() {
		err := fileDb.Close()
		if err != nil {
			log.Error().Caller().Err(err).Send()
		}
	}()

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return &Client{
		DB:  db,
		log: log,
	}, nil
}
