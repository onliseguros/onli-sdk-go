package brokerchannel

import (
	"context"
	"encoding/json"
	"strconv"
)

// ListBrokerChannel list broker channel for broker.
func (svc service) ListBrokerChannel(ctx context.Context, req *ListBrokerChannelRequest) (resp *ListBrokerChannelResponse, err error) {
	params := make(map[string]string)

	if req.Rows != nil {
		params["rows"] = strconv.Itoa(*req.Rows)
	}

	if req.Page != nil {
		params["page"] = strconv.Itoa(*req.Page)
	}

	r, err := svc.client.Request(ctx).
		SetQueryParams(params).
		Get("/v1/broker-channels")
	if err != nil {
		return nil, err
	}

	resp = new(ListBrokerChannelResponse)
	err = json.Unmarshal(r.Body(), resp)
	return resp, err
}
