package library

import "fmt"

func MainMenu(c *Character, enemy *Character) {
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
			DisplayInfo(c)

		case 2:
			accessInventory(c, enemy)

		case 3:
			fmt.Println("Au revoir !")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func menu(c *Character) {
	fmt.Println("===== MENU =====")
	fmt.Println("1. Inventaire")
	fmt.Println("2. Marchand")
	fmt.Println("3. Forgeron")
	fmt.Println("4. Quitter")

	var choix int
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		accessInventory(c)

	case 2:
		merchantMenu(c)

	case 3:
		ForgeronMenu(c)

	case 4:
		fmt.Println("Au revoir")
	}
}
