package library

import "fmt"

func merchantMenu(player *Character) {
	for {
		fmt.Println()
		fmt.Println("===== MARCHAND =====")
		fmt.Printf("Pièces d'or : %d\n", player.Gold)
		fmt.Println()
		fmt.Println("1. Potion de vie - 3 pièces d'or")
		fmt.Println("2. Potion de poison - 6 pièces d'or")
		fmt.Println("3. Livre de Sort : Boule de Feu - 25 pièces d'or")
		fmt.Println("4. Fourrure de Loup - 4 pièces d'or")
		fmt.Println("5. Peau de Troll - 7 pièces d'or")
		fmt.Println("6. Cuir de Sanglier - 3 pièces d'or")
		fmt.Println("7. Plume de Corbeau - 1 pièce d'or")
		fmt.Println("0. Retour")

		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			buyItem(player, "Potion de vie", 3)

		case 2:
			buyItem(player, "Potion de poison", 6)

		case 3:
			buyItem(player, "Livre de Sort : Boule de Feu", 25)

		case 4:
			buyItem(player, "Fourrure de Loup", 4)

		case 5:
			buyItem(player, "Peau de Troll", 7)

		case 6:
			buyItem(player, "Cuir de Sanglier", 3)

		case 7:
			buyItem(player, "Plume de Corbeau", 1)

		case 0:
			fmt.Println("Retour au menu.")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func buyItem(player *Character, itemName string, price int) {
	if player.Gold < price {
		fmt.Println()
		fmt.Println("Vous n'avez pas assez de pièces d'or.")
		return
	}

	player.Gold -= price

	player.Inventory = append(
		player.Inventory,
		Item{Name: itemName},
	)

	fmt.Println()
	fmt.Printf("Vous avez acheté : %s\n", itemName)
	fmt.Printf("Prix : %d pièces d'or\n", price)
	fmt.Printf("Pièces d'or restantes : %d\n", player.Gold)
}
