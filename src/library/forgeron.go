package library

import "fmt"

func forgeronMenu(c *Character) {
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
	const cout = 5
	if c.Gold < cout {
		fmt.Println("Pas assez d'or pour fabriquer cet objet.")
		return
	}
	c.Gold -= cout
	c.Inventory = append(c.Inventory, Item{Name: name, Quantity: 1})
	fmt.Printf("%s fabriqué(e) et ajouté(e) à l'inventaire !\n", name)
}
