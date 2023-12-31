package main

import (
	"fmt"
	"github.com/michaels1337/puppy"
	"github.com/michaels1337/puppy/kitty"
)

func main() {
	s := puppy.Bark()
	s1 := puppy.Barks()
	fmt.Println(s)
	fmt.Println(s1)

	// also
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())

	//
	fmt.Println(kitty.Meow())
	fmt.Println(puppy.BigBark())
	fmt.Println(puppy.BigBarks())

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// Version 1.0.1

}
