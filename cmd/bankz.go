package main

import (
	"github.com/yakiza/BankZ/api"
	"net/http"
)

func main() {
	handler := api.MainHandler()

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		return
	}
}
