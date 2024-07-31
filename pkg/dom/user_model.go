package dom

// user model store the user details
type User struct {
	UserName string
	Email    string
	Password string
	Role     string
}

type Login struct {
	Email    string
	Password string
}
