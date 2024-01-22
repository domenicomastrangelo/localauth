package userservice

import (
	"context"
	"localauth/config"
	"localauth/container"
	"localauth/database"
	"localauth/database/user"
	"localauth/database/user/userrepository"
	"testing"
)

func TestAddUser(t *testing.T) {
	conf := config.New(config.ENV_TEST)

	cont := container.New()
	cont.AddElement("conf", container.Element{
		Element: conf,
	})

	db := database.New(cont)

	cont.AddElement("db", container.Element{
		Element: db,
	})

	var (
		user        = user.New()
		userService = userrepository.New(db)
	)

	user.Name = "John"
	user.Surname = "Doe"
	user.Email = "john.doe@example.com"

	err := userService.AddUser(user, context.Background())
	if err != nil {
		t.Error(err)
	} else if user.ID == 0 {
		t.Error("User not added")
	}

	t.Cleanup(func() {
		db.Exec(context.Background(), "DELETE FROM users WHERE id = $1", user.ID)
	})
}
