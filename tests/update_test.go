package tests

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
	"sync"
	"testing"
)

func TestUpdate(t *testing.T) {
	var wg sync.WaitGroup
	routines := 1000
	wg.Add(routines)
	start := make(chan struct{})
	for i := 0; i < routines; i++ {
		go func() {
			<-start
			defer wg.Done()
			client := &http.Client{}
			req, err := http.NewRequest("PUT", "http://127.0.0.1:8080/user?id=2", strings.NewReader(`{
				 "name": "Hahan",
				 "birth_date": "2000-12-20"
				}`))
			if err != nil {
				t.Error(err)
				return
			}
			res, err := client.Do(req)
			if err != nil {
				t.Error(err)
				return
			}
			log.Debug().Msg(res.Status)
		}()
	}
	close(start)
	wg.Wait()
}
