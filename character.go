package main

type Item struct {
	Name     string
	Quantity int
}

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHP     int
	CurrentHP int
	Inventory []Item
}

func initCharacter(name string, class string, maxHP int) Character {
	return Character{
		Name:      name,
		Class:     class,
		Level:     1,
		MaxHP:     maxHP,
		CurrentHP: maxHP / 2,
		Inventory: []Item{{Name: "Potion", Quantity: 3}},
	}
}
