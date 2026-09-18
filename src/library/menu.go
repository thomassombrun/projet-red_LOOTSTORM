package library

import "fmt"

func MainMenu(player *Character, enemy *Character) {
	for {
		fmt.Println()
		fmt.Println("===== MENU PRINCIPAL =====")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Quitter")

		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			DisplayInfo(player)

		case 2:
			accessInventory(player, enemy)

		case 3:
			fmt.Println("Au revoir !")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}
