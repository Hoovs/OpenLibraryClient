package main

import (
	"errors"
	"github.com/Hoovs/OpenLibraryClient/server/handlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	portDefault      = ":8080"
	searchAPIBaseURI = "http://openlibrary.org/search.json?q="
)

var (
	logger *zap.Logger
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("ok")); err != nil {
		zap.Error(errors.New("couldn't write out status message"))
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = portDefault
	}

	logger, _ = zap.NewProduction()
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error("Couldn't sync log")
		}
	}()
	logger.Info("Running server", zap.String("port", port))

	r := mux.NewRouter()
	// Handle the status endpoint
	r.HandleFunc("/status", statusHandler).Methods("GET")
	sh := handlers.SearchHandler{
		Logger:        logger,
		BaseSearchUrl: searchAPIBaseURI,
	}

	wh := handlers.WishListHandler{
		Logger: logger,
	}

	r.HandleFunc("/search", sh.SearchHandler).Methods("GET")
	r.HandleFunc("/wishList", wh.GetWishListHandler).Methods("POST")
	r.HandleFunc("/wishList", nil).Methods("DELETE")

	// Bind to a port and pass our router in
	logger.Fatal(http.ListenAndServe(port, r).Error())
}
