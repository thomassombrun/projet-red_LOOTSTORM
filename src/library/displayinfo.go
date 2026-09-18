package library

import "fmt"

func (c Character) DisplayInfo() {
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Class: %s\n", c.Class)
	fmt.Printf("Level: %d\n", c.Level)
	fmt.Printf("MaxHP: %d\n", c.MaxHP)
	fmt.Printf("CurrentHP: %d\n", c.CurrentHP)
}

func DisplayInfo(player *Character) {
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
