package product

import "github.com/onliseguros/onli-sdk-go/service/insurer"

// Product represents the insurer product itself.
type Product struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Currency    string           `json:"currency"`
	Description string           `json:"description"`
	Susep       string           `json:"susep"`
	Type        int              `json:"type"`
	InsurerID   string           `json:"insurer_id"`
	Insurer     *insurer.Insurer `json:"insurer"`
}
