package library

import "fmt"

func CharacterTurn(c *Character, m *Monster) {
	for {
		fmt.Println()
		fmt.Println("===== TOUR DU JOUEUR =====")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Inventaire")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			damage := 5
			m.CurrentHP -= damage
			if m.CurrentHP < 0 {
				m.CurrentHP = 0
			}
			fmt.Printf(
				"%s inflige %d dégâts à %s\n",
				c.Name,
				damage,
				m.Name,
			)
			fmt.Printf(
				"%s : PV %d / %d\n",
				m.Name,
				m.CurrentHP,
				m.MaxHP,
			)
			return
		case 2:
			AccessInventory(c, m)
		}
	}
}

func TrainingFight(c *Character) {
	monster := InitGoblin("Gobelin d'entrainement", 40, 5)

	turn := 1

	fmt.Println()
	fmt.Println("===== COMBAT D'ENTRAÎNEMENT =====")
	fmt.Printf("%s affronte %s !\n", c.Name, monster.Name)

	for c.CurrentHP > 0 && monster.CurrentHP > 0 {

		fmt.Println()
		fmt.Printf("===== TOUR %d =====\n", turn)

		CharacterTurn(c, &monster)

		if monster.CurrentHP <= 0 {
			fmt.Printf("%s est vaincu !\n", monster.Name)
			break
		}

		GoblinPattern(&monster, c, turn)

		if c.CurrentHP <= 0 {
			fmt.Printf("%s est vaincu !\n", c.Name)
			break
		}
		turn++
	}
	fmt.Println()
	fmt.Println("===== FIN DU COMBAT =====")
}
