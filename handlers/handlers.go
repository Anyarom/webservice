package handlers

import (
	"anyrom/webservice/db"
	"github.com/rs/zerolog"
)

type WrapperHandler struct {
	log      *zerolog.Logger
	dbClient *db.Client
}

func InitWrapperHandler(log *zerolog.Logger, sqliteClient *db.Client) *WrapperHandler {
	return &WrapperHandler{log: log, dbClient: sqliteClient}
}
