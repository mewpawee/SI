package models
import (
	"encoding/json"
	//"time"
)
type InputPool struct {
	PoolID        string `json:"poolid"` 
	//Company       string `json:"company"`
}
type Pool struct {
	PoolID        string `json:"poolid" gorm:"primary_key"` 
	Company       string `json:"company"`
}
type InputEndpoint struct {
	Endpoint      string `json:"endpoint"`
}
type Endpoint struct {
	Endpoint      string `json:"endpoint" gorm:"primary_key"`
	PoolID        string `json:"poolid"`
}
type Scan struct {
	ScanID        string `json:"scanid" gorm:"primary_key"`
	Company       string `json:"company"`
	PoolID        string `json:"poolid"`
	Status        string `json:"status"` 
}
type Scanlog struct {
	ScanID        string `json:"scanid" gorm:"primary_key"`
	ScanEndpoint  string `json:"scanendpoint"`
}
type Report struct {
	ScanID        string `json:"scanid" gorm:"primary_key"`
	DIR           string `json:"dir"`
}
type Log struct {
	FileName      string `json:"filename" gorm:"primary_key"`
	ScanID        string `json:"scanid"`
	ToolName      string `json:"toolname"`
	DIR           string `json:"dir"`
}
type Resulth struct {
	Data    json.RawMessage 
}
type Resultd struct {  
	Data    string  `db:"data"`
}
