package library

import "fmt"

func poisonPot(c *Character, enemy *Character, index int) {
	item := c.Inventory[index]

	if item.Name != "Potion de poison" {
		fmt.Println("Cet objet n'est pas une Potion de poison.")
		return
	}

	c.Inventory = append(
		c.Inventory[:index],
		c.Inventory[index+1:]...,
	)

	enemy.PoisonTurns = 3

	fmt.Println()
	fmt.Printf("%s utilise une Potion de poison sur %s.\n",
		c.Name,
		enemy.Name,
	)

	fmt.Printf("%s est empoisonné pendant 3 tours.\n",
		enemy.Name,
	)
}
