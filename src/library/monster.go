package library

import "fmt"

type Monster struct {
	Name        string
	Pattern     string
	Level       int
	MaxHP       int
	CurrentHP   int
	Attack      int
	PoisonTurns int
	Effects     []Effect
	Initiative  int
	XPReward    int
	GoldReward  int
}

func InitGoblin(name string, maxHP int, attack int) Monster {
	return Monster{
		Name:        name,
		Pattern:     "goblin",
		Level:       1,
		MaxHP:       maxHP,
		CurrentHP:   maxHP,
		Attack:      attack,
		PoisonTurns: 0,
		Effects:     []Effect{},
		Initiative:  100,
		XPReward:    100,
		GoldReward:  50,
	}

}

func InitGoblinLevel(name string, level int) Monster {
	return initMonster(name, level, 40, 5, 100, 100, 50)
}

func InitSlimeLevel(name string, level int) Monster {
	monster := initMonster(name, level, 30, 4, 80, 80, 35)
	monster.Pattern = "slime"
	return monster
}

func InitGhostLevel(name string, level int) Monster {
	monster := initMonster(name, level, 35, 7, 125, 120, 60)
	monster.Pattern = "ghost"
	return monster
}

func InitGolem() Monster {
	return Monster{Name: "Golem de Pierre", Pattern: "golem", Level: 1, MaxHP: 100, CurrentHP: 100, Attack: 12, Initiative: 3, XPReward: 250, GoldReward: 80, Effects: []Effect{}}
}

func InitTroll() Monster {
	return Monster{Name: "Troll des Cavernes", Pattern: "troll", Level: 1, MaxHP: 80, CurrentHP: 80, Attack: 10, Initiative: 6, XPReward: 200, GoldReward: 70, Effects: []Effect{}}
}

func InitDuck() Monster {
	return Monster{Name: "Canard Démoniaque", Pattern: "duck", Level: 1, MaxHP: 300, CurrentHP: 300, Attack: 25, Initiative: 130, XPReward: 382, GoldReward: 37, Effects: []Effect{}}
}

func InitSkeleton() Monster {
	return Monster{Name: "Squelette Guerrier", Pattern: "skeleton", Level: 1, MaxHP: 50, CurrentHP: 50, Attack: 8, Initiative: 15, XPReward: 130, GoldReward: 35, Effects: []Effect{}}
}

func InitWolf() Monster {
	return Monster{Name: "Loup Sauvage", Pattern: "wolf", Level: 1, MaxHP: 45, CurrentHP: 45, Attack: 9, Initiative: 18, XPReward: 110, GoldReward: 30, Effects: []Effect{}}
}

func InitWizard() Monster {
	return Monster{Name: "Sorcier Maudit", Pattern: "wizard", Level: 1, MaxHP: 55, CurrentHP: 55, Attack: 11, Initiative: 12, XPReward: 220, GoldReward: 60, Effects: []Effect{}}
}

func InitDragon() Monster {
	return Monster{Name: "Dragon Ancien", Pattern: "dragon", Level: 1, MaxHP: 200, CurrentHP: 200, Attack: 20, Initiative: 8, XPReward: 500, GoldReward: 200, Effects: []Effect{}}
}

func initMonster(name string, level int, baseHP int, baseAttack int, initiative int, baseXP int, baseGold int) Monster {
	if level < 1 {
		level = 1
	}

	return Monster{
		Name:        name,
		Pattern:     "goblin",
		Level:       level,
		MaxHP:       baseHP + (level-1)*10,
		CurrentHP:   baseHP + (level-1)*10,
		Attack:      baseAttack + (level-1)*2,
		PoisonTurns: 0,
		Effects:     []Effect{},
		Initiative:  initiative + (level-1)*2,
		XPReward:    baseXP + (level-1)*20,
		GoldReward:  baseGold + (level-1)*10,
	}
}

func MonsterPatternDamage(m *Monster, turn int) int {
	damage := m.Attack
	switch m.Pattern {
	case "golem":
		if turn%3 == 0 {
			damage = m.Attack * 2
			fmt.Println("Le Golem charge son poing de pierre !")
		}
	case "troll":
		if turn%3 == 0 {
			m.CurrentHP += 10
			if m.CurrentHP > m.MaxHP {
				m.CurrentHP = m.MaxHP
			}
			fmt.Printf("%s récupère 10 PV !\n", m.Name)
		}
	case "duck":
		if turn%2 == 0 {
			damage = m.Attack * 2
			fmt.Printf("%s esquive l'attaque et contre-attaque !\n", m.Name)
		}
	case "skeleton":
		if turn%4 == 0 {
			damage = m.Attack * 3
			fmt.Println("Le Squelette effectue une attaque fracassante !")
		}
	case "wolf":
		if turn%3 == 0 {
			damage = m.Attack * 2
			fmt.Println("Le Loup bondit sur sa proie !")
		}
	case "wizard":
		if turn%3 == 0 {
			damage = m.Attack * 2
			fmt.Println("Le Sorcier Maudit lance une magie noire !")
		}
	case "dragon":
		if turn%3 == 0 {
			damage = m.Attack * 3
			fmt.Println("Le Dragon crache une immense flamme !")
		}
	}
	return damage
}

func GoblinPattern(m *Monster, c *Character, turn int) {
	damage := m.Attack

	if turn%3 == 0 {
		damage = m.Attack * 2
	}

	c.CurrentHP -= damage

	if c.CurrentHP < 0 {
		c.CurrentHP = 0
	}

	fmt.Printf("%s inflige à %s %d de dégâts\n",
		m.Name,
		c.Name,
		damage,
	)

	fmt.Printf("PV de %s : %d / %d\n",
		c.Name,
		c.CurrentHP,
		c.MaxHP,
	)
}
