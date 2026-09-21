package library

import "fmt"

func AccessInventory(c *Character, enemy *Monster) bool {
	for {
		fmt.Println()
		fmt.Println("===== INVENTAIRE =====")

		if len(c.Inventory) == 0 {
			fmt.Println("L'inventaire est vide.")
		} else {
			for i, item := range c.Inventory {
				if item.Quantity > 1 {
					fmt.Printf("%d. %s x%d\n", i+1, item.Name, item.Quantity)
				} else {
					fmt.Printf("%d. %s\n", i+1, item.Name)
				}
			}
		}

		fmt.Println("0. Retour")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scanln(&choice)

		if choice == 0 {
			return false
		}

		if choice < 1 || choice > len(c.Inventory) {
			fmt.Println("Choix invalide.")
			continue
		}

		item := c.Inventory[choice-1]

		switch item.Name {
		case "Potion de vie":
			TakePot(c, choice-1)
			return true

		case "Potion de poison":
			if enemy == nil {
				fmt.Println("La Potion de poison ne peut être utilisée que pendant un combat.")
				continue
			}
			PoisonPot(c, enemy, choice-1)
			return true

		case "Amelioration d'inventaire":
			UpgradeInventorySlot(c)
			c.Inventory = append(c.Inventory[:choice-1], c.Inventory[choice:]...)
			return true

		case "Livre de Sort : Boule de Feu":
			SpellBook(c)
			c.Inventory = append(c.Inventory[:choice-1], c.Inventory[choice:]...)
			return true

		case "Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier":
			EquipItem(c, item.Name, choice-1)
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

func IsInventoryFull(c *Character) bool {
	if len(c.Inventory) >= c.LimitInventory {
		fmt.Println("Votre inventaire est plein !")
		fmt.Printf("Capacité : %d / %d\n", len(c.Inventory), c.LimitInventory)
		return true
	}
	return false
}

func (c *Character) AddOrMergeItem(itemName string, quantity int) {
	if len(c.Inventory) >= c.LimitInventory {
		fmt.Println("Votre inventaire est plein ! L'ancien équipement n'a pas pu y être replacé.")
		return
	}
	for i := range c.Inventory {
		if c.Inventory[i].Name == itemName {
			c.Inventory[i].Quantity += quantity
			return
		}
	}
	c.Inventory = append(c.Inventory, Item{Name: itemName, Quantity: quantity})
}
