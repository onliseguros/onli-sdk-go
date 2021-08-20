package merchant

import (
	"context"
	"github.com/onliseguros/onli-sdk-go/onli"
)

// ServiceMerchant represents the broker merchant, deals with
// products, plans and channels available for him.
type ServiceMerchant interface {
	GetStore(ctx context.Context, brokerChannelID string) (resp *GetStoreResponse, err error)
	GetVendor(ctx context.Context, productID string) (resp *GetVendorResponse, err error)
}

type service struct {
	client *onli.Client
}

// New initializes a new merchant service.
func New(client *onli.Client) ServiceMerchant {
	return &service{
		client: client,
	}
}
