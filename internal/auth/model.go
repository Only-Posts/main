package auth

type Login struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type SignUp struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserData struct {
	ID       uint   `json:"id" gorm:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
