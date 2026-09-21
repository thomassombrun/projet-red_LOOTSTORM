package library

import "fmt"

func EquipItem(c *Character, itemName string, index int) {
	var slotType string
	var hpBonus int
	var weaponDamage int

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
	case "Dague de l'assassin", "Arc du chasseur", "Marteau du guerrier", "Épée du chevalier", "Lame de gobelin", "Bave de slime":
		slotType = "Weapon"
		weaponDamage = weaponDamageBonus(itemName)
	case "Lame spectrale", "Marteau de golem", "Massue de troll", "Bec du canard", "Épée squelette", "Crocs du loup", "Bâton maudit", "Griffe du dragon":
		slotType = "Weapon"
		weaponDamage = weaponDamageBonus(itemName)
	case "Casque de gobelin", "Carapace de slime", "Capuche spectrale", "Casque de golem", "Peau de troll renforcée", "Plumes du canard", "Heaume squelette", "Fourrure du loup", "Chapeau du sorcier", "Écailles de dragon":
		slotType = "Helmet"
		hpBonus = equipmentHPBonus(itemName)
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
			c.MaxHP -= equipmentHPBonus(c.Equip.Helmet)
		}
		c.Equip.Helmet = itemName
	case "Chestplate":
		if c.Equip.Chestplate != "" && c.Equip.Chestplate != "Aucun" {
			c.AddOrMergeItem(c.Equip.Chestplate, 1)
			c.MaxHP -= equipmentHPBonus(c.Equip.Chestplate)
		}
		c.Equip.Chestplate = itemName
	case "Boots":
		if c.Equip.Boots != "" && c.Equip.Boots != "Aucun" {
			c.AddOrMergeItem(c.Equip.Boots, 1)
			c.MaxHP -= equipmentHPBonus(c.Equip.Boots)
		}
		c.Equip.Boots = itemName
	case "Weapon":
		if c.Equip.Weapon != "" && c.Equip.Weapon != "Aucun" {
			c.AddOrMergeItem(c.Equip.Weapon, 1)
			c.Equip.Weapon = "Aucun"
			c.Equip.WeaponDamage = 0
		}
		c.Equip.Weapon = itemName
		c.Equip.WeaponDamage = weaponDamage
	}
	c.MaxHP += hpBonus
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
	fmt.Printf("Vous avez équipé : %s (+%d PV Max)\n", itemName, hpBonus)
}

func weaponDamageBonus(itemName string) int {
	switch itemName {
	case "Dague de l'assassin", "Bave de slime":
		return 8
	case "Arc du chasseur", "Lame de gobelin":
		return 10
	case "Lame spectrale", "Bâton maudit":
		return 14
	case "Marteau du guerrier":
		return 12
	case "Épée du chevalier":
		return 15
	case "Marteau de golem":
		return 18
	case "Massue de troll":
		return 16
	case "Bec du canard":
		return 14
	case "Épée squelette":
		return 12
	case "Crocs du loup":
		return 9
	case "Griffe du dragon":
		return 25
	default:
		return 0
	}
}

func equipmentHPBonus(itemName string) int {
	switch itemName {
	case "Chapeau de l'aventurier":
		return 10
	case "Tunique de l'aventurier":
		return 25
	case "Bottes de l'aventurier":
		return 15
	case "Casque de gobelin":
		return 15
	case "Carapace de slime":
		return 20
	case "Capuche spectrale":
		return 25
	case "Casque de golem":
		return 30
	case "Peau de troll renforcée":
		return 35
	case "Plumes du canard":
		return 20
	case "Heaume squelette":
		return 18
	case "Fourrure du loup":
		return 12
	case "Chapeau du sorcier":
		return 15
	case "Écailles de dragon":
		return 50
	default:
		return 0
	}
}

func UnequipWeapon(c *Character) {
	if c.Equip.Weapon == "" || c.Equip.Weapon == "Aucun" {
		fmt.Println("Aucune arme n'est équipée.")
		return
	}

	weapon := c.Equip.Weapon
	c.AddOrMergeItem(weapon, 1)
	c.Equip.Weapon = "Aucun"
	c.Equip.WeaponDamage = 0
	fmt.Printf("Vous avez déséquipé : %s.\n", weapon)
}
