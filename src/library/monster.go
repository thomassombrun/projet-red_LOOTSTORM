package library

import "fmt"

type Monster struct {
	Name        string
	MaxHP       int
	CurrentHP   int
	Attack      int
	PoisonTurns int
}

func InitGobelin(name string, maxHP int, attack int) Monster {
	return Monster{
		Name:        "Gobelin d'entrainement",
		MaxHP:       40,
		CurrentHP:   40,
		Attack:      5,
		PoisonTurns: 0,
	}

}

func goblinPattern(m *Monster, c *Character, turn int) {
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
