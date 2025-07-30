package entities

// CustomerAddress customer address
type CustomerAddress struct {
	Model
	CustomerID      uint   `json:"customer_id"`
	Address         string `json:"address"`
	Village         string `json:"village"`
	TelephoneNumber string `json:"telephone_number"`
	Detail          string `json:"detail"`
	Actor
}
