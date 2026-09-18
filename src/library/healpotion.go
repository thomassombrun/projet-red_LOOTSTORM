package library

import "fmt"

func TakePot(c *Character, index int) {
	item := c.Inventory[index]

	if item.Name != "Potion de vie" {
		fmt.Println("Cet objet n'est pas une Potion de vie.")
		return
	}

	c.Inventory = append(
		c.Inventory[:index],
		c.Inventory[index+1:]...,
	)

	c.CurrentHP += 50

	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}

	fmt.Println()
	fmt.Println("Vous utilisez une Potion de vie.")
	fmt.Printf("PV : %d / %d\n", c.CurrentHP, c.MaxHP)
}
