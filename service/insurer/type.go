package insurer

// Insure is the insurance company that undertakes
// to indemnify for losses and perform other
// insurance-related operations.
type Insurer struct {
	ID   string `json:"id"`
	Cnpj string `json:"cnpj"`
	Name string `json:"name"`
}
