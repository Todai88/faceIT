package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
			if createNewUser(user) {
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": getSlicedUsers()})
			} else {
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "A user with that nickname already exists, unable to add new user"})
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusBadRequest, "data": "The body of your request can not be parsed into a user."})
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusBadRequest, "data": "Something went wrong with your request. Contact support@faceit.com for more information."})
	}
}

/*
	PUT: UpdateUser
	Returns: The updated user and a status code, or a status code and error message.
	Logic:   Finds and updates a user, also notifies the competition microservice.
*/
func UpdateUser(c *gin.Context) {
	if id := c.Param("id"); id == "0" {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusBadRequest, "data": "You didn't provide an id to update with."})
		return
	}

	var user User
	if err := c.BindJSON(&user); err == nil {
		if user.validate() {
			if doesUserExist(user.ID) {
				users[user.ID] = user
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
			} else {
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "No user exists with that nickname, unable to update"})
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusBadRequest, "data": "The body of your request can not be parsed into a user."})
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusBadRequest, "data": "Something went wrong with your request. Contact support@faceit.com for more information."})
	}
}

/*
	DeleteUser
	Returns: A status code and a message.
	Logic:   Deletes a user, also notifies the search microservice
*/
func DeleteUser(c *gin.Context) {
	if id := c.Param("id"); id == "0" {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusBadRequest, "data": "You didn't provide an id to update with."})
		return
	}
	var user User
	if err := c.BindJSON(&user); err == nil {
		if user.validate() {
			if doesUserExist(user.ID) {
				delete(users, user.ID)
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": getSlicedUsers()})
			} else {
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "No user exists with that ID, unable to proceed with delete."})
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusBadRequest, "data": "The body of your request can not be parsed into a user."})
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusBadRequest, "data": "Something went wrong with your request. Contact support@faceit.com for more information."})
	}
}
