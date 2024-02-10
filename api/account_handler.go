package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/yakiza/BankZ/accounts"
	"github.com/yakiza/BankZ/api/internal/marshaller"
	"net/http"
)

type AccountHandler struct {
	chi.Router
	useCase accounts.UseCase
}

func (h AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account marshaller.Account
	err := json.NewDecoder(r.Body).Decode(&account)

	err = h.useCase.Create(r.Context(), marshaller.UnmarshalAccount(account))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func NewAccountHandler() AccountHandler {
	h := AccountHandler{
		Router: chi.NewRouter(),
	}

	h.Post("/create", h.CreateAccount)

	return h
}
