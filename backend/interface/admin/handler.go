package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAdmin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "CreateAdmin"})
}

func ListAdmins(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ListAdmins"})
}

func GetAdmin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetAdmin"})
}

func UpdateAdmin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateAdmin"})
}

func DeleteAdmin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteAdmin"})
}
