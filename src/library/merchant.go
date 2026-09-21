package library

import "fmt"

func MerchantMenu(c *Character) {
	for {
		ClearTerminal()

		fmt.Println()
		fmt.Println("===== MARCHAND =====")
		fmt.Printf("Pièces d'or : %d\n", c.Gold)
		fmt.Println()
		fmt.Println("1. Potion de vie - 3 pièces d'or")
		fmt.Println("2. Potion de poison - 6 pièces d'or")
		fmt.Println("3. Livre de Sort : Boule de Feu - 25 pièces d'or")
		fmt.Println("4. Fourrure de Loup - 4 pièces d'or")
		fmt.Println("5. Peau de Troll - 7 pièces d'or")
		fmt.Println("6. Cuir de Sanglier - 3 pièces d'or")
		fmt.Println("7. Plume de Corbeau - 1 pièce d'or")
		fmt.Println("8. Amelioration d'inventaire - 30 pièces d'or")
		fmt.Println("0. Retour")

		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			BuyItem(c, "Potion de vie", 3)

		case 2:
			BuyItem(c, "Potion de poison", 6)

		case 3:
			for _, skill := range c.Skill {
				if skill.Name == "Boule de Feu" {
					fmt.Println("Vous connaissez déjà le sort Boule de Feu.")
					return
				}
			}
			if !BuyItem(c, "Livre de Sort : Boule de Feu", 25) {
				return
			}

		case 4:
			BuyItem(c, "Fourrure de Loup", 4)

		case 5:
			BuyItem(c, "Peau de Troll", 7)

		case 6:
			BuyItem(c, "Cuir de Sanglier", 3)

		case 7:
			BuyItem(c, "Plume de Corbeau", 1)

		case 8:
			BuyItem(c, "Amelioration d'inventaire", 30)

		case 0:
			fmt.Println("Retour au menu.")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func BuyItem(c *Character, itemName string, price int) bool {
	if IsInventoryFull(c) {
		return false
	}
	if c.Gold < price {
		fmt.Println()
		fmt.Println("Vous n'avez pas assez de pièces d'or.")
		return false
	}

	c.Gold -= price

	c.Inventory = append(
		c.Inventory,
		Item{Name: itemName, Quantity: 1},
	)

	fmt.Println()
	fmt.Printf("Vous avez acheté : %s\n", itemName)
	fmt.Printf("Prix : %d pièces d'or\n", price)
	fmt.Printf("Pièces d'or restantes : %d\n", c.Gold)
	return true
}
