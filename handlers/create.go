package handlers

import (
	"anyrom/webservice/types"
	"net/http"
	"time"
)

func (wrapHandler *WrapperHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	if name == "" {
		wrapHandler.log.Error().Msg("параметр name не корректный")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	birthDate := r.URL.Query().Get("birth_date")
	if birthDate == "" {
		wrapHandler.log.Error().Msg("параметр birthDate не корректный")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// birth_date в date
	date, errParse := time.Parse(types.Layout, birthDate)
	if errParse != nil {
		wrapHandler.log.Error().Err(errParse).Caller().Msg("не удалось преобразовать дату в формат")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	wrapHandler.log.Debug().Msgf("получены данные name %s, birthDate %s", name, date.String())

	// запись в БД
	errInsert := wrapHandler.dbClient.InsertUser(name, date)
	if errInsert != nil {
		wrapHandler.log.Error().Err(errInsert).Caller().Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	wrapHandler.log.Debug().Msgf("пользователь %s записан в БД", name)

}
