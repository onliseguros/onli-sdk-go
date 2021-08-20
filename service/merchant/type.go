package merchant

import (
	"github.com/onliseguros/onli-sdk-go/service/brokerproduct"
	"time"
)

// GetStoreResponse represents the Broker’s Channel distribution configuration.
// Allows the Broker to control his products available in different channels,
// configuring some details important to control the behavior for the merchant in channels.
type GetStoreResponse struct {
	BrokerChannelID string                         `json:"broker_channel_id"`
	LogoURL         string                         `json:"logo_url"`
	BrokerProducts  []*brokerproduct.BrokerProduct `json:"broker_products"`
}

// GetVendorResponse represents the relationship of the Broker and Insurers,
// in the way in which Product and Plans are being available to Brokers.
// It establishes a reliable bond between the two, controlling aspects
// like commission, agency, and pro-labor percentages.
type GetVendorResponse struct {
	ID                string    `json:"id"`
	ProductID         string    `json:"product_id"`
	Active            bool      `json:"active"`
	AgencyPercent     float64   `json:"agency_percent"`
	CommissionPercent float64   `json:"commission_percent"`
	ProLaborPercent   float64   `json:"pro_labor_percent"`
	LastUpdate        time.Time `json:"last_update"`
	Date              time.Time `json:"date"`
}
