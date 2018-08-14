FaceIT Microservice for users

The service is split up into two a main compoent and an API package, that can be loaded using `go get github.com/Todai88/faceIt/microservice/api`. 

The user entity consists of
- [x] First name
- [x] Last name
- [x] Nickname
- [x] Password
- [x] email
- [x] country

The service must allow to:
- [x] add a new user (`POST => /api/v1/users/`)
- [x] modify an existing user (`PUT => /api/v1/users/:id`)
- [x] remove a user (`DELETE => /api/v1/users/:id`)
- [x] return the list of the users satisfying certain criteria (e.g. for country) (`GET => /api/v1/users/?query=value?query2=value2`...)

The microservice will be part of a more complex architecture, so consider for example that
the Search microservice will need to be notified when a new user is added, or that the
Competition microservice will need to be notified when the user changes his nickname.

`//=> These above examples should be possible by using a event / messaging service such as kafka / rabbitMQ. However, this seemed as a bit superflous to be integrated for a programming test.`

Think at how to implement a system that is scalable.

The application must be a “good citizen”:
-[x] Meaningful logs
-[x] Self-documented end points
-[x] Health checks

You can “mock” the database if you like (e.g. saving the data in memory).
Please provide the instructions to start the application on localhost.
We expect to be able to run unit tests and to be able to add/modify/delete/list the users by
calling the end points using http calls after starting the application.
Please explain what are the criteria and assumptions you used to take decisions. A clear
and correct explanation is part of the test.

`// => For the tests, check /microservice/api/user_test.go`

`// => Documentation and reasoning is provided in the code.`

To run application:

Either clone repo and run the above mentioned go get command, then run the microservice/main.go

Or run docker-compose up in the root of the folder. This should create a container with 5050:5050 exposed and available for testing of its endpoints.

The endpoints are self-explanatory and available using localhost:5050/api/v1/users + correct verb + http request body / query params.