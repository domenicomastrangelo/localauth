package user

type Service interface {
	Authorize(user *User) error
}
