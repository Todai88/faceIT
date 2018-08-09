package api

type User struct {
	FirstName       	string `json:"firstname"`
	LastName     	 	string `json:"lastname"`
	NickName      	 	string `json:"nickname"`
	Email		 		string `json:"email"`
	Password		 	string `json:"password"`
	Country		 		string `json:"country"`
}

var users = map[string]User{
	"1": User{FirstName: "Jane", LastName: "Doe", NickName: "1337",   Email: "1337@hltv.org", Password: "FnaticFanGrrl91", Country: "USA"},
	"2": User{FirstName: "John", LastName: "Doe", NickName: "h4xx0r", Email: "h4xx0r@SKgaming.com", Password: "ILoveGrubby4eva!", Country: "Netherlands"}
}