package handlers

import (
	"net/http"
	"strconv"
)

func (wrapHandler *WrapperHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		wrapHandler.log.Error().Caller().Msg("параметр id не корректный")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		wrapHandler.log.Error().Err(err).Caller().Send()
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// удаление записи из БД
	errDelete := wrapHandler.dbClient.DeleteUser(idInt)
	if errDelete != nil {
		wrapHandler.log.Error().Err(err).Caller().Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	wrapHandler.log.Debug().Msgf("пользователь удален из БД с id %s", id)
}
