package library

import "fmt"

func ForgeronMenu(c *Character) {
	for {
		fmt.Println("\n=== Forgeron ===")
		fmt.Println("1. Chapeau de l'aventurier (5 or)")
		fmt.Println("2. Tunique de l'aventurier (5 or)")
		fmt.Println("3. Bottes de l'aventurier (5 or)")
		fmt.Println("4. Retour au menu principal")
		fmt.Printf("Or actuel: %d\n", c.Gold)
		fmt.Print("Votre choix: ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			FabriquerObjet(c, "Chapeau de l'aventurier")
		case 2:
			FabriquerObjet(c, "Tunique de l'aventurier")
		case 3:
			FabriquerObjet(c, "Bottes de l'aventurier")
		case 4:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func FabriquerObjet(c *Character, name string) {
	materials := map[string]int{}
	switch name {
	case "Chapeau de l'aventurier":
		materials = map[string]int{"Plume de Corbeau": 1, "Cuir de Sanglier": 1}
	case "Tunique de l'aventurier":
		materials = map[string]int{"Fourrure de Loup": 2, "Peau de Troll": 1}
	case "Bottes de l'aventurier":
		materials = map[string]int{"Fourrure de Loup": 1, "Cuir de Sanglier": 1}
	default:
		fmt.Println("Objet inconnu.")
		return
	}

	for material, quantity := range materials {
		found := false
		for i := range c.Inventory {
			if c.Inventory[i].Name == material {
				if c.Inventory[i].Quantity < quantity {
					fmt.Printf("Vous n'avez pas assez de %s pour fabriquer %s.\n", material, name)
					return
				}
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Vous n'avez pas de %s pour fabriquer %s.\n", material, name)
			return
		}
	}

	if IsInventoryFull(c) {
		return
	}

	const cout = 5
	if c.Gold < cout {
		fmt.Println("Pas assez d'or pour fabriquer cet objet.")
		return
	}

	for material, quantity := range materials {
		for i := range c.Inventory {
			if c.Inventory[i].Name == material {
				c.Inventory[i].Quantity -= quantity
				if c.Inventory[i].Quantity <= 0 {
					c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
				}
				break
			}
		}
	}

	c.Gold -= cout
	c.Inventory = append(c.Inventory, Item{Name: name, Quantity: 1})
	fmt.Printf("%s fabriqué(e) et ajouté(e) à l'inventaire !\n", name)
}
