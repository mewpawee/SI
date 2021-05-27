package models

import (
	//"encoding/json"
	"time"
	"github.com/jmoiron/sqlx/types"
)

/*type Result struct {
	Data JSON `json:"company"`
	ScanID  string `json:"scanid"`
}*/

type InputEndpoint struct {
	Endpoint string `form:"endpoint" json:"endpoint"`
}
type Endpoint struct {
	Endpoint string `form:"endpoint" json:"endpoint" gorm:"primary_key"`
	Company  string `form:"company" json:"company"`
}
type InputScan struct {
	ScanID  string `json:"scanid"`
	Company string `json:"company"`
}
type ScanID struct {
	ScanID  string `json:"scanid"`
}
type ScanStatus struct {
	ScanID  string `json:"scanid" gorm:"primary_key"`
	Status  string `json:"status"`
}
type ScanEnd struct {
	ScanID  string `json:"scanid" gorm:"primary_key"`
	Status  string `json:"status"`
	Complete	time.Time `json:"Complete"`
}
type Scan struct {
	ScanID  string `json:"scanid" gorm:"primary_key"`
	Company string `json:"company"`
	Status  string `json:"status"`
	Start	time.Time `json:"start"`
}
type ScanInfo struct {
	ScanID  string `json:"scanid" gorm:"primary_key"`
	Company string `json:"company"`
	Status  string `json:"status"`
	Start	time.Time `json:"start"`
	Complete	time.Time `json:"Complete"`
}
type Scanlog struct {
	ScanID       string `json:"scanid" gorm:"primary_key"`
	ScanEndpoint string `json:"scanendpoint"`
}
type Report struct {
	ScanID string `json:"scanid" gorm:"primary_key"`
	DIR    string `json:"dir"`
}
type Log struct {
	FileName string `json:"filename" gorm:"primary_key"`
	ScanID   string `json:"scanid"`
	ToolName string `json:"toolname"`
	DIR      string `json:"dir"`
}
type Result struct {
	ScanID string `json:"scanid" gorm:"primary_key"`
	Data   types.JSONText `json:"data"`
}
type Resultd struct {
	Data string `db:"data"`
}
type Booking struct {
	DayStart string `json:"daystart"`
	TimeStart time.Time `json:"timestart"`
	DayEnd string `json:"dayend"`
	TimeEnd time.Time `json:"timeend"`
}
