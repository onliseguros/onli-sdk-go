package brokerproduct

import (
	"context"
	"encoding/json"
	"strconv"
)

// ListBrokerProduct list broker product for broker.
func (svc service) ListBrokerProduct(ctx context.Context, req *ListBrokerProductRequest) (resp *ListBrokerProductResponse, err error) {
	params := make(map[string]string)

	if req.Rows != nil {
		params["rows"] = strconv.Itoa(*req.Rows)
	}

	if req.Page != nil {
		params["page"] = strconv.Itoa(*req.Page)
	}

	r, err := svc.client.Request(ctx).
		SetQueryParams(params).
		Get("/v1/broker-products")
	if err != nil {
		return nil, err
	}

	resp = new(ListBrokerProductResponse)
	err = json.Unmarshal(r.Body(), resp)
	return resp, err
}
