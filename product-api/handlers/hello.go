package handlers

import (
	"io"
	"net/http"

	"github.com/hashicorp/go-hclog"
)

type Hello struct {
	loggs *hclog.Logger
}

func NewHello(loggs *hclog.Logger) *Hello {
	return &Hello{
		loggs: loggs,
	}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	(*h.loggs).Info("We are in Hello struct handler method")

	d, err := io.ReadAll(r.Body)
	if err != nil {
		(*h.loggs).Error("Some error occured", err)
		http.Error(rw, "some error", http.StatusBadGateway)
		return
	}

	(*h.loggs).Info("request body", "body", string(d))

	// response
	rw.Write([]byte("Response mil gaya he"))
}
