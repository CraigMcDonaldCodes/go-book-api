package handlers

import (
	"net/http"

	"github.com/craigmcdonaldcodes/go-book-api/data"
)

type Handlers struct {
	db data.Db
}

func NewHandlers(db data.Db) *Handlers {

	return &Handlers{
		db: db,
	}
}

func (h *Handlers) GetAll(wr http.ResponseWriter, req *http.Request) {
	wr.Write([]byte("Get all the books"))
}
