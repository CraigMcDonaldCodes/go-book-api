package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/craigmcdonaldcodes/go-book-api/data"
)

var logger = log.Default()

type Handlers struct {
	db data.Db
}

func NewHandlers(db data.Db) *Handlers {

	return &Handlers{
		db: db,
	}
}

func (h *Handlers) GetAll(wr http.ResponseWriter, req *http.Request) {

	wr.Header().Set("Content-Type", "application/json")

	books := h.db.GetAll()

	json, err := json.Marshal(books)

	if err != nil {
		// TODO return a 500 error?
	}

	logger.Printf("URL: %s, Client IP: %s\n", req.URL, req.RemoteAddr)
	wr.Write(json)
}
