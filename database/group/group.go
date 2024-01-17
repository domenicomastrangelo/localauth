package group

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func New() *Group {
	return &Group{}
}

func (g *Group) GetName() string {
	return g.Name
}

func (g *Group) SetName(name string) {
	g.Name = name
}

func (g *Group) GetID() int {
	return g.ID
}

func (g *Group) SetID(id int) {
	g.ID = id
}
