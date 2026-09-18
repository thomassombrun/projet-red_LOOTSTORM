package library

import "fmt"

func AccessInventory(c *Character) {
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

		default:
			fmt.Println("Cet objet ne peut pas encore être utilisé.")
		}
	}
}
