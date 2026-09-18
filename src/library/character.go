package library

type Item struct {
	Name     string
	Quantity int
}

type Skill struct {
	Name   string
	Damage int
}

type Character struct {
	Name        string
	Class       string
	Level       int
	MaxHP       int
	CurrentHP   int
	Inventory   []Item
	Skill       []Skill
	PoisonTurns int
	Gold        int
}

func InitCharacter(name string, class string, maxHP int) Character {
	return Character{
		Name:        name,
		Class:       class,
		Level:       1,
		MaxHP:       maxHP,
		CurrentHP:   maxHP / 2,
		Inventory:   []Item{{Name: "Potion", Quantity: 3}},
		Skill:       []Skill{{Name: "Coup de poing", Damage: 5}},
		PoisonTurns: 0,
		Gold:        100,
	}
}
