package controllers

//D:\capstone\backend
import (
	models "backendAPI/models"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"time"

	//"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	//"encoding/json"
	//"github.com/tbaehler/gin-keycloak/pkg/ginkeycloak"
	"log"
	//"github.com/jmoiron/sqlx"
)

func AddNewScan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.InputScan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loc, _ := time.LoadLocation("Asia/Bangkok")
	//log.Printf(input.ScanID)
	scan := models.Scan{ScanID: input.ScanID, Company: input.Company, Status: "running", Start: time.Now().In(loc)}
	dbc := db.Create(&scan)
	if dbc.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": dbc.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": scan})
}
func Result(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.Result
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//log.Printf(input.ScanID)
	result := models.Result{ScanID: input.ScanID, Data: input.Data}
	dbc := db.Create(&result)
	if dbc.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": dbc.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}
func GenerateReport(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.ScanID
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//check if scan end
	var result models.Result
	if err := db.Where("scan_id = ?", input.ScanID).Find(&result).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "record not found"})
		return
	}
	jsonData := result.Data
	url := "http://carbone:3000/generateReport"
	jsonString := string(jsonData)
	payload := strings.NewReader(jsonString)
	fmt.Print(jsonString)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Error"})
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Error"})
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Error"})
		return
	}

	c.Data(http.StatusOK, "application/pdf", body)
}

/*func ActivateScan(c *gin.Context) {
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
}*/
/*func Raw(c *gin.Context) {
	db := c.MustGet("NoSQL").(*sqlx.DB)
	var body, _ = ioutil.ReadAll(c.Request.Body)
	var str = string(body)
	log.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	log.Printf(str)
	//json.Unmarshal(body, &hdata)
	j := types.JSONText(string(body))
	db.MustExec(`INSERT INTO results(data) VALUES($1)`, j) //record)
	c.JSON(http.StatusOK, gin.H{"data": str})              //record})
}*/
/*func GetJSON(c *gin.Context) {
	db := c.MustGet("NoSQL").(*sqlx.DB)
	var body, _ = ioutil.ReadAll(c.Request.Body)
	var str = string(body)
	log.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	log.Printf(str)
	//json.Unmarshal(body, &hdata)
	j := types.JSONText(string(body))
	db.MustExec(`SELECT * FROM results WHERE (data #>> scanid)`, j) //record)
	c.JSON(http.StatusOK, gin.H{"data": str})              //record})
}*/
func AddEndpoint(c *gin.Context) { //input poolid, endpoint
	db := c.MustGet("db").(*gorm.DB)
	company, okCompany := c.Get("company")
	if okCompany {
		strCompany := fmt.Sprintf("%v", company)
		strCompany = strings.Trim(strCompany, "[/]")
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
		strCompany = strings.Trim(strCompany, "[/]")
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
	log.Println("GetEndpoints")
	db := c.MustGet("db").(*gorm.DB)
	company, okCompany := c.Get("company")
	if okCompany {
		strCompany := fmt.Sprintf("%v", company)
		strCompany = strings.Trim(strCompany, "[/]")
		var endpoints []models.Endpoint
		if err := db.Where("company = ?", strCompany).Find(&endpoints).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "record not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": endpoints})
	}
}

/*func AddScan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	company, okCompany := c.Get("company")
	if okCompany {
		strCompany := fmt.Sprintf("%v", company)
		var input models.Scan
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		scan := models.Scan{ScanID: input.ScanID, Company: input.Company, Status: input.Status}
		dbc := db.Create(&scan)
		if dbc.Error != nil {
			c.JSON(http.StatusOK, gin.H{"error": dbc.Error})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": pool})
	}
}*/

/*func GetScan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("scan_id = ?", c.Param("scan_id")).Find(&hosts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hosts})
}
func ActivateScan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	company, okCompany := c.Get("company")
	if okCompany {
		strCompany := fmt.Sprintf("%v", company)

		scan := models.Scan{ScanID: input.ScanID, Company: input.Company, PoolID: input.PoolID, Status: input.Status}
		dbc := db.Create(&scan)
		if dbc.Error != nil {
			c.JSON(http.StatusOK, gin.H{"error": dbc.Error})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": pool})
	}
}*/

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
}*/
func UpdateScan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var scan models.Scan
	if err := db.Where("scan_id = ?", c.Param("scan_id")).First(&scan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.ScanStatus
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	db.Model(&scan).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": scan})
}
func GetScanStatus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var input models.ScanID
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	var scan models.Scan
	if err := db.Where("scan_id = ?", input.ScanID).First(&scan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": scan.Status})
}

/*func DeleteEndpoint(c *gin.Context) {
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
