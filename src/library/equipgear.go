package library

import "fmt"

func EquipItem(c *Character, itemName string, index int) {
	var slotType string
	var hpBonus int

	switch itemName {
	case "Chapeau de l'aventurier":
		slotType = "Helmet"
		hpBonus = 10
	case "Tunique de l'aventurier":
		slotType = "Chestplate"
		hpBonus = 25
	case "Bottes de l'aventurier":
		slotType = "Boots"
		hpBonus = 15
	default:
		fmt.Println("Cet objet ne peut pas être équipé.")
		return
	}
	item := &c.Inventory[index]
	item.Quantity--
	if item.Quantity <= 0 {
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
	}
	switch slotType {
	case "Helmet":

		if c.Equip.Helmet != "" && c.Equip.Helmet != "Aucun" {
			c.AddOrMergeItem(c.Equip.Helmet, 1)
			c.MaxHP -= 10
		}
		c.Equip.Helmet = itemName
	case "Chestplate":
		if c.Equip.Chestplate != "" && c.Equip.Chestplate != "Aucun" {
			c.AddOrMergeItem(c.Equip.Chestplate, 1)
			c.MaxHP -= 25
		}
		c.Equip.Chestplate = itemName
	case "Boots":
		if c.Equip.Boots != "" && c.Equip.Boots != "Aucun" {
			c.AddOrMergeItem(c.Equip.Boots, 1)
			c.MaxHP -= 15
		}
		c.Equip.Boots = itemName
	}
	c.MaxHP += hpBonus
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
	fmt.Printf("Vous avez équipé : %s (+%d PV Max)\n", itemName, hpBonus)
}
