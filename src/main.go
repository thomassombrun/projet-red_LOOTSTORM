package main

import (
	"fmt"
	"projet/src/library"
)

func main() {
	c1 := library.InitCharacter("Himiko Toga", "Assassin", 100)
	c2 := library.InitCharacter("Link", "Chevalier", 120)
	c3 := library.InitCharacter("Patrick Bouldefeu", "Mage", 80)

	c1.DisplayInfo()
	fmt.Println("---")
	c2.DisplayInfo()
	fmt.Println("---")
	c3.DisplayInfo()
}
