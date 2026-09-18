package library

import "fmt"

func takePot(player *Character, index int) {
	item := player.Inventory[index]

	if item.Name != "Potion de vie" {
		fmt.Println("Cet objet n'est pas une Potion de vie.")
		return
	}

	player.Inventory = append(
		player.Inventory[:index],
		player.Inventory[index+1:]...,
	)

	player.CurrentHP += 50

	if player.CurrentHP > player.MaxHP {
		player.CurrentHP = player.MaxHP
	}

	fmt.Println()
	fmt.Println("Vous utilisez une Potion de vie.")
	fmt.Printf("PV : %d / %d\n", player.CurrentHP, player.MaxHP)
}
