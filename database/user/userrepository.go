package user

type Repository interface {
	GetUsers() ([]*User, error)
	GetUser(id int) (*User, error)
	AddUser(group *User) error
	EditUser(group *User) error
}

type RepositoryImpl struct {
}

func (g *RepositoryImpl) GetUsers() ([]*User, error) {
	return nil, nil
}

func (g *RepositoryImpl) GetUser(id int) (*User, error) {
	return nil, nil
}

func (g *RepositoryImpl) AddUser(user *User) error {
	return nil
}

func (g *RepositoryImpl) EditUser(user *User) error {
	return nil
}
