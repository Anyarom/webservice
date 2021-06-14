package main

import (
	"anyrom/webservice/db"
	"anyrom/webservice/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

func main() {

	// зададим настройки логирования
	log := zerolog.New(os.Stdout).With().Logger()

	sqliteClient, err := db.InitSqliteClient(&log)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	defer func() {
		err = sqliteClient.Close()
		if err != nil {
			log.Error().Err(err).Send()
		}
	}()

	// сформируем структуру БД
	errCreate := sqliteClient.CreateTable()
	if errCreate != nil {
		log.Fatal().Err(errCreate).Send()
	}

	// инициализируем структуру handler для всех запросов
	wrapperHandler := handlers.InitWrapperHandler(&log, sqliteClient)

	router := mux.NewRouter()
	router.HandleFunc("/user", wrapperHandler.CreateHandler).Methods(http.MethodPost)
	router.HandleFunc("/user", wrapperHandler.ReadHandler).Methods(http.MethodGet)
	router.HandleFunc("/user", wrapperHandler.DeleteHandler).Methods(http.MethodDelete)
	router.HandleFunc("/user", wrapperHandler.UpdateHandler).Methods(http.MethodPut)

	log.Debug().Msg("запустим web server")
	errServ := http.ListenAndServe(":8080", router)
	if errServ != nil {
		log.Fatal().Err(errServ).Caller().Send()
	}
}
