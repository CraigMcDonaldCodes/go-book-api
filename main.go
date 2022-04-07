package main

import (
	"fmt"
	"net/http"

	"github.com/craigmcdonaldcodes/go-book-api/handlers"
)

const address = "localhost:8080"

func main() {

	fmt.Println("Go Book API")

	http.HandleFunc("/books", handlers.GetAll)
	err := http.ListenAndServe(address, nil)

	if err != nil {
		fmt.Println("Error launching web server:", err)
	}
}
