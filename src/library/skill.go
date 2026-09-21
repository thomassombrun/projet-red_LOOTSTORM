package library

import "fmt"

func SpellBook(c *Character) {
	for _, skill := range c.Skill {
		if skill.Name == "Boule de Feu" {
			fmt.Println("Vous connaissez déjà ce sort.")
			return
		}
	}

	c.Skill = append(c.Skill, Skill{Name: "Boule de Feu", Damage: 15, ManaCost: 10})
	fmt.Println("Vous avez appris le sort : Boule de Feu !")
}
