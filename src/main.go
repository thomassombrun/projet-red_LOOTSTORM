package main

import "projet/src/library"

func main() {
	c1 := library.InitCharacter("Himiko Toga", "Assassin", 100)
	c2 := library.InitCharacter("Link", "Chevalier", 120)
	c3 := library.InitCharacter("Patrick Bouldefeu", "Mage", 80)

	_ = c1
	_ = c2
	_ = c3
}
