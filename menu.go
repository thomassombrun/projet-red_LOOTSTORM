package main

import "fmt"

type Item struct {
	Name string
}

type Character struct {
	Name        string
	Class       string
	Level       int
	MaxHP       int
	CurrentHP   int
	Inventory   []Item
	PoisonTurns int
}

// ====================
// TÂCHE 4 : INVENTAIRE
// ====================

func accessInventory(player *Character, enemy *Character) {
	for {
		fmt.Println()
		fmt.Println("===== INVENTAIRE =====")

		if len(player.Inventory) == 0 {
			fmt.Println("L'inventaire est vide.")
		} else {
			for i, item := range player.Inventory {
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

		if choice < 1 || choice > len(player.Inventory) {
			fmt.Println("Choix invalide.")
			continue
		}

		item := player.Inventory[choice-1]

		switch item.Name {
		case "Potion de vie":
			takePot(player, choice-1)

		case "Potion de poison":
			poisonPot(player, enemy, choice-1)

		default:
			fmt.Println("Cet objet ne peut pas encore être utilisé.")
		}
	}
}

// ====================
// TÂCHE 5 : POTION DE VIE
// ====================

func takePot(player *Character, index int) {
	item := player.Inventory[index]

	if item.Name != "Potion de vie" {
		fmt.Println("Cet objet n'est pas une Potion de vie.")
		return
	}

	// Retire la potion de l'inventaire
	player.Inventory = append(
		player.Inventory[:index],
		player.Inventory[index+1:]...,
	)

	// Rend 50 PV
	player.CurrentHP += 50

	// Empêche de dépasser les PV maximum
	if player.CurrentHP > player.MaxHP {
		player.CurrentHP = player.MaxHP
	}

	fmt.Println()
	fmt.Println("Vous utilisez une Potion de vie.")
	fmt.Printf("PV : %d / %d\n", player.CurrentHP, player.MaxHP)
}

// ====================
// TÂCHE 9 : POTION DE POISON
// ====================

func poisonPot(player *Character, enemy *Character, index int) {
	item := player.Inventory[index]

	if item.Name != "Potion de poison" {
		fmt.Println("Cet objet n'est pas une Potion de poison.")
		return
	}

	// Retire la potion de l'inventaire du joueur
	player.Inventory = append(
		player.Inventory[:index],
		player.Inventory[index+1:]...,
	)

	// Applique le poison à l'ennemi
	enemy.PoisonTurns = 3

	fmt.Println()
	fmt.Printf("%s utilise une Potion de poison sur %s.\n",
		player.Name,
		enemy.Name,
	)

	fmt.Printf("%s est empoisonné pendant 3 tours.\n",
		enemy.Name,
	)
}

// Dégâts du poison.
// Cette fonction doit être appelée une fois par tour de combat.

func poisonEffect(enemy *Character) {
	if enemy.PoisonTurns <= 0 {
		return
	}

	// 10 dégâts de poison
	enemy.CurrentHP -= 10

	// Empêche les PV de passer sous 0
	if enemy.CurrentHP < 0 {
		enemy.CurrentHP = 0
	}

	// Un tour de poison est consommé
	enemy.PoisonTurns--

	fmt.Println()
	fmt.Printf("Le poison inflige 10 dégâts à %s.\n",
		enemy.Name,
	)

	fmt.Printf("PV de %s : %d / %d\n",
		enemy.Name,
		enemy.CurrentHP,
		enemy.MaxHP,
	)

	if enemy.PoisonTurns > 0 {
		fmt.Printf("Il reste %d tours de poison.\n",
			enemy.PoisonTurns,
	)
	} else {
		fmt.Printf("L'effet du poison sur %s est terminé.\n",
			enemy.Name,
		)
	}
}

// ====================
// TÂCHE 3 : INFORMATIONS
// ====================

func displayInfo(player *Character) {
	fmt.Println()
	fmt.Println("===== INFORMATIONS DU PERSONNAGE =====")

	fmt.Printf("Nom : %s\n", player.Name)
	fmt.Printf("Classe : %s\n", player.Class)
	fmt.Printf("Niveau : %d\n", player.Level)
	fmt.Printf("PV : %d / %d\n",
		player.CurrentHP,
		player.MaxHP,
	)

	if player.PoisonTurns > 0 {
		fmt.Printf("Poison : %d tours restants\n",
			player.PoisonTurns,
		)
	}
}

// ====================
// TÂCHE 6 : MENU
// ====================

func mainMenu(player *Character, enemy *Character) {
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
			displayInfo(player)

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

// ====================
// MAIN
// ====================

func main() {

	// Joueur de test
	player := Character{
		Name:        "Samuel",
		Class:       "Elfe",
		Level:       1,
		MaxHP:       100,
		CurrentHP:   100,
		Inventory: []Item{
			{Name: "Potion de vie"},
			{Name: "Potion de poison"},
			{Name: "Potion de vie"},
		},
	}

	// Ennemi de test
	enemy := Character{
		Name:        "Gobelin",
		Class:       "Monstre",
		Level:       1,
		MaxHP:       100,
		CurrentHP:   100,
	}

	mainMenu(&player, &enemy)
}