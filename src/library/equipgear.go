package library

import "fmt"

func currentEquippedForSlot(c *Character, slotType string) string {
	switch slotType {
	case "Helmet":
		return c.Equip.Helmet
	case "Chestplate":
		return c.Equip.Chestplate
	case "Boots":
		return c.Equip.Boots
	case "Weapon":
		return c.Equip.Weapon
	default:
		return "Aucun"
	}
}

func EquipItem(c *Character, itemName string, index int) {
	var slotType string
	var hpBonus int
	var weaponDamage int
	var replaceExisting bool

	normalizedName := normalizeItemName(itemName)
	itemRarity := parseItemRarity(itemName)
	baseName := normalizedName

	switch baseName {
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
		weaponDamage = weaponDamageBonus(baseName)
	case "Lame spectrale", "Massue de troll", "Bec du canard", "Épée squelette", "Crocs du loup", "Bâton maudit", "Griffe du dragon":
		slotType = "Weapon"
		weaponDamage = weaponDamageBonus(baseName)
	case "Marteau de golem":
		slotType = "Weapon"
		weaponDamage = weaponDamageBonus(baseName)
	case "Casque de gobelin", "Carapace de slime", "Capuche spectrale", "Casque de golem", "Peau de troll renforcée", "Plumes du canard", "Heaume squelette", "Fourrure du loup", "Chapeau du sorcier", "Écailles de dragon":
		slotType = "Helmet"
		hpBonus = equipmentHPBonus(baseName)
	case "Tunique de gobelin", "Tunique de slime", "Tunique spectrale", "Tunique de golem", "Tunique de troll", "Tunique du canard", "Tunique du squelette", "Tunique du loup", "Tunique du sorcier", "Tunique du dragon":
		slotType = "Chestplate"
		hpBonus = equipmentHPBonus(baseName)
	case "Bottes de gobelin", "Bottes de slime", "Bottes spectrales", "Bottes de golem", "Bottes de troll", "Bottes du canard", "Bottes du squelette", "Bottes du loup", "Bottes du sorcier", "Bottes du dragon":
		slotType = "Boots"
		hpBonus = equipmentHPBonus(baseName)
	default:
		fmt.Println("Cet objet ne peut pas être équipé.")
		return
	}
	currentEquipped := currentEquippedForSlot(c, slotType)
	if currentEquipped != "" && currentEquipped != "Aucun" && normalizeItemName(currentEquipped) != baseName {
		fmt.Println("Vous avez déjà un équipement dans cette catégorie :")
		fmt.Printf("- Actuel : %s%s\n", currentEquipped, itemStatSummary(currentEquipped))
		fmt.Printf("- Nouveau : %s%s\n", equipmentDisplayName(baseName, itemRarity), itemStatSummary(itemName))
		fmt.Println("1. Équiper le nouvel objet et remettre l'ancien dans l'inventaire")
		fmt.Println("2. Garder l'ancien équipement")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scanln(&choice)
		if choice == 2 {
			if c.Inventory[index].Quantity <= 0 {
				c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
			}
			return
		}
		replaceExisting = true
	}

	item := &c.Inventory[index]
	item.Quantity--
	if item.Quantity <= 0 {
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
	}

	if slotType == "Weapon" {
		weaponDamage = scaledEquipmentStat(weaponDamage, itemRarity)
	}
	if slotType != "Weapon" {
		hpBonus = scaledEquipmentStat(hpBonus, itemRarity)
	}

	switch slotType {
	case "Helmet":
		if replaceExisting && c.Equip.Helmet != "" && c.Equip.Helmet != "Aucun" {
			c.AddOrMergeItem(c.Equip.Helmet, 1)
			c.MaxHP -= scaledEquipmentStat(equipmentHPBonus(c.Equip.Helmet), parseItemRarity(c.Equip.Helmet))
		}
		c.Equip.Helmet = equipmentDisplayName(baseName, itemRarity)
	case "Chestplate":
		if replaceExisting && c.Equip.Chestplate != "" && c.Equip.Chestplate != "Aucun" {
			c.AddOrMergeItem(c.Equip.Chestplate, 1)
			c.MaxHP -= scaledEquipmentStat(equipmentHPBonus(c.Equip.Chestplate), parseItemRarity(c.Equip.Chestplate))
		}
		c.Equip.Chestplate = equipmentDisplayName(baseName, itemRarity)
	case "Boots":
		if replaceExisting && c.Equip.Boots != "" && c.Equip.Boots != "Aucun" {
			c.AddOrMergeItem(c.Equip.Boots, 1)
			c.MaxHP -= scaledEquipmentStat(equipmentHPBonus(c.Equip.Boots), parseItemRarity(c.Equip.Boots))
		}
		c.Equip.Boots = equipmentDisplayName(baseName, itemRarity)
	case "Weapon":
		if replaceExisting && c.Equip.Weapon != "" && c.Equip.Weapon != "Aucun" {
			c.AddOrMergeItem(c.Equip.Weapon, 1)
			c.Equip.Weapon = "Aucun"
			c.Equip.WeaponDamage = 0
		}
		c.Equip.Weapon = equipmentDisplayName(baseName, itemRarity)
		c.Equip.WeaponDamage = weaponDamage
	}
	c.MaxHP += hpBonus
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
	fmt.Printf("Vous avez équipé : %s (+%d PV Max)\n", equipmentDisplayName(baseName, itemRarity), hpBonus)
}

func weaponDamageBonus(itemName string) int {
	base := normalizeItemName(itemName)
	switch base {
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
	base := normalizeItemName(itemName)
	switch base {
	case "Chapeau de l'aventurier":
		return 10
	case "Tunique de l'aventurier":
		return 25
	case "Bottes de l'aventurier":
		return 15
	case "Casque de gobelin":
		return 15
	case "Tunique de gobelin":
		return 22
	case "Bottes de gobelin":
		return 12
	case "Carapace de slime":
		return 20
	case "Tunique de slime":
		return 25
	case "Bottes de slime":
		return 15
	case "Capuche spectrale":
		return 25
	case "Tunique spectrale":
		return 30
	case "Bottes spectrales":
		return 18
	case "Casque de golem":
		return 30
	case "Tunique de golem":
		return 35
	case "Bottes de golem":
		return 20
	case "Peau de troll renforcée":
		return 35
	case "Tunique de troll":
		return 40
	case "Bottes de troll":
		return 22
	case "Plumes du canard":
		return 20
	case "Tunique du canard":
		return 25
	case "Bottes du canard":
		return 15
	case "Heaume squelette":
		return 18
	case "Tunique du squelette":
		return 22
	case "Bottes du squelette":
		return 14
	case "Fourrure du loup":
		return 12
	case "Tunique du loup":
		return 18
	case "Bottes du loup":
		return 10
	case "Chapeau du sorcier":
		return 15
	case "Tunique du sorcier":
		return 20
	case "Bottes du sorcier":
		return 12
	case "Écailles de dragon":
		return 50
	case "Tunique du dragon":
		return 60
	case "Bottes du dragon":
		return 30
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
