package library

import "fmt"

func PoisonPot(c *Character, enemy *Monster, index int) {
	c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
	enemy.PoisonTurns = 3

	fmt.Println()
	fmt.Printf("%s utilise une Potion de poison sur %s.\n", c.Name, enemy.Name)
	fmt.Printf("%s est empoisonné pendant 3 tours.\n", enemy.Name)
}

func PoisonPotSilent(c *Character, enemy *Monster, index int) bool {
	if enemy == nil || index < 0 || index >= len(c.Inventory) || c.Inventory[index].Name != "Potion de poison" {
		return false
	}
	c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
	enemy.PoisonTurns = 3
	return true
}
