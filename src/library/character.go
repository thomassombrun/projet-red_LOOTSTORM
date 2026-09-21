package library

import (
	"fmt"
	"math/rand"
)

type Item struct {
	Name     string
	Quantity int
}

type Skill struct {
	Name        string
	Damage      int
	HealAmount  int
	EffectTurns int
	ManaCost    int
}

type Effect struct {
	Name      string
	Value     int
	TurnsLeft int
}

type Character struct {
	Name                  string
	Class                 string
	Level                 int
	MaxHP                 int
	CurrentHP             int
	Inventory             []Item
	Skill                 []Skill
	Effects               []Effect
	HolyBarrier           bool
	PoisonTurns           int
	Gold                  int
	LimitInventory        int
	LimitInventoryUpgrade int
	Equip                 Equipment
	Initiative            int
	CurrentXP             int
	MaxXP                 int
	Mana                  int
	MaxMana               int
	LastClearedRoom       int
}

func InitCharacter(name string, class string, maxHP int) Character {
	initiative := 100
	switch class {
	case "Assassin":
		initiative = 130
	case "Mage":
		initiative = 110
	case "Guerrier":
		initiative = 90
	case "Archer":
		initiative = 120
	case "Chevalier":
		initiative = 80
	}

	return Character{
		Name:                  name,
		Class:                 class,
		Level:                 1,
		MaxHP:                 maxHP,
		CurrentHP:             maxHP / 2,
		Inventory:             []Item{{Name: "Potion de vie", Quantity: 3}},
		Skill:                 []Skill{{Name: "Coup de poing", Damage: 5, ManaCost: 0}},
		Effects:               []Effect{},
		HolyBarrier:           false,
		PoisonTurns:           0,
		Gold:                  100,
		LimitInventory:        10,
		LimitInventoryUpgrade: 0,
		Initiative:            initiative,
		CurrentXP:             0,
		MaxXP:                 100,
		Mana:                  50,
		MaxMana:               50,
		LastClearedRoom:       0,
		Equip: Equipment{
			Helmet:     "Aucun",
			Chestplate: "Aucun",
			Boots:      "Aucun",
			Weapon:     "Aucun",
		},
	}
}

func TryDodge(c *Character) bool {
	chance := 0
	switch c.Class {
	case "Archer":
		chance = 20
	case "Assassin":
		chance = 30
	}
	if chance == 0 || rand.Intn(100) >= chance {
		return false
	}
	fmt.Printf("%s profite de son agilité (%d%% d'esquive).\n", c.Name, chance)
	return true
}

func BlockDamage(c *Character, damage int) int {
	chance := 0
	switch c.Class {
	case "Guerrier":
		chance = 25
	case "Chevalier":
		chance = 35
	}
	if chance == 0 || rand.Intn(100) >= chance {
		return damage
	}

	blockedDamage := damage / 2
	fmt.Printf("%s bloque une partie de l'attaque (%d%% de chance) !\n", c.Name, chance)
	fmt.Printf("Dégâts réduits : %d -> %d.\n", damage, blockedDamage)
	return blockedDamage
}
