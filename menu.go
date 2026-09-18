package main

import "fmt"

type Item struct {
	Name string
}

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHP     int
	CurrentHP int
	Inventory []Item
}

// accessInventory affiche tous les objets présents dans l'inventaire.
func accessInventory(character *Character) {
	for {
		fmt.Println()
		fmt.Println("===== INVENTAIRE =====")

		if len(character.Inventory) == 0 {
			fmt.Println("L'inventaire est vide.")
		} else {
			for i, item := range character.Inventory {
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

		if choice < 1 || choice > len(character.Inventory) {
			fmt.Println("Choix invalide.")
			continue
		}

		item := character.Inventory[choice-1]

		if item.Name == "Potion de vie" {
			takePot(character, choice-1)
		} else {
			fmt.Println("Cet objet ne peut pas encore être utilisé.")
		}
	}
}

// takePot utilise la potion sélectionnée.
func takePot(character *Character, index int) {
	item := character.Inventory[index]

	if item.Name != "Potion de vie" {
		fmt.Println("Cet objet n'est pas une Potion de vie.")
		return
	}

	// Suppression de la potion de l'inventaire.
	character.Inventory = append(
		character.Inventory[:index],
		character.Inventory[index+1:]...,
	)

	// Récupération de 50 PV.
	character.CurrentHP += 50

	// Les PV ne peuvent pas dépasser les PV maximum.
	if character.CurrentHP > character.MaxHP {
		character.CurrentHP = character.MaxHP
	}

	fmt.Println()
	fmt.Println("Vous utilisez une Potion de vie.")
	fmt.Printf("PV : %d / %d\n", character.CurrentHP, character.MaxHP)
}

// mainMenu affiche le menu principal.
func mainMenu(character *Character) {
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
			displayInfo(character)

		case 2:
			accessInventory(character)

		case 3:
			fmt.Println("Au revoir !")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// Exemple temporaire de displayInfo.
// Tes coéquipiers pourront remplacer cette fonction par leur version
// de la tâche 3.
func displayInfo(character *Character) {
	fmt.Println()
	fmt.Println("===== INFORMATIONS =====")
	fmt.Printf("Nom : %s\n", character.Name)
	fmt.Printf("Classe : %s\n", character.Class)
	fmt.Printf("Niveau : %d\n", character.Level)
	fmt.Printf("PV : %d / %d\n", character.CurrentHP, character.MaxHP)
}

func main() {
	character := Character{
		Name:      "Samuel",
		Class:     "Elfe",
		Level:     1,
		MaxHP:     100,
		CurrentHP: 40,
		Inventory: []Item{
			{Name: "Potion de vie"},
			{Name: "Potion de vie"},
			{Name: "Potion de vie"},
		},
	}

	mainMenu(&character)
}