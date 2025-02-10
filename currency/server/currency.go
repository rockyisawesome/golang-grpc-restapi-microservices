package server

import (
	"context"
	protos "currency/protos/currency"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	loggs hclog.Logger
	// I dont know what is this
	protos.UnimplementedCurrencyServer
}

func NewCurrency(loggs hclog.Logger) *Currency {
	return &Currency{loggs: loggs}
}

func (currency *Currency) GetRate(ctx context.Context, rateRequest *protos.RateRequest) (*protos.RateResponse, error) {
	currency.loggs.Info("Handle request for GetRate", "base", rateRequest.GetBase(), "destination", rateRequest.GetDestination())

	return &protos.RateResponse{
		Base:        rateRequest.Base,
		Destination: rateRequest.Destination,
		Rate:        0.8,
	}, nil

}
