package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	firstName string `json:"firstname"`
	lastName  string `json:"lastname"`
	nickName  string `json:"nickname"`
	email     string `json:"email"`
	password  string `json:"password"`
	country   string `json:"country"`
}

// Use map for scalability
var users = map[string]User{
	"1": User{firstName: "Jane", lastName: "Doe", nickName: "1337", email: "1337@hltv.org", password: "FnaticFanGrrl91", country: "USA"},
	"2": User{firstName: "John", lastName: "Doe", nickName: "h4xx0r", email: "h4xx0r@SKgaming.com", password: "ILoveGrubby4eva!", country: "Netherlands"},
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
