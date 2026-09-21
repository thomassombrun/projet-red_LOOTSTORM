package main

import (
	"fmt"
	"projet/src/library"
)

func normalizeClass(class string) string {
	switch class {
	case "Guerrier", "guerrier":
		return "Guerrier"
	case "Mage", "mage":
		return "Mage"
	case "Archer", "archer":
		return "Archer"
	case "Assassin", "assassin":
		return "Assassin"
	case "Chevalier", "chevalier":
		return "Chevalier"
	default:
		return class
	}
}

func getHPForClass(class string) int {
	switch class {
	case "Guerrier":
		return 150
	case "Mage":
		return 80
	case "Archer":
		return 110
	case "Assassin":
		return 100
	case "Chevalier":
		return 120
	default:
		return 100
	}
}

func main() {
	fmt.Println("========================================")
	fmt.Println("       BIENVENUE DANS LOOTSTORM         ")
	fmt.Println("========================================")

	var player library.Character
	selected := false

	for !selected {
		fmt.Println("\nChoisissez une option :")
		fmt.Println("1. Choisir un héros prédéfini (Himiko, Link, Patrick)")
		fmt.Println("2. Créer votre propre personnage")
		fmt.Print("Votre choix : ")

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("Saisie invalide.")
			continue
		}

		switch choice {
		case 1:
			c1 := library.InitCharacter("Himiko Toga", "Assassin", 100)
			c2 := library.InitCharacter("Link", "Chevalier", 120)
			c3 := library.InitCharacter("Patrick Bouldefeu", "Mage", 80)

			fmt.Println("\nChoisissez votre héros :")
			fmt.Println("1. Himiko Toga (Assassin)")
			fmt.Println("2. Link (Chevalier)")
			fmt.Println("3. Patrick Bouldefeu (Mage)")
			fmt.Print("Votre choix : ")

			var heroChoice int
			if _, err := fmt.Scanln(&heroChoice); err != nil {
				fmt.Println("Choix invalide.")
				continue
			}

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

			selected = true

		case 2:
			var name, class string

			fmt.Print("Entrez le nom de votre héros : ")
			if _, err := fmt.Scanln(&name); err != nil {
				fmt.Println("Nom invalide.")
				continue
			}

			fmt.Print("Entrez sa classe (Guerrier, Mage, Archer, Assassin, Chevalier) : ")
			if _, err := fmt.Scanln(&class); err != nil {
				fmt.Println("Classe invalide.")
				continue
			}

			class = normalizeClass(class)
			player = library.InitCharacter(name, class, getHPForClass(class))
			selected = true

		default:
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
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("Saisie invalide, veuillez réessayer.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println()
			library.DisplayInfo(&player)

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
