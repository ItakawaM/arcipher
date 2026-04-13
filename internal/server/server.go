package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ItakawaM/arcipher/ciphers"
	"github.com/pkg/browser"
)

func handleSubmit(w http.ResponseWriter, r *http.Request) *ciphers.CardanKey {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return nil
	}

	var key *ciphers.CardanKey
	if err := json.NewDecoder(r.Body).Decode(&key); err != nil {
		http.Error(w, "Invalid JSON Body!", http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)

	return key
}

func GetCardanKeyUI(gridSize int, ctx context.Context) <-chan *ciphers.CardanKey {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./internal/server/static"))) // TODO: Think of a way to change this hardcoded value?

	srv := &http.Server{Addr: ":8080", Handler: mux}
	ch := make(chan *ciphers.CardanKey, 1)

	mux.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		key := handleSubmit(w, r)
		if key != nil {
			ch <- key
		}
	})

	go func() {
		<-ctx.Done()
		srv.Shutdown(context.Background())
	}()

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	fmt.Printf("Cardan UI running at http://localhost:8080?n=%d\n", gridSize)
	browser.OpenURL(fmt.Sprintf("http://localhost:8080?n=%d", gridSize))
	return ch
}
