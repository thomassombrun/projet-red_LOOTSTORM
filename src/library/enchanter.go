package library

import "fmt"

func EnchanterMenu(c *Character) {
	for {
		ClearTerminal()
		fmt.Println("===== ENCHANTEUR =====")
		fmt.Printf("Or : %d\n", c.Gold)

		if len(c.Skill) == 0 {
			fmt.Println("Vous ne connaissez aucun sort.")
			WaitForEnter()
			return
		}

		for i, skill := range c.Skill {
			if skill.Damage <= 0 {
				continue
			}
			cost := enchantmentCost(skill)
			fmt.Printf("%d. %s (+5 dégâts, %d or)\n", i+1, skill.Name, cost)
		}
		fmt.Println("0. Retour")
		fmt.Print("Choisissez un sort à enchanter : ")

		var choice int
		fmt.Scanln(&choice)
		if choice == 0 {
			return
		}
		if choice < 1 || choice > len(c.Skill) {
			fmt.Println("Choix invalide.")
			WaitForEnter()
			continue
		}

		enchantSkill(c, choice-1)
		WaitForEnter()
	}
}

func EnchantSkill(c *Character, index int) bool {
	if index < 0 || index >= len(c.Skill) || c.Skill[index].Damage <= 0 {
		return false
	}
	cost := enchantmentCost(c.Skill[index])
	if c.Gold < cost {
		return false
	}
	c.Gold -= cost
	c.Skill[index].Damage += 5
	return true
}

func enchantmentCost(skill Skill) int {
	return 20 + skill.Damage*2
}

func enchantSkill(c *Character, index int) {
	skill := &c.Skill[index]
	if skill.Damage <= 0 {
		fmt.Println("Ce sort n'inflige pas de dégâts et ne peut pas être enchanté ici.")
		return
	}
	cost := enchantmentCost(*skill)
	if c.Gold < cost {
		fmt.Println("Vous n'avez pas assez d'or.")
		return
	}
	c.Gold -= cost
	skill.Damage += 5
	fmt.Printf("%s a été enchanté : dégâts des sorts +5.\n", skill.Name)
	fmt.Printf("Or restant : %d\n", c.Gold)
}

func EnchantSkillSilent(c *Character, index int) bool {
	if index < 0 || index >= len(c.Skill) || c.Skill[index].Damage <= 0 {
		return false
	}
	cost := enchantmentCost(c.Skill[index])
	if c.Gold < cost {
		return false
	}
	c.Gold -= cost
	c.Skill[index].Damage += 5
	return true
}

func EnchantmentCost(skill Skill) int {
	return enchantmentCost(skill)
}
