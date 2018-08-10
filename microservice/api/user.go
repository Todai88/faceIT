package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	NickName  string `json:"nickname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Country   string `json:"country"`
}

// Use map for scalability
var users = map[string]User{
	"1": User{FirstName: "Jane", LastName: "Doe", NickName: "1337", Email: "1337@hltv.org", Password: "FnaticFanGrrl91", Country: "USA"},
	"2": User{FirstName: "John", LastName: "Doe", NickName: "h4xx0r", Email: "h4xx0r@SKgaming.com", Password: "ILoveGrubby4eva!", Country: "Netherlands"},
}

func (u User) ToJSON() []byte {
	ToJSON, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

func GetUsers(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	fmt.Println(queryParams)
	if nickname, ok := queryParams[queryParams.Get("nickname")]; ok {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users[nickname[0]]})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
	}
}

func GetUser(c *gin.Context) {
	if user, ok := users[c.Param("id")]; ok {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": make([]User, 0)})
	}
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}
