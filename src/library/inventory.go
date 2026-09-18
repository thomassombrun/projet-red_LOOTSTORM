package library

import "fmt"

func AccessInventory(c *Character, enemy *Monster) {
	for {
		fmt.Println()
		fmt.Println("===== INVENTAIRE =====")

		if len(c.Inventory) == 0 {
			fmt.Println("L'inventaire est vide.")
		} else {
			for i, item := range c.Inventory {
				fmt.Printf("%d. %s\n", i+1, item.Name)
			}
		}

		fmt.Println("0. Retour")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scanln(&choice)

		if choice == 0 {
			return
		}

		if choice < 1 || choice > len(c.Inventory) {
			fmt.Println("Choix invalide.")
			continue
		}

		item := c.Inventory[choice-1]

		switch item.Name {
		case "Potion de vie":
			TakePot(c, choice-1)

		case "Potion de poison":
			PoisonPot(c, enemy, choice-1)

		case "Amelioration d'inventaire":
			upgradeInventorySlot(c)
			c.Inventory = append(c.Inventory[:choice-1], c.Inventory[choice:]...)

		case "Livre de Sort : Boule de Feu":
			SpellBook(c)
			c.Inventory = append(c.Inventory[:choice-1], c.Inventory[choice:]...)

		default:
			fmt.Println("Cet objet ne peut pas encore être utilisé.")
		}
	}
}

func upgradeInventorySlot(c *Character) {
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

func isInventoryFull(c *Character) bool {
	if len(c.Inventory) >= c.LimitInventory {
		fmt.Println("Votre inventaire est plein !")
		fmt.Printf("Capacité : %d / %d\n", len(c.Inventory), c.LimitInventory)
		return true
	}
	return false
}
