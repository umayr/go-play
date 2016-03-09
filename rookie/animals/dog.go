package animals

import "go-play/rookie/conf"

type Dog struct {
	Name string
}

func (d *Dog) Speak() string {
	return "Type: Dog\nName: " + d.Name
}

func NewDog(src conf.AnimalConfiguration) *Dog {
	d := &Dog{}
	d.Name = src.Name
	return d;
}