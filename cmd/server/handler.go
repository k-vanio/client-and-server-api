package main

import (
	"encoding/json"
	"net/http"

	"github.com/x-vanio/client-and-server-api/internal/db"
	"github.com/x-vanio/client-and-server-api/pkg/request"
)

type handlerQuote struct {
	client     request.HTTPClient
	repository db.Repository
}

type responseErr struct {
	Err string `json:"err"`
}

type responseBid struct {
	Bid string `json:"bid"`
}

func NewHandler(client request.HTTPClient, repository db.Repository) *handlerQuote {
	return &handlerQuote{client: client, repository: repository}
}

func (h *handlerQuote) GetDollarQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := h.client.ByDollarQuote("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		json.NewEncoder(w).Encode(&responseErr{Err: err.Error()})
		return
	}

	if err := h.repository.Save(*res); err != nil {
		json.NewEncoder(w).Encode(&responseErr{Err: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&responseBid{Bid: res.Currency.Bid})
}
