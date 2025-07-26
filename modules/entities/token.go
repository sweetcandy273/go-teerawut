package entities

// Token token model
type Token struct {
	Model
	AccessToken string `json:"access_token" gorm:"-"`
	UserName    string `json:"username"`
}
