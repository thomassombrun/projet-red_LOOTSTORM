package library

import "fmt"

type Monster struct {
	Name        string
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
	return initMonster(name, level, 30, 4, 80, 80, 35)
}

func initMonster(name string, level int, baseHP int, baseAttack int, initiative int, baseXP int, baseGold int) Monster {
	if level < 1 {
		level = 1
	}

	return Monster{
		Name:        name,
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
