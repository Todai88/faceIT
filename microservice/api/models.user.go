package api

import (
	"strings"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	NickName  string `json:"nickname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Country   string `json:"country"`
	ID        int    `json:"ID"`
}

/*
	A slice is used as persistent "storage" as it's quicker to iterate through (marginally on bigger datasets).
	Further, a map was used for quicker lookup capabilities, with the nickname as a key. Reason being that nicknames are unique on the FaceIT platform.
	Realistically speaking a database would be used for all of these capabilities and our main focus would be to get parameters,
	send them to the DB and return the results.
*/
var users = map[int]User{
	1: User{FirstName: "Jane", LastName: "Doe", NickName: "1337", Email: "1337@hltv.org", Password: "FnaticFanGrrl91", Country: "USA", ID: 1},
	2: User{FirstName: "John", LastName: "Doe", NickName: "h4xx0r", Email: "h4xx0r@SKgaming.com", Password: "ILoveGrubby4eva!", Country: "Netherlands", ID: 2},
}

func createNewUser(user User) bool {
	if !doesUserExist(user.ID) {
		users[user.ID] = user
		return true
	} else {
		return false
	}
}

func (u *User) validate() bool {
	return u.FirstName != "" && u.LastName != "" && u.NickName != "" && u.Email != "" && u.Password != "" && u.Country != ""
}

func getSlicedUsers() []User {
	var slicedUsers = make([]User, len(users))
	index := 0
	for _, value := range users {
		slicedUsers[index] = value
		index++
	}
	return slicedUsers
}

func compareEquals(expected, actual string) bool {
	if strings.ToLower(expected) == strings.ToLower(actual) {
		return true
	}
	return false
}

func doesUserExist(ID int) bool {
	if _, ok := users[ID]; ok {
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
