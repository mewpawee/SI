package models

import (
	//models "backendAPI/models"
	//"net/http"
	//"github.com/gin-gonic/gin"
	//"database/sql"
	//"database/sql"
	"fmt"
	//"log"

	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectPsql() *gorm.DB {
	dbUri := "<POSTGRESQL_URI>"
	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}
	return db
}
func ConnectPNoSQL() *sqlx.DB {
	dsn := "<POSTGRESQL_URI>"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		fmt.Println("Error in postgres connection: ", err)
	}
	return db
}

