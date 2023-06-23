package main

import (
	"fmt"
	"github.com/nomadcoders/nomadcoin/person"
)

func main() {
	nico := person.Person{}
	nico.SetDetails("Nico", 12)

	fmt.Println("Main's nico", nico)

}
