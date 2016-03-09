package main

import (
	"fmt"
	"go-play/interfaces/conf"
	"go-play/interfaces/animals"
)

func main() {
	conf := conf.Get()
	animals := []animals.Animal{
		animals.NewCat(conf.Animals["cat"]),
		animals.NewDog(conf.Animals["dog"]),
	}

	for _, v := range animals {
		fmt.Println(v.Speak())
		fmt.Println("--------")
	}
}
