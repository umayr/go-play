package animals

import "go-play/interfaces/conf"

type Cat struct {
	Name string
}

func (c *Cat) Speak() string {
	return "Type: Cat\nName: " + c.Name
}

func NewCat(src conf.AnimalConfiguration) *Cat {
	c := &Cat{}
	c.Name = src.Name
	return c;
}