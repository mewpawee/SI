package models

type Company struct {
	Name string `json:"name"`
	Tel string `json:"tel"`
	Contactperson string `json:"contactperson"`
}
type Account struct {
	UserName string `json:"username"`
	CompanyName string `json:"companyname"`
	GoogleID string `json:"googleid"`
}
type logs struct {
	ID int `json:"id"` 
	Log string `json:"log"` 
}