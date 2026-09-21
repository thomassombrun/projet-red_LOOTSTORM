package library

import "fmt"

func SpellBook(c *Character) {
	learnSkill(c, Skill{Name: "Boule de Feu", Damage: 35, ManaCost: 50})
}

func LearnHeal(c *Character) {
	learnSkill(c, Skill{Name: "Soin", HealAmount: 30, ManaCost: 40})
}

func LearnRegeneration(c *Character) {
	learnSkill(c, Skill{Name: "Régénération", HealAmount: 10, EffectTurns: 3, ManaCost: 30})
}

func LearnPoison(c *Character) {
	learnSkill(c, Skill{Name: "Poison", Damage: 10, EffectTurns: 3, ManaCost: 25})
}

func LearnBurn(c *Character) {
	learnSkill(c, Skill{Name: "Brûlure", Damage: 15, EffectTurns: 2, ManaCost: 30})
}

func LearnHolyBarrier(c *Character) {
	learnSkill(c, Skill{Name: "Barrière Sacrée", EffectTurns: 1, ManaCost: 35})
}

func LearnSummonSoldier(c *Character) {
	learnSkill(c, Skill{Name: "Invocation de soldat", ManaCost: 35})
}

func LearnSpellBook(c *Character, itemName string) {
	switch itemName {
	case "Livre de Sort : Boule de Feu":
		SpellBook(c)
	case "Livre de Sort : Soin":
		LearnHeal(c)
	case "Livre de Sort : Régénération":
		LearnRegeneration(c)
	case "Livre de Sort : Poison":
		LearnPoison(c)
	case "Livre de Sort : Brûlure":
		LearnBurn(c)
	case "Livre de Sort : Barrière Sacrée":
		LearnHolyBarrier(c)
	}
}

func learnSkill(c *Character, skill Skill) {
	for _, knownSkill := range c.Skill {
		if knownSkill.Name == skill.Name {
			fmt.Printf("Vous connaissez déjà le sort %s.\n", skill.Name)
			return
		}
	}

	c.Skill = append(c.Skill, skill)
	fmt.Printf("Vous avez appris le sort : %s !\n", skill.Name)
}

