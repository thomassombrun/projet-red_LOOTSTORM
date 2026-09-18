package library

import "fmt"

func PoisonPot(c *Character, enemy *Monster, index int) {
	c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
	enemy.PoisonTurns = 3

	fmt.Println()
	fmt.Printf("%s utilise une Potion de poison sur %s.\n", c.Name, enemy.Name)
	fmt.Printf("%s est empoisonné pendant 3 tours.\n", enemy.Name)
}
