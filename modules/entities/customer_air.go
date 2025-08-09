package entities

// CustomerAirCondition customer air condition
type CustomerAirCondition struct {
	Model
	CustomerID uint   `json:"customer_id"`
	AirBrandID *uint  `json:"air_brand_id,omitempty"`
	AirTypeID  *uint  `json:"air_type_id,omitempty"`
	BtuID      *uint  `json:"btu_id,omitempty"`
	RoomName   string `json:"room_name"`
	FromUs     *bool  `gorm:"column:from_us;type:boolean" json:"from_us,omitempty"`
}
