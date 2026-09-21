package library

import "fmt"

func TakePot(c *Character, index int) {
	item := c.Inventory[index]

	if item.Name != "Potion de vie" {
		fmt.Println("Cet objet n'est pas une Potion de vie.")
		return
	}

	if c.CurrentHP >= c.MaxHP {
		fmt.Println("Vous êtes déjà au maximum de vos PV, vous ne pouvez pas vous soigner.")
		return
	}

	healAmount := 50
	missingHP := c.MaxHP - c.CurrentHP
	if missingHP < healAmount {
		healAmount = missingHP
	}

	if item.Quantity > 1 {
		c.Inventory[index].Quantity--
	} else {
		c.Inventory = append(
			c.Inventory[:index],
			c.Inventory[index+1:]...,
		)
	}

	c.CurrentHP += healAmount
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}

	fmt.Println()
	fmt.Println("Vous utilisez une Potion de vie.")
	fmt.Printf("Vous regagnez %d PV.", healAmount)
	fmt.Printf(" PV : %d / %d\n", c.CurrentHP, c.MaxHP)
}

func TakeManaPot(c *Character, index int) {
	item := c.Inventory[index]
	if item.Name != "Potion de mana" {
		fmt.Println("Cet objet n'est pas une Potion de mana.")
		return
	}
	if c.Mana >= c.MaxMana {
		fmt.Println("Votre mana est déjà au maximum.")
		return
	}

	const manaAmount = 30
	c.Mana += manaAmount
	if c.Mana > c.MaxMana {
		c.Mana = c.MaxMana
	}
	removeInventoryQuantity(c, index, 1)
	fmt.Printf("Vous utilisez une Potion de mana et récupérez %d mana. Mana : %d / %d\n", manaAmount, c.Mana, c.MaxMana)
}
