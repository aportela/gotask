package database

import (
	"errors"

	"github.com/aportela/doneo/internal/config"
	sqliteWrapper "github.com/aportela/doneo/internal/database/sqlite"
)

// TODO: DSN
func Open(cfg config.DatabaseConfiguration) (Database, error) {
	switch cfg.Type {

	case "sqlite":
		databaseHandler := &sqliteWrapper.Handler{}

		if err := databaseHandler.Open(cfg); err != nil {
			return nil, err
		}

		return databaseHandler, nil

	default:
		return nil, errors.New("unknown database type")
	}
}
