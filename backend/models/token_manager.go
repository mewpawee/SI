package models

import (
	"google.golang.org/api/oauth2/v2"
	"net/http"
	"github.com/gin-gonic/gin"
)

var httpClient = &http.Client{}
func GetUserID(c *gin.Context) (string, error){
	//
	header := c.Request.Header["Token"]
	token := header[0]
	oauth2Service, err := oauth2.New(httpClient)
    tokenInfoCall := oauth2Service.Tokeninfo() 
    tokenInfoCall.AccessToken(token)
    tokenInfo, err := tokenInfoCall.Do()
    if err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return "", err
    }
	
	userID := tokenInfo.UserId // this thing works, comment bc I don't use it now 
	return userID, nil
	//c.JSON(http.StatusOK, gin.H{"tokenInfo": tokenInfo})
	//return tokenInfo.user_ID
}


