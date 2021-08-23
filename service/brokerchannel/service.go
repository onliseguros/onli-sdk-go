package brokerchannel

import (
	"context"
	"github.com/onliseguros/onli-sdk-go/onli"
)

// ServiceBrokerChannel represents the broker channel, deals with
// channels available for broker.
type ServiceBrokerChannel interface {
	ListBrokerChannel(ctx context.Context, req *ListBrokerChannelRequest) (resp *ListBrokerChannelResponse, err error)
}

type service struct {
	client *onli.Client
}

// New initializes a new broker channel service.
func New(client *onli.Client) ServiceBrokerChannel {
	return &service{
		client: client,
	}
}
