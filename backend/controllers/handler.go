package controllers
import (
    models "backendAPI/models"
    "net/http"
    "github.com/gin-gonic/gin"
    //"database/sql"
    "github.com/jinzhu/gorm"
	//"fmt"
	//"log"
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
