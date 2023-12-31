package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redstonerayy/discounter/db"
)

/*------------ check if code is valid, invalid or non existend ------------*/
func CheckCode(c *gin.Context) {
	/*------------ parse client request ------------*/
	var codeinfo CodeInfo
	if err := c.BindJSON(&codeinfo); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	/*------------ check ------------*/
	codestatus := db.Check(codeinfo.Code)

	/*------------ return status of code ------------*/
	c.JSON(http.StatusOK, gin.H{
		"codestatus": codestatus,
	})
}
