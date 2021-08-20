package merchant

import (
	"context"
	"encoding/json"
)

// GetStore get the broker merchant store from broker channel.
func (svc service) GetStore(ctx context.Context, brokerChannelID string) (resp *GetStoreResponse, err error) {
	r, err := svc.client.Request(ctx).
		SetPathParam("broker_channel_id", brokerChannelID).
		Get("/v1/merchant/stores/{broker_channel_id}")
	if err != nil {
		return nil, err
	}

	resp = new(GetStoreResponse)
	err = json.Unmarshal(r.Body(), resp)
	return resp, err
}

// GetVendor get the broker merchant vendor from product.
func (svc service) GetVendor(ctx context.Context, productID string) (resp *GetVendorResponse, err error) {
	r, err := svc.client.Request(ctx).
		SetPathParam("product_id", productID).
		Get("/v1/merchant/vendor/{product_id}")
	if err != nil {
		return nil, err
	}

	resp = new(GetVendorResponse)
	err = json.Unmarshal(r.Body(), resp)
	return resp, err
}
