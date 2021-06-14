package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type RespRead struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
}

func (wrapHandler *WrapperHandler) ReadHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		wrapHandler.log.Error().Caller().Msg("Параметр id не корректный")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		wrapHandler.log.Error().Err(err).Caller().Send()
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// получим из БД имя пользователя
	user, birthDate, errGet := wrapHandler.dbClient.GetUser(idInt)
	if errGet == sql.ErrNoRows {
		wrapHandler.log.Error().Err(errGet).Caller().Msg("нет такого id в БД")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if errGet != nil {
		wrapHandler.log.Error().Err(errGet).Caller().Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(RespRead{Id: idInt, Name: user, BirthDate: birthDate})
	if err != nil {
		wrapHandler.log.Error().Err(err).Caller().Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	_, errRespBody := w.Write(respBody)
	if errRespBody != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
