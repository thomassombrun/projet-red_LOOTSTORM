func characterTurn(c *Character, monster *Character) {
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
			monster.CurrentHP -= damage
			if monster.CurrentHP < 0 {
				monster.CurrentHP = 0
			}
			fmt.Printf(
				"%s inflige %d dégâts à %s\n",
				c.Name,
				damage,
				monster.Name,
			)
			fmt.Printf(
				"%s : PV %d / %d\n",
				monster.Name,
				monster.CurrentHP,
				monster.MaxHP,
			)
			return
		case 2:
			accessInventory(c, monster)
		}
	}
}

func TrainingFight(c *Character) {
	monster := InitGoblin()

	turn := 1

	fmt.Println()
	fmt.Println("===== COMBAT D'ENTRAÎNEMENT =====")
	fmt.Printf("%s affronte %s !\n", c.Name, monster.Name)

	for c.CurrentHP > 0 && monster.CurrentHP > 0 {

		fmt.Println()
		fmt.Printf("===== TOUR %d =====\n", turn)

		characterTurn(c, &monster)

		if monster.CurrentHP <= 0 {
			fmt.Printf("%s est vaincu !\n", goblin.Name)
			break
		}

		goblinPattern(&monster, c, turn)

		if c.CurrentHP <= 0 {
			fmt.Printf("%s est vaincu !\n", c.Name)
			break
		}
		turn++
	}
	fmt.Println()
	fmt.Println("===== FIN DU COMBAT =====")
}