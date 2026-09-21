package library

import "fmt"

type Experience struct {
	CurrentXP int
	MaxXP     int
	Level     int
}

func gainExperience(player *Experience, amount int) {
	player.CurrentXP += amount

	fmt.Printf("Vous gagnez %d XP.\n", amount)

	for player.CurrentXP >= player.MaxXP {
		player.CurrentXP -= player.MaxXP
		player.Level++
		player.MaxXP += 10

		fmt.Printf("Vous passez niveau %d !\n", player.Level)
	}

	fmt.Printf("XP : %d / %d\n",
		player.CurrentXP,
		player.MaxXP,
	)
}
