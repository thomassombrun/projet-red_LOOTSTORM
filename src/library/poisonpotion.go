package library

import "fmt"

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
