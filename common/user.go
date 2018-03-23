package common

// User is a user object
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
