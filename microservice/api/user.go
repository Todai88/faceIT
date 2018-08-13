package api

import (
	"net/http"
	"strings"

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

/*
	A slice is used as persistent "storage" as it's quicker to iterate through (marginally on bigger datasets).
	Further, a map was used for quicker lookup capabilities, with the nickname as a key. Reason being that nicknames are unique on the FaceIT platform.
	Realistically speaking a database would be used for all of these capabilities and our main focus would be to get parameters,
	send them to the DB and return the results.
*/
var users = map[string]User{
	"1337":   User{FirstName: "Jane", LastName: "Doe", NickName: "1337", Email: "1337@hltv.org", Password: "FnaticFanGrrl91", Country: "USA"},
	"h4xx0r": User{FirstName: "John", LastName: "Doe", NickName: "h4xx0r", Email: "h4xx0r@SKgaming.com", Password: "ILoveGrubby4eva!", Country: "Netherlands"},
}

func (u *User) validate() bool {
	if u.FirstName == "" {
		return false
	}
	if u.LastName == "" {
		return false
	}
	if u.NickName == "" {
		return false
	}
	if u.Email == "" {
		return false
	}
	if u.Password == "" {
		return false
	}
	if u.Country == "" {
		return false
	}
	return true
}

func getSlicedUsers() []User {
	var slicedUsers = make([]User, len(users))
	for index := range users {
		slicedUsers = append(slicedUsers, users[index])
	}
	return slicedUsers
}

func compareEquals(expected, actual string) bool {
	if strings.ToLower(expected) == strings.ToLower(actual) {
		return true
	}
	return false
}

func filterUsers(parameters map[string][]string) []User {
	var slicedUsers = []User{}
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
			Further work on this could include a case-insensitive key in the dictionary.
		*/
		if nickname, ok := parameters["nickname"]; ok {
			if compareEquals(users[index].NickName, nickname[0]) {
				fullFillsAllFilters = false
			}
		}
		if country, ok := parameters["country"]; ok {
			if compareEquals(users[index].Country, country[0]) {
				fullFillsAllFilters = false
			}
		}
		if firstname, ok := parameters["firstname"]; ok {
			if compareEquals(users[index].FirstName, firstname[0]) {
				fullFillsAllFilters = false
			}
		}
		if lastname, ok := parameters["lastname"]; ok {
			if compareEquals(users[index].LastName, lastname[0]) {
				fullFillsAllFilters = false
			}
		}
		if email, ok := parameters["email"]; ok {
			if compareEquals(users[index].Email, email[0]) {
				fullFillsAllFilters = false
			}
		}
		if fullFillsAllFilters {
			slicedUsers = append(slicedUsers, users[index])
		}
	}
	return slicedUsers
}

/*
	GET: GetUsers
	Returns: Slice of users of length >= 0.
	Logic: If there are query parameters, a filter is added.
	Considerations: Would be good to chunk out the slice of users to
					somehow concurrently run over the chunks, increasing performance.
					At the current stage of the application I can't see any reason to add any errors.
					However, if there would be a layer of authentication, I would suggest adding errors to
					handle unauthorized requests.
*/
func GetUsers(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	if len(queryParams) >= 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": filterUsers(queryParams)})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": getSlicedUsers()})
	}
}

/*
	POST: CreateUser
	Returns: The created user, or an error message.
	Logic:   Adds the user to our in-memory slice and adds an event to the queue of RabbitMQ to notify the search microservice.
			 If an entry in users exists with the nickname, then the request will fail.
			 If the request has an erroneous user in it (ie, can't be parsed), it will return an internal server error.
*/
func CreateUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err == nil {
		if user.validate() {
			if _, ok := users[user.NickName]; !ok {
				users[user.NickName] = user
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": getSlicedUsers()})
			} else if ok {
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "A user with that nickname already exists"})
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "data": "Unable to validate your data"})
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "data": "Something went wrong with your request. Contact support@faceit.com for more information."})
	}
}

/*
	UpdateUser
	Returns: The updated user and a status code, or a status code and error message.
	Logic:   Finds and updates a user, also notifies the competition microservice.
*/
func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

/*
	DeleteUser
	Returns: A status code and a message.
	Logic:   Deletes a user, also notifies the search microservice
*/
func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}
