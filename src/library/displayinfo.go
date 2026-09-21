package library

import "fmt"

func (c Character) DisplayInfo() {
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Class: %s\n", c.Class)
	fmt.Printf("Level: %d\n", c.Level)
	fmt.Printf("MaxHP: %d\n", c.MaxHP)
	fmt.Printf("CurrentHP: %d\n", c.CurrentHP)
}

func DisplayInfo(c *Character) {
	ClearTerminal()

	fmt.Println()
	fmt.Println("===== INFORMATIONS DU PERSONNAGE =====")

	fmt.Printf("Nom : %s\n", c.Name)
	fmt.Printf("Classe : %s\n", c.Class)
	fmt.Printf("Niveau : %d\n", c.Level)
	fmt.Printf("PV : %d / %d\n", c.CurrentHP, c.MaxHP)
	fmt.Printf("XP : %d / %d\n", c.CurrentXP, c.MaxXP)
	fmt.Printf("Or : %d\n", c.Gold)
	fmt.Printf("Initiative : %d\n", c.Initiative)
	fmt.Printf("Mana : %d / %d\n", c.Mana, c.MaxMana)
	fmt.Printf("Inventaire : %d / %d emplacements utilisés\n", len(c.Inventory), c.LimitInventory)
	fmt.Printf("Slots libres : %d\n", c.LimitInventory-len(c.Inventory))
	fmt.Printf("Améliorations d'inventaire : %d / 3\n", c.LimitInventoryUpgrade)
	fmt.Printf("Salle la plus avancée : %d\n", c.LastClearedRoom)

	if c.PoisonTurns > 0 {
		fmt.Printf("Poison : %d tours restants\n", c.PoisonTurns)
	}

	DisplayEquipmentAndSkills(c)
}

func DisplayEquipmentAndSkills(c *Character) {
	firstDisplay := true

	for {
		if !firstDisplay {
			ClearTerminal()
		}
		firstDisplay = false

		fmt.Println()
		fmt.Println("===== FICHE PERSONNAGE =====")
		fmt.Println("1. Voir mes sorts")
		fmt.Println("2. Voir mon équipement")
		fmt.Println("0. Retour")
		fmt.Print("Votre choix : ")

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("Choix invalide.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println()
			fmt.Println("===== SORTS =====")
			if len(c.Skill) == 0 {
				fmt.Println("Aucun sort connu.")
			} else {
				for i, s := range c.Skill {
					fmt.Printf("%d. %s (Dégâts : %d, Coût : %d mana)\n", i+1, s.Name, s.Damage, s.ManaCost)
				}
			}
		case 2:
			fmt.Println()
			fmt.Println("===== ÉQUIPEMENT =====")
			fmt.Printf("Casque : %s\n", c.Equip.Helmet)
			fmt.Printf("Plastron : %s\n", c.Equip.Chestplate)
			fmt.Printf("Bottes : %s\n", c.Equip.Boots)
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
