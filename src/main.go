package main

import (
	"fmt"
	"projet/src/library"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("       BIENVENUE DANS LOOTSTORM         ")
	fmt.Println("========================================")

	c1 := library.InitCharacter("Himiko Toga", "Assassin", 100)
	c2 := library.InitCharacter("Link", "Chevalier", 120)
	c3 := library.InitCharacter("Patrick Bouldefeu", "Mage", 80)

	var player library.Character
	for {
		fmt.Println("\nChoisissez votre héros :")
		fmt.Println("1. Himiko Toga (Assassin)")
		fmt.Println("2. Link (Chevalier)")
		fmt.Println("3. Patrick Bouldefeu (Mage)")
		fmt.Print("Votre choix : ")

		var heroChoice int
		fmt.Scan(&heroChoice)

		switch heroChoice {
		case 1:
			player = c1
		case 2:
			player = c2
		case 3:
			player = c3
		default:
			fmt.Println("Choix invalide, réessayez.")
			continue
		}
		break

		fmt.Printf("\nVous avez choisi %s (%s) !\n", player.Name, player.Class)

		for {
			fmt.Println()
			fmt.Println("===== MENU PRINCIPAL =====")
			fmt.Println("1. Afficher mes statistiques")
			fmt.Println("2. Combat d'entraînement")
			fmt.Println("3. Ouvrir l'inventaire")
			fmt.Println("4. Quitter le jeu")
			fmt.Print("Votre choix : ")

			var choice int
			fmt.Scan(&choice)

			switch choice {
			case 1:
				fmt.Println()
				player.DisplayInfo()

			case 2:
				library.TrainingFight(&player)

			case 3:
				library.AccessInventory(&player, nil)

			case 4:
				fmt.Println("Merci d'avoir joué à Lootstorm ! À bientôt.")
				return

			default:
				fmt.Println("Choix invalide, veuillez réessayer.")
			}
		}
	}
}
