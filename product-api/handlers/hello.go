package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"product-api/database"

	"github.com/hashicorp/go-hclog"
)

type Hello struct {
	loggs   *hclog.Logger
	mongodb *database.MongoDB
	ctx     context.Context
}

func NewHello(loggs *hclog.Logger, mgdb *database.MongoDB, ctxe context.Context) *Hello {
	return &Hello{
		loggs:   loggs,
		mongodb: mgdb,
		ctx:     ctxe,
	}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	(*h.loggs).Info("A call made to Hello GetMethod")

	d, err := io.ReadAll(r.Body)
	if err != nil {
		(*h.loggs).Error("Some error occured", err)
		http.Error(rw, "some error", http.StatusBadGateway)
		return
	}

	(*h.loggs).Info("request body", "body", string(d))
	data, err := (*h.mongodb).ListUsers(h.ctx)
	if err != nil {
		(*h.loggs).Error("SOme issue occured in data fetching")
		http.Error(rw, "Dekho error aa gayi bhai", http.StatusBadRequest)
		return
	}

	(*h.loggs).Info("Data fetched properly")

	json.NewEncoder(rw).Encode(data)
	return
	// response
	// rw.Write([]byte("Response mil gaya he"))
}
