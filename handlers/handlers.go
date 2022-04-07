package handlers

import "net/http"

func GetAll(wr http.ResponseWriter, req *http.Request) {
	wr.Write([]byte("Get all the books"))
}
