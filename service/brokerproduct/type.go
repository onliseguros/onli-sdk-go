package brokerproduct

import "github.com/onliseguros/onli-sdk-go/service/product"

// BrokerProduct represents the broker product,
// always associated with an insurer product but carries
// data for each broker related to it.
type BrokerProduct struct {
	ID                 string           `json:"id"`
	ProductID          string           `json:"product_id"`
	Name               string           `json:"name"`
	CommercialImage    string           `json:"commercial_image"`
	CommercialText     string           `json:"commercial_text"`
	CommissionPercent  int              `json:"commission_percent"`
	Description        string           `json:"description"`
	ParticipationRules string           `json:"participation_rules"`
	Product            *product.Product `json:"product"`
}

// ListBrokerProductRequest prepares the request to list broker products.
type ListBrokerProductRequest struct {
	Rows *int
	Page *int
}

// ListBrokerProductResponse defines the response of broker products.
type ListBrokerProductResponse struct {
	Total int              `json:"total"`
	Items []*BrokerProduct `json:"items"`
}
