package library

import "fmt"

func PoisonEffect(enemy *Monster) {
	if enemy == nil || enemy.PoisonTurns <= 0 {
		return
	}

	enemy.CurrentHP -= 10
	if enemy.CurrentHP <= 0 {
		enemy.CurrentHP = 0
	}
	enemy.PoisonTurns--

	fmt.Println()
	fmt.Printf("%s subit 10 dégâts de poison (%d PV restants).\n", enemy.Name, enemy.CurrentHP)

	if enemy.PoisonTurns > 0 {
		fmt.Printf("Il reste %d tours de poison.\n", enemy.PoisonTurns)
	} else {
		fmt.Printf("L'effet du poison sur %s est terminé.\n", enemy.Name)
	}
}

func ApplyPoisonSilent(enemy *Monster) int {
	if enemy == nil || enemy.PoisonTurns <= 0 {
		return 0
	}
	enemy.CurrentHP -= 10
	if enemy.CurrentHP < 0 {
		enemy.CurrentHP = 0
	}
	enemy.PoisonTurns--
	return 10
}
