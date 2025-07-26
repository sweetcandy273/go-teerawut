package entities

import "time"

// CustomerServiceList customer service list entity
type CustomerServiceList struct {
	Model
	CustomerID  uint      `json:"customer_id"`
	Date        time.Time `json:"date"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	Actor
}

// CustomerService customer service entity
type CustomerService struct {
	ID                    uint   `json:"id"`
	CustomerServiceListID uint   `json:"customer_service_list_id"`
	TypeID                uint   `json:"type_id"`
	Description           string `json:"description"`
}

// CustomerServiceItem customer service item entity
type CustomerServiceItem struct {
	CustomerServiceID uint `json:"customer_service_id"`
	ItemID            uint `json:"item_id"`
}
