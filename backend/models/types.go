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
type Host struct {
	IpAddress string `json:"ipaddress"` 
	ScanID string `json:"scanid"` 
	HostName string `json:"hostname"` 
	HostVulnerabilty float64 `json:"hostvulnerabilty"`
}
type TestData struct {
	ID int `json:"id"` 
	Name string `json:"name"`
	Word string `json:"word"`
}
type Vulnerability struct {
	ScanID string `json:"scanid"` 
	Critical int `json:"critical"` 
	High int `json:"high"` 
	Low int `json:"high"` 
	SystemVulnerabilty float64 `json:"systemvulnerabilty"`
}
type Endpoint struct {
	Endpoint int `json:"endpoint"` 
	GoogleID string `json:"googleid"` 
}