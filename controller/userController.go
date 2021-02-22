package controller

import (
	"hw/dal"
	models "hw/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	users []models.User
)

func init() {
	users = models.Users
}

//HttpGet user
func GetUserAction(c *gin.Context) {
	dal.Gets()
	c.JSON(http.StatusOK, users)
}

//HttpGet /user/:id
func GetUserByIdAction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	for i := 0; i < len(users); i++ {
		if users[i].ID == id {
			c.JSON(http.StatusOK, users[i])
			break
		}
	}
}

//HttpPost /user
func InsertUserAction(c *gin.Context) {
	var u models.User
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	users = append(users, u)
	c.Status(http.StatusOK)
}

//HttpPut /user/:id
func InserActionById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var u models.User
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	for i := 0; i < len(users); i++ {
		if users[i].ID == id {
			users[i].Name = u.Name
			break
		}
	}
	c.Status(http.StatusNoContent)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[0:i], users[i+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}
