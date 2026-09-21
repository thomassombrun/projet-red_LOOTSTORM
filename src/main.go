package main

import (
	"fmt"
	"projet/src/library"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("       BIENVENUE DANS LOOTSTORM         ")
	fmt.Println("========================================")

	var player library.Character

	for {
		fmt.Println("\nChoisissez une option :")
		fmt.Println("1. Choisir un héros prédéfini (Himiko, Link, Patrick)")
		fmt.Println("2. Créer votre propre personnage")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scanln(&choice)

		if choice == 1 {
			c1 := library.InitCharacter("Himiko Toga", "Assassin", 100)
			c2 := library.InitCharacter("Link", "Chevalier", 120)
			c3 := library.InitCharacter("Patrick Bouldefeu", "Mage", 80)

			fmt.Println("\nChoisissez votre héros :")
			fmt.Println("1. Himiko Toga (Assassin)")
			fmt.Println("2. Link (Chevalier)")
			fmt.Println("3. Patrick Bouldefeu (Mage)")
			fmt.Print("Votre choix : ")

			var heroChoice int
			fmt.Scanln(&heroChoice)

			switch heroChoice {
			case 1:
				player = c1
			case 2:
				player = c2
			case 3:
				player = c3
			default:
				fmt.Println("Choix invalide.")
				continue
			}
			break

		} else if choice == 2 {

			var name, class string
			fmt.Print("Entrez le nom de votre héros : ")
			fmt.Scanln(&name)
			fmt.Print("Entrez sa classe (Guerrier, Mage, Archer) : ")
			fmt.Scanln(&class)

			player = library.InitCharacter(name, class, 100)
			break
		} else {
			fmt.Println("Choix invalide, réessayez.")
		}
	}

	fmt.Printf("\nC'est parti, %s (%s) entre dans la légende !\n", player.Name, player.Class)

	for {
		fmt.Println()
		fmt.Println("===== MENU PRINCIPAL =====")
		fmt.Println("1. Afficher mes statistiques")
		fmt.Println("2. Combat d'entraînement")
		fmt.Println("3. Ouvrir l'inventaire")
		fmt.Println("4. Quitter le jeu")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scanln(&choice)

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
