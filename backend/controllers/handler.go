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
	//"fmt"
	_ "github.com/lib/pq"
	//"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"
)

func AddNewCompany(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.Company
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	company := models.Company{Name: input.Name, Tel: input.Tel, Contactperson: input.Contactperson}
	dbc := db.Create(&company)
	if dbc.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": dbc.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": company})
}
func Raw(c *gin.Context) {
	db := c.MustGet("NoSQL").(*sqlx.DB)
	//var result, err = c.GetRawData()
	//if result == nil {
	//	log.Printf("this is bad")
	//}
	//var str = string(result)
	//log.Printf(str)
	//if err != nil {
	//	log.Printf("this sucks!!!")
	//}
	//hdata := models.Resulth{}
	var body, _ = ioutil.ReadAll(c.Request.Body)
	var str = string(body)
	log.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	log.Printf(str)
	//json.Unmarshal(body, &hdata)
	j := types.JSONText(string(body))
	
	//record := models.Resultd{}
	//record.Data = hdata.Data
	//var _, err = db.NamedExec(`INSERT INTO results (data) VALUES (:data)`, body)//record)
	db.MustExec(`INSERT INTO results(data) VALUES($1)`,j)//record)
	
	//var input models.Company
	/*if err != nil {
	    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	    return
	}*/
	//company := models.Company{Name: input.Name, Tel: input.Tel, Contactperson: input.Contactperson}
	//db.Create(&company)
	c.JSON(http.StatusOK, gin.H{"data": str})//record})
}

//func AddScanResult(c *gin.Context) {
//db := c.MustGet("db").(*gorm.DB)
//var input models.Result
//if err := c.ShouldBindJSON(&input); err != nil {
//  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//return
//}
//result := models.Result{Data: input.Data}
//db.Create(&result)
//c.JSON(http.StatusOK, gin.H{"data": result})
//}
/*func AddNewAccount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.Account
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account := models.Account{UserName: input.UserName, CompanyName: input.CompanyName, GoogleID: input.GoogleID}
	dbc := db.Create(&account)
	if dbc.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": dbc.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": account})
}*/

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
func AddNewEndpoint(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.Endpoint
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	endpoint := models.Endpoint{Endpoint: input.Endpoint, GoogleID: input.GoogleID}
	dbc := db.Create(&endpoint)
	if dbc.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": dbc.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": endpoint})
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

/*func GetToken(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var input models.GoogleID
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var account models.Account
	if err := db.Where("google_id = ?", input.GoogleID).Find(&account).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Record not found!"})
		return
	}
	token, err := models.CreateToken(account.GoogleID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}*/

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
