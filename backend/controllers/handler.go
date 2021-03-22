package controllers

//D:\capstone\backend
import (
	models "backendAPI/models"
	"net/http"
	"github.com/jmoiron/sqlx/types"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	//"database/sql"
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/lib/pq"
	//"encoding/json"
	//"github.com/tbaehler/gin-keycloak/pkg/ginkeycloak"
	"github.com/jmoiron/sqlx"
	"log"
)

func AddPool(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	company, okCompany := c.Get("company")
	if okCompany {
		strCompany := fmt.Sprintf("%v", company)
		input := c.Param("poolid")//models.Pool
		/*if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/
		pool := models.Pool{PoolID: input, Company: strCompany}
		dbc := db.Create(&pool)
		if dbc.Error != nil {
			c.JSON(http.StatusOK, gin.H{"error": dbc.Error})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": pool})
	}
}
func Raw(c *gin.Context) {
	db := c.MustGet("NoSQL").(*sqlx.DB)
	var body, _ = ioutil.ReadAll(c.Request.Body)
	var str = string(body)
	log.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	log.Printf(str)
	//json.Unmarshal(body, &hdata)
	j := types.JSONText(string(body))
	db.MustExec(`INSERT INTO results(data) VALUES($1)`,j)//record)
	c.JSON(http.StatusOK, gin.H{"data": str})//record})
}

func AddEndpoint(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	company, okCompany := c.Get("company")
	if okCompany {
		strCompany := fmt.Sprintf("%v", company)
		var input models.Endpoint
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var pool models.Pool 
		if err := db.Where("pool_id = ?", c.Param("poolid")).First(&pool).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "Pool does not exist"})
			return
		}
		if pool.Company != strCompany {
			c.JSON(http.StatusForbidden, gin.H{"error": "Why the hell are you try to add endpoint to other company's endpoint pool?")
			return
		}
		endpoint := models.Endpoint{Endpoint: input.Endpoint, PoolID: input.PoolID}
		dbc := db.Create(&endpoint)
		if dbc.Error != nil {
			c.JSON(http.StatusOK, gin.H{"error": dbc.Error})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": endpoint})
	}
}
func DeleteEndpoint(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	company, okCompany := c.Get("company")
	if okCompany {
		strCompany := fmt.Sprintf("%v", company)
		var endpoint models.Endpoint
		if err := db.Where("endpoint = ?", c.Param("endpoint")).First(&endpoint).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "Record not found!"})
			return
		}
		var pool models.Pool 
		if err := db.Where("pool_id = ?", endpoint.PoolID).First(&pool).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "Pool does not exist"})
			return
		}
		if pool.Company != strCompany {
			c.JSON(http.StatusForbidden, gin.H{"error": "Why the hell are you try to delete endpoint in other company's endpoint pool?")
			return
		}
		dbc := db.Delete(&endpoint)
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
func GetCompanyPools(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	company, okCompany := c.Get("company")
	if okCompany {
		strCompany := fmt.Sprintf("%v", company)
		var pools []models.Pool
		if err := db.Where("company = ?", strCompany).Find(&pools).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "Record not found!"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": pools})
	}
}
func GetPoolEndpoints(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	company, okCompany := c.Get("company")
	if okCompany {
		strCompany := fmt.Sprintf("%v", company)
		var pool models.Pool 
		if err := db.Where("pool_id = ?", c.Param("poolid")).Find(&pool).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "Pool does not exist"})
			return
		}
		if pool.Company != strCompany {
			c.JSON(http.StatusForbidden, gin.H{"error": "Why the hell are you peak in to other company's pool?")
			return
		}
		var endpoints []models.Endpoint
		if err := db.Where("pool_id = ?", c.Param("poolid")).Find(&endpoints).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "Pool does not exist in endpoints"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": endpoints})
	}
}
/*

func AddNewScan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.Scan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf(input.ScanID)
	scan := models.Scan{ScanID: input.ScanID, GoogleID: input.GoogleID, Status: input.Status}
	dbc := db.Create(&scan)
	if dbc.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": dbc.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": scan})
}
func UpdateScan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var scan models.Scan
	if err := db.Where("scan_id = ?", c.Param("scan_id")).First(&scan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.Scan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	db.Model(&scan).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": scan})
}

func DeleteEndpoint(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var endpoint models.Endpoint
	log.Printf(c.Param("endpoint"))
	if err := db.Where("endpoint = ?", c.Param("endpoint")).First(&endpoint).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&endpoint)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func GetTestData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var test models.TestData
	if err := db.Where("id = ?", c.Param("id")).First(&test).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": test})
}

func GetEndpoints(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var endpoints []models.Endpoint
	if err := db.Where("google_id = ?", c.Param("google_id")).Find(&endpoints).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": endpoints})
}

func GetHost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	log.Printf("host")
	// Get model if exist
	var hosts []models.Host
	if err := db.Where("scan_id = ?", c.Param("scan_id")).Find(&hosts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hosts})
}

func GetVul(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var vulnerability []models.Vulnerability
	if err := db.Where("scan_id = ?", c.Param("scan_id")).Find(&vulnerability).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": vulnerability})
}
*/
