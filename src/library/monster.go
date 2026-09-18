package library

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
