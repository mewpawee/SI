package models

import (
	"encoding/json"
	//"time"
)

type InputEndpoint struct {
	Endpoint string `form:"endpoint" json:"endpoint"`
}
type Endpoint struct {
	Endpoint string `form:"endpoint" json:"endpoint" gorm:"primary_key"`
	Company  string `form:"company" json:"company"`
}
type Scan struct {
	ScanID  string `json:"scanid" gorm:"primary_key"`
	Company string `json:"company"`
	Status  string `json:"status"`
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
type Resulth struct {
	Data json.RawMessage
}
type Resultd struct {
	Data string `db:"data"`
}
