package api

import (
	"encoding/json"
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

func FromJSON(data []byte) User {
	user := User{}
	err := json.Unmarshal(data, &user)
	if err != nil {
		panic(err)
	}
	return user
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}
