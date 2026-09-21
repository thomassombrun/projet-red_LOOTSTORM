package library

import (
	"fmt"
	"math/rand"
)

func TrainingFight(c *Character) {
	monster := InitGoblin("Gobelin d'entrainement", 40, 5)

	turn := 1

	fmt.Println()
	fmt.Println("===== COMBAT D'ENTRAÎNEMENT =====")
	fmt.Printf("%s affronte %s !\n", c.Name, monster.Name)
	AssassinOpeningAttack(c, &monster)
	WaitForEnter()

	for c.CurrentHP > 0 && monster.CurrentHP > 0 {
		ClearTerminal()

		fmt.Println()
		fmt.Printf("===== TOUR %d =====\n", turn)

		ApplyCharacterEffects(c)
		if c.CurrentHP <= 0 {
			break
		}
		RegenerateMana(c)
		CharacterActions(c, &monster)

		if monster.CurrentHP <= 0 {
			fmt.Printf("%s est vaincu !\n", monster.Name)
			break
		}
		SummonAttack(c, &monster)
		if monster.CurrentHP <= 0 {
			fmt.Printf("%s est vaincu !\n", monster.Name)
			break
		}

		ApplyMonsterEffects(&monster)
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

	fmt.Printf("%s rencontre %s (niveau %d) !\n", c.Name, m.Name, m.Level)

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
	AssassinOpeningAttack(c, m)
	WaitForEnter()

	for c.CurrentHP > 0 && m.CurrentHP > 0 {
		ClearTerminal()

		fmt.Println()
		fmt.Printf("========== TOUR %d ==========\n", turn)

		if playerTurn {

			ApplyCharacterEffects(c)
			if c.CurrentHP <= 0 {
				break
			}
			RegenerateMana(c)
			CharacterActions(c, m)

			if c.CurrentHP <= 0 {
				break
			}

			if m.CurrentHP <= 0 {
				break
			}
			SummonAttack(c, m)
			if m.CurrentHP <= 0 {
				break
			}

			ApplyMonsterEffects(m)
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

			ApplyMonsterEffects(m)
			if m.CurrentHP <= 0 {
				break
			}

			MonsterAttack(m, c, turn)

			if c.CurrentHP <= 0 {
				break
			}

			ApplyCharacterEffects(c)
			if c.CurrentHP <= 0 {
				break
			}
			RegenerateMana(c)

			CharacterActions(c, m)

			if c.CurrentHP <= 0 {
				break
			}

			if m.CurrentHP <= 0 {
				break
			}
			SummonAttack(c, m)
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
			damage := BasicAttackDamage(c)
			skillName := "Coup de poing"
			if IsWeaponEquipped(c) {
				skillName = "Coup d'arme"
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

func CharacterActions(c *Character, m *Monster) {
	CharacterTurn(c, m)
	if c.Class == "Assassin" && c.CurrentHP > 0 && m.CurrentHP > 0 {
		fmt.Println("L'assassin profite de son double tour !")
		WaitForEnter()
		CharacterTurn(c, m)
	}
}

func SummonAttack(c *Character, m *Monster) {
	if c.Summon == nil || c.Summon.CurrentHP <= 0 || m.CurrentHP <= 0 {
		return
	}
	damage := c.Summon.Attack
	m.CurrentHP -= damage
	if m.CurrentHP < 0 {
		m.CurrentHP = 0
	}
	fmt.Printf("%s attaque et inflige %d dégâts à %s.\n", c.Summon.Name, damage, m.Name)
	if c.Summon.CurrentHP <= 0 {
		c.Summon = nil
	}
}

func MonsterAttack(m *Monster, c *Character, turn int) {
	if CheckHolyBarrier(c) {
		return
	}
	if TryDodge(c) {
		fmt.Printf("%s esquive l'attaque de %s !\n", c.Name, m.Name)
		return
	}

	damage := m.Attack

	if turn%3 == 0 {
		damage = m.Attack * 2
		fmt.Printf("%s utilise son attaque renforcée !\n",
			m.Name,
		)
	}
	damage = BlockDamage(c, damage)
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
	if TryCounterAttack(c, m) {
		fmt.Printf("PV de %s : %d / %d\n", m.Name, m.CurrentHP, m.MaxHP)
	}
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
	dropMonsterEquipment(c, m)
	fmt.Printf(
		"Or actuel : %d\n",
		c.Gold,
	)
	fmt.Println("=======================")
}

func dropMonsterEquipment(c *Character, m *Monster) {
	if rand.Intn(100) >= 35 {
		return
	}

	var drops []string
	switch m.Name {
	case "Gobelin", "Boss gobelin":
		drops = []string{"Casque de gobelin", "Lame de gobelin"}
	case "Slime":
		drops = []string{"Carapace de slime", "Bave de slime"}
	default:
		return
	}

	drop := drops[rand.Intn(len(drops))]
	fmt.Printf("%s a laissé tomber : %s !\n", m.Name, drop)
	c.AddOrMergeItem(drop, 1)
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
		manaGain := 10
		if c.Class == "Mage" {
			manaGain = 20
		}
		c.MaxMana += manaGain
		c.Mana += manaGain

		for i := range c.Skill {
			c.Skill[i].Damage += 2
		}

		fmt.Printf("Niveau %d atteint ! PV max +10, mana max +%d, initiative +5, dégâts des sorts +2.\n", c.Level, manaGain)
	}

	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
	fmt.Printf("XP : %d / %d\n", c.CurrentXP, c.MaxXP)
}

func RegenerateMana(c *Character) {
	const regeneration = 5
	if c.Mana >= c.MaxMana {
		return
	}
	c.Mana += regeneration
	if c.Mana > c.MaxMana {
		c.Mana = c.MaxMana
	}
	fmt.Printf("%s récupère de la mana : %d / %d.\n", c.Name, c.Mana, c.MaxMana)
}
