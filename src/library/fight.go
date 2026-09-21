package library

import "fmt"

func TrainingFight(c *Character) {
	monster := InitGoblin("Gobelin d'entrainement", 40, 5)

	turn := 1

	fmt.Println()
	fmt.Println("===== COMBAT D'ENTRAÎNEMENT =====")
	fmt.Printf("%s affronte %s !\n", c.Name, monster.Name)
	WaitForEnter()

	for c.CurrentHP > 0 && monster.CurrentHP > 0 {
		ClearTerminal()

		fmt.Println()
		fmt.Printf("===== TOUR %d =====\n", turn)

		CharacterTurn(c, &monster)

		if monster.CurrentHP <= 0 {
			fmt.Printf("%s est vaincu !\n", monster.Name)
			break
		}

		GoblinPattern(&monster, c, turn)

		PoisonEffect(&monster)
		if monster.CurrentHP <= 0 {
			fmt.Printf("%s est vaincu !\n", monster.Name)
			break
		}

		if c.IsDead() {
			c.Respawn()
			break
		}

		WaitForEnter()
		turn++
	}
	fmt.Println()
	fmt.Println("===== FIN DU COMBAT =====")
}

func UseSkill(c *Character, m *Monster) {
	if len(c.Skill) == 0 {
		fmt.Println("Vous ne connaissez aucun sort !")
		return
	}

	fmt.Println("\n--- Vos Sorts ---")
	for i, s := range c.Skill {
		fmt.Printf("%d. %s (Dégâts : %d)\n", i+1, s.Name, s.Damage)
	}
	fmt.Println("0. Retour")
	fmt.Print("Choisissez un sort : ")

	var choix int
	fmt.Scanln(&choix)

	if choix == 0 {
		return
	}

	if choix < 1 || choix > len(c.Skill) {
		fmt.Println("Choix invalide.")
		return
	}

	chosenSkill := c.Skill[choix-1]

	m.CurrentHP -= chosenSkill.Damage
	if m.CurrentHP < 0 {
		m.CurrentHP = 0
	}

	fmt.Printf("%s lance %s et inflige %d dégâts à %s !\n",
		c.Name, chosenSkill.Name, chosenSkill.Damage, m.Name)
	fmt.Printf("%s : PV %d / %d\n", m.Name, m.CurrentHP, m.MaxHP)
}

func normalFight(c *Character, m *Monster, previousRoom int) bool {
	turn := 1

	fmt.Println()
	fmt.Println("================================")
	fmt.Println("          COMBAT")
	fmt.Println("================================")

	fmt.Printf("%s rencontre %s !\n", c.Name, m.Name)

	fmt.Printf("%s : %d / %d PV\n",
		c.Name,
		c.CurrentHP,
		c.MaxHP,
	)

	fmt.Printf("%s : %d / %d PV\n",
		m.Name,
		m.CurrentHP,
		m.MaxHP,
	)

	playerTurn := c.Initiative >= m.Initiative

	if playerTurn {
		fmt.Printf("%s commence le combat !\n", c.Name)
	} else {
		fmt.Printf("%s commence le combat !\n", m.Name)
	}
	WaitForEnter()

	for c.CurrentHP > 0 && m.CurrentHP > 0 {
		ClearTerminal()

		fmt.Println()
		fmt.Printf("========== TOUR %d ==========\n", turn)

		if playerTurn {

			CharacterTurn(c, m)

			if c.CurrentHP <= 0 {
				break
			}

			if m.CurrentHP <= 0 {
				break
			}

			MonsterAttack(m, c, turn)

			if c.CurrentHP <= 0 {
				break
			}

			PoisonEffect(m)

			if m.CurrentHP <= 0 {
				break
			}

		} else {

			MonsterAttack(m, c, turn)

			if c.CurrentHP <= 0 {
				break
			}

			CharacterTurn(c, m)

			if c.CurrentHP <= 0 {
				break
			}

			if m.CurrentHP <= 0 {
				break
			}

			PoisonEffect(m)

			if m.CurrentHP <= 0 {
				break
			}
		}

		WaitForEnter()

		playerTurn = !playerTurn

		turn++
	}

	if m.CurrentHP <= 0 {
		fmt.Println()
		fmt.Println("================================")
		fmt.Println("           VICTOIRE")
		fmt.Println("================================")

		fmt.Printf("%s est vaincu !\n", m.Name)

		GiveCombatReward(c, m)

		return true
	}

	if c.CurrentHP <= 0 {
		fmt.Println()
		fmt.Println("================================")
		fmt.Println("           DÉFAITE")
		fmt.Println("================================")

		fmt.Printf("%s a été vaincu.\n", c.Name)

		c.CurrentHP = c.MaxHP / 2

		c.LastClearedRoom = previousRoom

		fmt.Printf("Vous réapparaissez dans la salle %d.\n",
			c.LastClearedRoom,
		)

		fmt.Printf("PV : %d / %d\n",
			c.CurrentHP,
			c.MaxHP,
		)
	}

	return false
}

