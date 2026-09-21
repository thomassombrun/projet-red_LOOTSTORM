package library

import "fmt"

func EquipItem(c *Character, itemName string, index int) {
	var slotType string
	var hpBonus int
	var weaponDamage int

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
	case "Lame spectrale", "Marteau de golem", "Massue de troll", "Bec du canard", "Épée squelette", "Crocs du loup", "Bâton maudit", "Griffe du dragon":
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
	item := &c.Inventory[index]
	item.Quantity--
	if item.Quantity <= 0 {
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
	}

	baseStatsMultiplier := rarityMultiplier(itemRarity)
	if slotType == "Weapon" {
		weaponDamage = int(float64(weaponDamage) * baseStatsMultiplier)
	}
	if slotType != "Weapon" {
		hpBonus = int(float64(hpBonus) * baseStatsMultiplier)
	}

	switch slotType {
	case "Helmet":
		if c.Equip.Helmet != "" && c.Equip.Helmet != "Aucun" {
			c.AddOrMergeItem(c.Equip.Helmet, 1)
			c.MaxHP -= equipmentHPBonus(normalizeItemName(c.Equip.Helmet))
		}
		c.Equip.Helmet = equipmentDisplayName(baseName, itemRarity)
	case "Chestplate":
		if c.Equip.Chestplate != "" && c.Equip.Chestplate != "Aucun" {
			c.AddOrMergeItem(c.Equip.Chestplate, 1)
			c.MaxHP -= equipmentHPBonus(normalizeItemName(c.Equip.Chestplate))
		}
		c.Equip.Chestplate = equipmentDisplayName(baseName, itemRarity)
	case "Boots":
		if c.Equip.Boots != "" && c.Equip.Boots != "Aucun" {
			c.AddOrMergeItem(c.Equip.Boots, 1)
			c.MaxHP -= equipmentHPBonus(normalizeItemName(c.Equip.Boots))
		}
		c.Equip.Boots = equipmentDisplayName(baseName, itemRarity)
	case "Weapon":
		if c.Equip.Weapon != "" && c.Equip.Weapon != "Aucun" {
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
