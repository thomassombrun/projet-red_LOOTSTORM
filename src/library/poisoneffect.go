package library

import "fmt"

func poisonEffect(enemy *Character) {
	if enemy.PoisonTurns <= 0 {
		return
	}

	enemy.CurrentHP -= 10

	if enemy.CurrentHP < 0 {
		enemy.CurrentHP = 0
	}

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
