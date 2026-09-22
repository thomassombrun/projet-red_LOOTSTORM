package library

import "fmt"

func AccessInventory(c *Character, enemy *Monster) bool {
	return accessInventory(c, enemy, false)
}

func AccessCombatInventory(c *Character, enemy *Monster) bool {
	return accessInventory(c, enemy, true)
}

func accessInventory(c *Character, enemy *Monster, combatOnly bool) bool {
	for {
		ClearTerminal()

		fmt.Println()
		fmt.Println("===== INVENTAIRE =====")

		availableItems := make([]int, 0, len(c.Inventory))
		for i, item := range c.Inventory {
			if !combatOnly || item.Name == "Potion de vie" || item.Name == "Potion de poison" || item.Name == "Potion de mana" {
				availableItems = append(availableItems, i)
			}
		}

		if len(availableItems) == 0 {
			fmt.Println("L'inventaire est vide.")
		} else {
			for displayIndex, inventoryIndex := range availableItems {
				item := c.Inventory[inventoryIndex]
				label := itemDisplayName(item)
				bonusText := itemStatSummary(item.Name)
				if item.Quantity > 1 {
					fmt.Printf("%d. %s%s x%d\n", displayIndex+1, label, bonusText, item.Quantity)
				} else {
					fmt.Printf("%d. %s%s\n", displayIndex+1, label, bonusText)
				}
			}
		}

		fmt.Println("0. Retour")
		if !combatOnly {
			fmt.Printf("%d. Jeter un objet\n", len(availableItems)+1)
			fmt.Printf("%d. Déséquiper l'arme\n", len(availableItems)+2)
		}
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scanln(&choice)

		if choice == 0 {
			return false
		}
		if !combatOnly && choice == len(availableItems)+1 {
			discardItem(c)
			return true
		}
		if !combatOnly && choice == len(availableItems)+2 {
			UnequipWeapon(c)
			WaitForEnter()
			return true
		}

		if choice < 1 || choice > len(availableItems) {
			fmt.Println("Choix invalide.")
			continue
		}

		inventoryIndex := availableItems[choice-1]
		item := c.Inventory[inventoryIndex]

		switch item.Name {
		case "Potion de vie":
			TakePot(c, inventoryIndex)
			if !combatOnly {
				WaitForEnter()
			}
			return true

		case "Potion de mana":
			TakeManaPot(c, inventoryIndex)
			if !combatOnly {
				WaitForEnter()
			}
			return true

		case "Potion de poison":
			if enemy == nil {
				fmt.Println("La Potion de poison ne peut être utilisée que pendant un combat.")
				continue
			}
			PoisonPot(c, enemy, inventoryIndex)
			if !combatOnly {
				WaitForEnter()
			}
			return true

		case "Amelioration d'inventaire":
			UpgradeInventorySlot(c)
			c.Inventory = append(c.Inventory[:inventoryIndex], c.Inventory[inventoryIndex+1:]...)
			WaitForEnter()
			return true

		case "Livre de Sort : Boule de Feu", "Livre de Sort : Soin", "Livre de Sort : Régénération", "Livre de Sort : Poison", "Livre de Sort : Brûlure", "Livre de Sort : Éclair", "Livre de Sort : Barrière Sacrée":
			LearnSpellBook(c, item.Name)
			c.Inventory = append(c.Inventory[:inventoryIndex], c.Inventory[inventoryIndex+1:]...)
			WaitForEnter()
			return true

		case "Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier", "Casque de gobelin", "Carapace de slime", "Capuche spectrale", "Casque de golem", "Peau de troll renforcée", "Plumes du canard", "Heaume squelette", "Fourrure du loup", "Chapeau du sorcier", "Écailles de dragon", "Dague de l'assassin", "Arc du chasseur", "Marteau du guerrier", "Épée du chevalier", "Lame de gobelin", "Bave de slime", "Lame spectrale", "Marteau de golem", "Massue de troll", "Bec du canard", "Épée squelette", "Crocs du loup", "Bâton maudit", "Griffe du dragon", "Tunique de gobelin", "Bottes de gobelin", "Tunique de slime", "Bottes de slime", "Tunique spectrale", "Bottes spectrales", "Tunique de golem", "Bottes de golem", "Tunique de troll", "Bottes de troll", "Tunique du canard", "Bottes du canard", "Tunique du squelette", "Bottes du squelette", "Tunique du loup", "Bottes du loup", "Tunique du sorcier", "Bottes du sorcier", "Tunique du dragon", "Bottes du dragon":
			EquipItem(c, itemDisplayName(item), inventoryIndex)
			WaitForEnter()
			return true

		default:
			fmt.Println("Cet objet ne peut pas encore être utilisé.")
		}
	}
}

func UpgradeInventorySlot(c *Character) {
	if c.LimitInventoryUpgrade >= 3 {
		fmt.Println("Vous avez déjà utilisé les 3 améliorations d'inventaire disponibles.")
		return
	}
	c.LimitInventory += 10
	c.LimitInventoryUpgrade++

	fmt.Println("Votre capacité d'inventaire augmente de 10.")
	fmt.Printf("Capacité maximale : %d\n", c.LimitInventory)
	fmt.Printf("Améliorations utilisées : %d / 3\n", c.LimitInventoryUpgrade)
}

func UpgradeInventorySlotSilent(c *Character) bool {
	if c.LimitInventoryUpgrade >= 3 {
		return false
	}
	c.LimitInventory += 10
	c.LimitInventoryUpgrade++
	return true
}

func IsInventoryFull(c *Character) bool {
	if len(c.Inventory) >= c.LimitInventory {
		fmt.Println("Votre inventaire est plein !")
		fmt.Printf("Capacité : %d / %d\n", len(c.Inventory), c.LimitInventory)
		return true
	}
	return false
}

func (c *Character) AddOrMergeItem(itemName string, quantity int) {
	baseName := normalizeItemName(itemName)
	itemRarity := parseItemRarity(itemName)
	for i := range c.Inventory {
		existingRarity := c.Inventory[i].Rarity
		if existingRarity == "" {
			existingRarity = RarityCommon
		}
		if normalizeItemName(c.Inventory[i].Name) == baseName && existingRarity == itemRarity {
			c.Inventory[i].Rarity = existingRarity
			c.Inventory[i].Quantity += quantity
			return
		}
	}
	if len(c.Inventory) >= c.LimitInventory {
		fmt.Println("Votre inventaire est plein ! L'ancien équipement n'a pas pu y être replacé.")
		return
	}
	c.Inventory = append(c.Inventory, Item{Name: baseName, Quantity: quantity, Rarity: itemRarity})
}

func discardItem(c *Character) {
	if len(c.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		WaitForEnter()
		return
	}

	ClearTerminal()
	fmt.Println("===== OBJETS À JETER =====")
	for i, item := range c.Inventory {
		fmt.Printf("%d. %s x%d\n", i+1, item.Name, item.Quantity)
	}
	fmt.Println("0. Annuler")
	fmt.Print("Choisissez l'objet à jeter : ")

	var choice int
	fmt.Scanln(&choice)
	if choice == 0 {
		return
	}
	if choice < 1 || choice > len(c.Inventory) {
		fmt.Println("Choix invalide.")
		WaitForEnter()
		return
	}

	item := c.Inventory[choice-1]
	removeInventoryQuantity(c, choice-1, 1)
	fmt.Printf("Vous avez jeté 1 %s.\n", item.Name)
	WaitForEnter()
}

func removeInventoryQuantity(c *Character, index int, quantity int) {
	c.Inventory[index].Quantity -= quantity
	if c.Inventory[index].Quantity <= 0 {
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
	}
}
