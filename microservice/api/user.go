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

var users = []User{
	User{FirstName: "Jane", LastName: "Doe", NickName: "1337", Email: "1337@hltv.org", Password: "FnaticFanGrrl91", Country: "USA"},
	User{FirstName: "John", LastName: "Doe", NickName: "h4xx0r", Email: "h4xx0r@SKgaming.com", Password: "ILoveGrubby4eva!", Country: "Netherlands"},
}

func (u User) ToJSON() []byte {
	ToJSON, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

//TODO: Add fan out -> fan in.
func filterUsers(parameters map[string][]string) []User {
	var tmpUsers = []User{}
	for index := range users {
		/*
			Check if parameter exists in the query string.
			If it does, check if the parameter's value matches the current user's corresponding value.
			If it does, append it to our temporary list.
		*/
		fullFillsAllFilters := true

		/*
			context.Request.URL.Query() returns a map of a list of string.
			In this case we are really only interested in the first item in that list of strings.
			If a user sends two or more country parameters, it can be considered to be an incorrect
			usage of the API, in which case we discard any superflous values.
		*/
		if nickname, ok := parameters["nickname"]; ok {
			if users[index].NickName != nickname[0] {
				fullFillsAllFilters = false
			}
		}
		if country, ok := parameters["country"]; ok {
			if users[index].Country != country[0] {
				fullFillsAllFilters = false
			}
		}
		if firstname, ok := parameters["firstname"]; ok {
			if users[index].FirstName != firstname[0] {
				fullFillsAllFilters = false
			}
		}
		if lastname, ok := parameters["lastname"]; ok {
			if users[index].LastName != lastname[0] {
				fullFillsAllFilters = false
			}
		}
		if email, ok := parameters["email"]; ok {
			if users[index].Email != email[0] {
				fullFillsAllFilters = false
			}
		}
		if fullFillsAllFilters {
			tmpUsers = append(tmpUsers, users[index])
		}
	}
	return tmpUsers
}

func GetUsers(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	if len(queryParams) >= 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": filterUsers(queryParams)})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
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
