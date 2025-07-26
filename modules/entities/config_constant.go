package entities

// ConfigConstant configuration constant entity
type ConfigConstant struct {
	Model
	Group       string `json:"group"`
	NameEn      string `json:"name_en"`
	NameTh      string `json:"name_th"`
	Option      string `json:"option"`
	Description string `json:"description"`
	Sort        int    `json:"sort"`
	IsActive    int8   `json:"is_active"`
}
