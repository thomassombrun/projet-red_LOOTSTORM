package library

import (
	"math/rand"
)

type Item struct {
	Name     string
	Quantity int
	Rarity   Rarity
}

type Skill struct {
	Name        string
	Damage      int
	HealAmount  int
	EffectTurns int
	Initiative  int
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
	Attack                int
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
	Summon                *Monster
	CombatWeaponBonus     int
	CombatSpellBonus      int
	CombatInitiativeBonus int
}

func InitCharacter(name string, class string, maxHP int) Character {
	initiative := 100
	attack := 5
	skills := []Skill{{Name: "Coup de poing", Damage: 5, ManaCost: 0}}
	if class == "Invocateur" {
		skills = append(skills, Skill{Name: "Invocation de soldat", ManaCost: 35})
	}
	switch class {
	case "Guerrier":
		attack = 8
	case "Mage":
		attack = 3
	case "Archer":
		attack = 7
	case "Assassin":
		attack = 6
	case "Chevalier":
		attack = 6
	case "Samourai":
		attack = 8
	case "Clerc":
		attack = 5
	case "Barbare":
		attack = 10
	case "Invocateur":
		attack = 4
	}

	switch class {
	case "Assassin":
		initiative = 125
	case "Mage":
		initiative = 110
	case "Guerrier":
		initiative = 90
	case "Archer":
		initiative = 120
	case "Chevalier":
		initiative = 80
	case "Samourai":
		initiative = 115
	case "Clerc":
		initiative = 95
	case "Barbare":
		initiative = 85
	case "Invocateur":
		initiative = 90
	}

	return Character{
		Name:                  name,
		Class:                 class,
		Level:                 1,
		Attack:                attack,
		MaxHP:                 maxHP,
		CurrentHP:             maxHP / 2,
		Inventory:             []Item{{Name: "Potion de vie", Quantity: 3, Rarity: RarityCommon}},
		Skill:                 skills,
		Effects:               []Effect{},
		HolyBarrier:           false,
		PoisonTurns:           0,
		Gold:                  100,
		LimitInventory:        10,
		LimitInventoryUpgrade: 0,
		Initiative:            initiative,
		CurrentXP:             0,
		MaxXP:                 100,
		Mana:                  baseMana(class),
		MaxMana:               baseMana(class),
		LastClearedRoom:       0,
		Summon:                nil,
		CombatWeaponBonus:     0,
		CombatSpellBonus:      0,
		CombatInitiativeBonus: 0,
		Equip: Equipment{
			Helmet:     "Aucun",
			Chestplate: "Aucun",
			Boots:      "Aucun",
			Weapon:     "Aucun",
		},
	}
}

func baseMana(class string) int {
	switch class {
	case "Mage":
		return 70
	case "Clerc":
		return 60
	case "Invocateur":
		return 80
	default:
		return 50
	}
}

func TryDodge(c *Character) bool {
	chance := 0
	switch c.Class {
	case "Archer":
		chance = 20
	case "Assassin":
		chance = 22
	case "Samourai":
		chance = 15
	}
	if chance == 0 || rand.Intn(100) >= chance {
		return false
	}
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
	return blockedDamage
}

func BasicAttackDamage(c *Character) int {
	damage := 5 + c.Attack + c.CombatWeaponBonus
	if c.Equip.Weapon != "" && c.Equip.Weapon != "Aucun" {
		damage += c.Equip.WeaponDamage
	}
	if c.Class == "Barbare" && c.MaxHP > 0 {
		missingRatio := float64(c.MaxHP-c.CurrentHP) / float64(c.MaxHP)
		damage += int(float64(damage) * missingRatio)
	}
	if c.Class == "Archer" && isBow(c.Equip.Weapon) {
		damage *= 2
	}
	return damage
}

func isBow(itemName string) bool {
	switch normalizeItemName(itemName) {
	case "Arc du chasseur", "Arc long", "Arc composite", "Arc elfique":
		return true
	default:
		return false
	}
}

func IsWeaponEquipped(c *Character) bool {
	return c.Equip.Weapon != "" && c.Equip.Weapon != "Aucun"
}

func StartCombatBonuses(c *Character) {
	c.CombatWeaponBonus = 0
	c.CombatSpellBonus = 0
	c.CombatInitiativeBonus = 0
	if c.Class != "Guerrier" {
		return
	}

	switch rand.Intn(3) {
	case 0:
		c.CombatWeaponBonus = 8
	case 1:
		c.CombatSpellBonus = 8
	case 2:
		c.CombatInitiativeBonus = 20
	}
}

func CombatInitiative(c *Character) int {
	return c.Initiative + c.CombatInitiativeBonus
}

func TryCounterAttack(c *Character, m *Monster) bool {
	if c.Class != "Samourai" || m.CurrentHP <= 0 || rand.Intn(100) >= 25 {
		return false
	}

	damage := BasicAttackDamage(c)
	m.CurrentHP -= damage
	if m.CurrentHP < 0 {
		m.CurrentHP = 0
	}

	return true
}

func AssassinOpeningAttack(c *Character, m *Monster) bool {
	if c.Class != "Assassin" || m.CurrentHP <= 0 {
		return false
	}

	damage := BasicAttackDamage(c) * 3 / 2
	m.CurrentHP -= damage
	if m.CurrentHP < 0 {
		m.CurrentHP = 0
	}

	ApplyAssassinBleed(c, m)
	return true
}

func ApplyAssassinBleed(c *Character, m *Monster) {
	if c.Class != "Assassin" || normalizeItemName(c.Equip.Weapon) != "Dague de l'assassin" {
		return
	}

	damage := 5 + c.Level*2
	m.Effects = append(m.Effects, Effect{Name: "Saignement", Value: damage, TurnsLeft: 3})
}

func SummonStats(c *Character) (int, int) {
	return 25 + c.Level*10, 5 + c.Level*3
}
