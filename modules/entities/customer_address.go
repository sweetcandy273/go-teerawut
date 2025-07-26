package entities

// CustomerAddress customer address
type CustomerAddress struct {
	Model
	CustomerID      uint   `json:"customer_id"`
	Address         string `json:"address"`
	Village         string `json:"village"`
	Moo             string `json:"moo"`
	Soi             string `json:"soi"`
	Road            string `json:"road"`
	SubDistrict     string `json:"sub_district"`
	District        string `json:"district"`
	Province        string `json:"province"`
	ZipCode         string `json:"zip_code"`
	TelephoneNumber string `json:"telephone_number"`
	Detail          string `json:"detail"`
	Actor
}
