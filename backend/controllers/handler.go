package controllers
import (
    models "backendAPI/models"
    "net/http"
    "github.com/gin-gonic/gin"
    //"database/sql"
    "github.com/jinzhu/gorm"
	//"fmt"
	"log"
)

func AddNewCompany(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
// Validate input
    var input models.Company
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
// Create Book
    company := models.Company{Name: input.Name, Tel: input.Tel, Contactperson: input.Contactperson}
    db.Create(&company)
    c.JSON(http.StatusOK, gin.H{"data": company})
}

func AddNewAccount(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
// Validate input
    var input models.Account
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
// Create Book
    account := models.Account{UserName: input.UserName, CompanyName: input.CompanyName, GoogleID: input.GoogleID}
    db.Create(&account)
    c.JSON(http.StatusOK, gin.H{"data": account})
}

func AddNewScan(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
// Validate input
    var input models.Account
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
// Create Book
    account := models.Account{UserName: input.UserName, CompanyName: input.CompanyName, GoogleID: input.GoogleID}
    db.Create(&account)
    c.JSON(http.StatusOK, gin.H{"data": account})
}

func AddNewEndpoint(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
// Validate input
    var input models.Endpoint
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
// Create Book
    endpoint := models.Endpoint{Endpoint: input.Endpoint, GoogleID: input.GoogleID}
    db.Create(&endpoint)
    c.JSON(http.StatusOK, gin.H{"data": endpoint})
}

func DeleteEndpoint(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
   // Get model if exist
    var endpoint models.Endpoint
    log.Printf(c.Param("endpoint"))
    if err := db.Where("endpoint = ?", c.Param("endpoint")).First(&endpoint).Error; err != nil {
     c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
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
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": test})
}

func GetEndpoints(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
   // Get model if exist
    var endpoints[] models.Endpoint
    if err := db.Where("google_id = ?", c.Param("google_id")).Find(&endpoints).Error; err != nil {
     c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
     return
    }
   c.JSON(http.StatusOK, gin.H{"data": endpoints})
}

func GetHost(c *gin.Context) {
 db := c.MustGet("db").(*gorm.DB)
 log.Printf("host")
// Get model if exist
 var hosts[] models.Host
 if err := db.Where("scan_id = ?", c.Param("scan_id")).Find(&hosts).Error; err != nil {
  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
  return
 }
c.JSON(http.StatusOK, gin.H{"data": hosts})
}

func GetVul(c *gin.Context) {
 db := c.MustGet("db").(*gorm.DB)
// Get model if exist
 var vulnerability[] models.Vulnerability
 if err := db.Where("scan_id = ?", c.Param("scan_id")).Find(&vulnerability).Error; err != nil {
  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
  return
 }
c.JSON(http.StatusOK, gin.H{"data": vulnerability})
}