package brokerchannel

// BrokerChannel represents the broker channel,
// always associated with an broker.
type BrokerChannel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ListBrokerChannelRequest prepares the request to list broker channels.
type ListBrokerChannelRequest struct {
	Rows *int
	Page *int
}

// ListBrokerChannelResponse defines the response of broker channels.
type ListBrokerChannelResponse struct {
	Total int              `json:"total"`
	Items []*BrokerChannel `json:"items"`
}
