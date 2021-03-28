package controllers

//D:\capstone\backend
import (
	models "backendAPI/models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx/types"

	//"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	//"encoding/json"
	//"github.com/tbaehler/gin-keycloak/pkg/ginkeycloak"
	"log"

	"github.com/jmoiron/sqlx"
)


func Raw(c *gin.Context) {
	db := c.MustGet("NoSQL").(*sqlx.DB)
	var body, _ = ioutil.ReadAll(c.Request.Body)
	var str = string(body)
	log.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	log.Printf(str)
	//json.Unmarshal(body, &hdata)
	j := types.JSONText(string(body))
	db.MustExec(`INSERT INTO results(data) VALUES($1)`, j) //record)
	c.JSON(http.StatusOK, gin.H{"data": str})              //record})
}

func AddEndpoint(c *gin.Context) { //input poolid, endpoint
	db := c.MustGet("db").(*gorm.DB)
	company, okCompany := c.Get("company")
	if okCompany {
		strCompany := fmt.Sprintf("%v", company)
		strCompany := strings.Trim(strCompany, "[/]")
		var input models.InputEndpoint
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		endpoint := models.Endpoint{Endpoint: input.Endpoint, Company: strCompany}
		dbc := db.Create(&endpoint)
		if dbc.Error != nil {
			c.JSON(http.StatusOK, gin.H{"error": dbc.Error})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": endpoint})
	}
	c.JSON(http.StatusOK, gin.H{"error": okCompany})
}
func DeleteEndpoint(c *gin.Context) { //input endpoint
	db := c.MustGet("db").(*gorm.DB)
	company, okCompany := c.Get("company")
	if okCompany {
		strCompany := fmt.Sprintf("%v", company)
		strCompany := strings.Trim(strCompany, "[/]")
		var input models.InputEndpoint
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//log.Printf(input.Endpoint)
		var endpoint models.Endpoint
		if err := db.Where("endpoint = ?", input.Endpoint).First(&endpoint).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "Record not found!"})
			return
		}
		if endpoint.Company != strCompany {
			c.JSON(http.StatusForbidden, gin.H{"error": "User do not have permission to delete this endpoint"})
			return
		}
		db.Delete(&endpoint)
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
func GetEndpoints(c *gin.Context) { // input poolid
	db := c.MustGet("db").(*gorm.DB)
	company, okCompany := c.Get("company")
	if okCompany {
		strCompany := fmt.Sprintf("%v", company)
		strCompany := strings.Trim(strCompany, "[/]")
		var endpoints []models.Endpoint
		if err := db.Where("endpoint = ?", strCompany).Find(&endpoints).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "record not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": endpoints})
	}
}