func CharacterTurn(c *Character, m *Monster) {
	for {
		ClearTerminal()

		fmt.Println()
		fmt.Println("===== TOUR DU JOUEUR =====")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Sorts")
		fmt.Println("3. Inventaire")
		fmt.Println("4. Abandonner")
		fmt.Println("0. Fin du tour")

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("Saisie invalide.")
			continue
		}

		switch choice {
		case 1:
			damage := 5
			skillName := "Coup de poing"

			if len(c.Skill) > 0 {
				damage = c.Skill[0].Damage
				skillName = c.Skill[0].Name
			}

			m.CurrentHP -= damage
			if m.CurrentHP < 0 {
				m.CurrentHP = 0
			}

			fmt.Printf("%s utilise %s et inflige %d dégâts à %s\n",
				c.Name, skillName, damage, m.Name,
			)
			fmt.Printf("PV de %s : %d / %d\n",
				m.Name, m.CurrentHP, m.MaxHP,
			)
			return

		case 2:
			if SkillMenu(c, m) {
				return
			}
			continue

		case 3:
			if AccessCombatInventory(c, m) {
				return
			}
			continue

		case 4:
			fmt.Println("Vous abandonnez le combat...")
			c.CurrentHP = 0
			return

		case 0:
			fmt.Println("Fin du tour du joueur.")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func MonsterAttack(m *Monster, c *Character, turn int) {

	damage := m.Attack

	if turn%3 == 0 {
		damage = m.Attack * 2
		fmt.Printf("%s utilise son attaque renforcée !\n",
			m.Name,
		)
	}
	c.CurrentHP -= damage
	if c.CurrentHP < 0 {
		c.CurrentHP = 0
	}
	fmt.Printf(
		"%s inflige à %s %d de dégâts\n",
		m.Name,
		c.Name,
		damage,
	)
	fmt.Printf(
		"PV de %s : %d / %d\n",
		c.Name,
		c.CurrentHP,
		c.MaxHP,
	)
}
func GiveCombatReward(c *Character, m *Monster) {

	fmt.Println()
	fmt.Println("===== RÉCOMPENSE =====")
	if m.XPReward > 0 {
		fmt.Printf(
			"Vous gagnez %d XP !\n",
			m.XPReward,
		)
		GainExperience(c, m.XPReward)
	}
	if m.GoldReward > 0 {
		c.Gold += m.GoldReward

		fmt.Printf(
			"Vous gagnez %d pièces d'or !\n",
			m.GoldReward,
		)
	}
	fmt.Printf(
		"Or actuel : %d\n",
		c.Gold,
	)
	fmt.Println("=======================")
}

func GainExperience(c *Character, amount int) {
	fmt.Printf("%s gagne %d points d'expérience !\n", c.Name, amount)
	c.CurrentXP += amount

	for c.CurrentXP >= c.MaxXP {
		c.CurrentXP -= c.MaxXP
		c.Level++
		c.MaxXP += 10
		c.MaxHP += 10
		c.CurrentHP += 10
		c.Initiative += 5

		for i := range c.Skill {
			c.Skill[i].Damage += 2
		}

		fmt.Printf("Niveau %d atteint ! PV max +10, initiative +5, dégâts des sorts +2.\n", c.Level)
	}

	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
	fmt.Printf("XP : %d / %d\n", c.CurrentXP, c.MaxXP)
}
