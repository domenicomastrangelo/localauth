package container

import (
	"localauth/config"
	"localauth/database/grouprepository"
	"localauth/database/user/userrepository"

	"github.com/jackc/pgx/v5"
)

type Container struct {
	Elements map[string]Element
}

type Element struct {
	Element interface{}
}

func New() *Container {
	return &Container{}
}

func (c *Container) GetElement(name string) interface{} {
	return c.Elements[name].Element
}

func (c *Container) AddElement(name string, element Element) {
	if c.Elements == nil {
		c.Elements = make(map[string]Element)
	}

	c.Elements[name] = element
}

func (c *Container) GetDB() *pgx.Conn {
	return c.Elements["db"].Element.(*pgx.Conn)
}

func (c *Container) GetGroupRepository() *grouprepository.RepositoryImpl {
	return c.Elements["grouprepository"].Element.(*grouprepository.RepositoryImpl)
}

func (c *Container) GetUserRepository() *userrepository.RepositoryImpl {
	return c.Elements["userrepository"].Element.(*userrepository.RepositoryImpl)
}

func (c *Container) GetConfigData() *config.ConfigData {
	return c.Elements["conf"].Element.(*config.ConfigData)
}