func SkillMenu(c *Character, m *Monster) bool {
	if len(c.Skill) == 0 {
		fmt.Println("Vous ne connaissez aucun sort !")
		return false
	}

	ClearTerminal()

	fmt.Println("\n--- Vos Sorts ---")
	for i, s := range c.Skill {
		fmt.Printf("%d. %s (", i+1, s.Name)
		statPrinted := false
		if s.Damage > 0 {
			fmt.Printf("Dégâts : %d", s.Damage)
			statPrinted = true
		}
		if s.HealAmount > 0 {
			if statPrinted {
				fmt.Print(", ")
			}
			fmt.Printf("Soin : %d", s.HealAmount)
			statPrinted = true
		}
		if statPrinted {
			fmt.Printf(", Mana : %d)\n", s.ManaCost)
		} else {
			fmt.Printf("Mana : %d)\n", s.ManaCost)
		}
	}
	fmt.Println("0. Retour")
	fmt.Print("Choisissez un sort : ")

	var choix int
	fmt.Scanln(&choix)

	if choix == 0 {
		return false
	}

	if choix < 1 || choix > len(c.Skill) {
		fmt.Println("Choix invalide.")
		return false
	}

	chosenSkill := c.Skill[choix-1]
	if c.Mana < chosenSkill.ManaCost {
		fmt.Println("Vous n'avez pas assez de mana !")
		return false
	}
	c.Mana -= chosenSkill.ManaCost

	switch chosenSkill.Name {
	case "Soin":
		healAmount := chosenSkill.HealAmount
		if c.Class == "Clerc" {
			healAmount = healAmount * 3 / 2
		}
		c.CurrentHP += healAmount
		if c.CurrentHP > c.MaxHP {
			c.CurrentHP = c.MaxHP
		}
		fmt.Printf("%s récupère %d PV.\n", c.Name, healAmount)
	case "Régénération":
		c.Effects = append(c.Effects, Effect{Name: "Régénération", Value: chosenSkill.HealAmount, TurnsLeft: chosenSkill.EffectTurns})
		fmt.Printf("%s récupérera %d PV pendant %d tours.\n", c.Name, chosenSkill.HealAmount, chosenSkill.EffectTurns)
	case "Poison", "Brûlure":
		damage := SpellDamage(c, chosenSkill.Damage)
		m.Effects = append(m.Effects, Effect{Name: chosenSkill.Name, Value: damage, TurnsLeft: chosenSkill.EffectTurns})
		fmt.Printf("%s applique %s à %s pendant %d tours.\n", c.Name, chosenSkill.Name, m.Name, chosenSkill.EffectTurns)
	case "Barrière Sacrée":
		c.HolyBarrier = true
		fmt.Println("La prochaine attaque ennemie sera complètement bloquée.")
	case "Invocation de soldat":
		if c.Summon != nil && c.Summon.CurrentHP > 0 {
			fmt.Println("Vous avez déjà un soldat invoqué.")
			c.Mana += chosenSkill.ManaCost
			return false
		}
		c.Summon = &Monster{
			Name: "Soldat invoqué", Level: c.Level, MaxHP: 25 + c.Level*5,
			CurrentHP: 25 + c.Level*5, Attack: 5 + c.Level*2,
			Initiative: c.Initiative, XPReward: 0, GoldReward: 0,
		}
		fmt.Printf("%s invoque un soldat qui combattra à ses côtés.\n", c.Name)
	default:
		damage := SpellDamage(c, chosenSkill.Damage)
		m.CurrentHP -= damage
		if m.CurrentHP < 0 {
			m.CurrentHP = 0
		}
		fmt.Printf("%s lance %s et inflige %d dégâts à %s !\n", c.Name, chosenSkill.Name, damage, m.Name)
	}

	if chosenSkill.Damage > 0 {
		fmt.Printf("%s : PV %d / %d\n", m.Name, m.CurrentHP, m.MaxHP)
	}

	return true
}

func SpellDamage(c *Character, damage int) int {
	if c.Class == "Mage" {
		return damage * 3 / 2
	}
	return damage
}

func ApplyCharacterEffects(c *Character) {
	for i := 0; i < len(c.Effects); {
		effect := &c.Effects[i]
		if effect.Name == "Régénération" {
			c.CurrentHP += effect.Value
			if c.CurrentHP > c.MaxHP {
				c.CurrentHP = c.MaxHP
			}
			fmt.Printf("%s récupère %d PV grâce à Régénération.\n", c.Name, effect.Value)
		}
		effect.TurnsLeft--
		if effect.TurnsLeft <= 0 {
			fmt.Printf("%s n'est plus affecté par %s.\n", c.Name, effect.Name)
			c.Effects = append(c.Effects[:i], c.Effects[i+1:]...)
			continue
		}
		i++
	}
}

func ApplyMonsterEffects(m *Monster) {
	for i := 0; i < len(m.Effects); {
		effect := &m.Effects[i]
		m.CurrentHP -= effect.Value
		if m.CurrentHP < 0 {
			m.CurrentHP = 0
		}
		fmt.Printf("%s subit %d dégâts de %s.\n", m.Name, effect.Value, effect.Name)
		effect.TurnsLeft--
		if effect.TurnsLeft <= 0 {
			fmt.Printf("%s n'est plus affecté par %s.\n", m.Name, effect.Name)
			m.Effects = append(m.Effects[:i], m.Effects[i+1:]...)
			continue
		}
		i++
	}
}

func CheckHolyBarrier(c *Character) bool {
	if !c.HolyBarrier {
		return false
	}

	c.HolyBarrier = false
	fmt.Printf("%s bloque complètement l'attaque ennemie grâce à Barrière Sacrée !\n", c.Name)
	return true
}
