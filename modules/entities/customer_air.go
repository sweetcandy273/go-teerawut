package entities

// CustomerAirCondition customer air condition
type CustomerAirCondition struct {
	Model
	CustomerID     uint   `json:"customer_id"`
	AirConditionID *uint  `json:"air_condition_id,omitempty"`
	BtuID          *uint  `json:"btu_id,omitempty"`
	RoomName       string `json:"room_name"`
	FromUs         bool   `json:"from_us"`
}
