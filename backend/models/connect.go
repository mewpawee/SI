package models
import (
    //models "backendAPI/models"
    //"net/http"
    //"github.com/gin-gonic/gin"
    //"database/sql"
    "github.com/jinzhu/gorm"
	"fmt"
	//"log"
)

func ConnectPsql() *gorm.DB {
	dbUri := "<POSTGRESQL_URI>"
	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}
	return db
}
