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

func SkillMenu(c *Character, m *Monster) bool {
	if len(c.Skill) == 0 {
		fmt.Println("Vous ne connaissez aucun sort !")
		return false
	}

	ClearTerminal()

	fmt.Println("\n--- Vos Sorts ---")
	for i, s := range c.Skill {
		fmt.Printf("%d. %s (Dégâts : %d)\n", i+1, s.Name, s.Damage)
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

	m.CurrentHP -= chosenSkill.Damage
	if m.CurrentHP < 0 {
		m.CurrentHP = 0
	}

	fmt.Printf("%s lance %s et inflige %d dégâts à %s !\n",
		c.Name, chosenSkill.Name, chosenSkill.Damage, m.Name)
	fmt.Printf("%s : PV %d / %d\n", m.Name, m.CurrentHP, m.MaxHP)

	return true
}
