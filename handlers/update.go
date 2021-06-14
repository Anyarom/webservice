package handlers

import (
	"anyrom/webservice/types"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type BodyUpdate struct {
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
}

func (wrapHandler *WrapperHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {

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

	// достанем тело
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wrapHandler.log.Error().Err(err).Caller().Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func() {
		err := r.Body.Close()
		if err != nil {
			wrapHandler.log.Error().Err(err).Caller().Send()
		}
	}()

	// парсинг тела ответа
	var bodyUpdate BodyUpdate
	err = json.Unmarshal(body, &bodyUpdate)
	if err != nil {
		wrapHandler.log.Error().Err(err).Caller().Send()
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// birth_date в date
	date, errParse := time.Parse(types.Layout, bodyUpdate.BirthDate)
	if errParse != nil {
		wrapHandler.log.Error().Err(errParse).Caller().Send()
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// обновление в БД
	errUpdate := wrapHandler.dbClient.UpdateUser(bodyUpdate.Name, date, idInt)
	if errUpdate == types.ErrNoRows {
		wrapHandler.log.Error().Err(errUpdate).Caller().Send()
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if errUpdate != nil {
		wrapHandler.log.Error().Err(errUpdate).Caller().Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	wrapHandler.log.Debug().Msgf("данные в БД успешно обновлены %s", bodyUpdate.Name)
}
