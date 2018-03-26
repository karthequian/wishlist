package common

// User is a user object
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

//Userlist is the list of users
var Userlist = []User{
	User{
		Name:     "Karthik",
		Password: "helloworld",
		Username: "karthik",
		Token:    "20d15fed-42f4-4a71-b9a3-7c7fee78d38d",
	},
	User{
		Name:     "Tracey",
		Password: "helloworld",
		Username: "tracey",
		Token:    "4d76c945-a946-4d2b-95a8-281aff55404f",
	},
	User{
		Name:     "Carisa",
		Password: "helloworld",
		Username: "carisa",
		Token:    "224b3200-d09b-4881-8a7b-d69d6d8ba543",
	},
	User{
		Name:     "Ernest",
		Password: "helloworld",
		Username: "ernest",
		Token:    "2115274e-34bc-4456-a1ee-1c4c171231a9",
	},
	User{
		Name:     "Amy",
		Password: "helloworld",
		Username: "amy",
		Token:    "5343b1d3-dfa3-4823-b544-e5907c3585f5",
	},
}
