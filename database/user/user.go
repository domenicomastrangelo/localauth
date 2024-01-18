package user

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	GroupID  int    `json:"group_id"`
}

func New() *User {
	return &User{}
}
