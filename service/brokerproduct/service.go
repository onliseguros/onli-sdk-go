package brokerproduct

import (
	"context"
	"github.com/onliseguros/onli-sdk-go/onli"
)

// ServiceBrokerProduct represents the broker product, deals with
// products available for broker.
type ServiceBrokerProduct interface {
	ListBrokerProduct(ctx context.Context, req *ListBrokerProductRequest) (resp *ListBrokerProductResponse, err error)
}

type service struct {
	client *onli.Client
}

// New initializes a new broker product service.
func New(client *onli.Client) ServiceBrokerProduct {
	return &service{
		client: client,
	}
}
