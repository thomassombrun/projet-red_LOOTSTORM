package library

import (
	"fmt"
	"strings"
)

func MerchantMenu(c *Character) {
	for {
		ClearTerminal()

		fmt.Println()
		fmt.Println("===== MARCHAND =====")
		fmt.Printf("Pièces d'or : %d\n", c.Gold)
		fmt.Println()
		fmt.Println("1. Potion de vie - 3 pièces d'or")
		fmt.Println("2. Potion de poison - 6 pièces d'or")
		fmt.Println("3. Potion de mana - 6 pièces d'or")
		fmt.Println("4. Livre de Sort : Boule de Feu - 25 pièces d'or")
		fmt.Println("5. Livre de Sort : Soin - 25 pièces d'or")
		fmt.Println("6. Livre de Sort : Régénération - 25 pièces d'or")
		fmt.Println("7. Livre de Sort : Poison - 25 pièces d'or")
		fmt.Println("8. Livre de Sort : Brûlure - 25 pièces d'or")
		fmt.Println("9. Livre de Sort : Éclair - 25 pièces d'or")
		fmt.Println("10. Livre de Sort : Barrière Sacrée - 25 pièces d'or")
		fmt.Println("11. Fourrure de Loup - 4 pièces d'or")
		fmt.Println("12. Peau de Troll - 7 pièces d'or")
		fmt.Println("13. Cuir de Sanglier - 3 pièces d'or")
		fmt.Println("14. Plume de Corbeau - 1 pièce d'or")
		fmt.Println("15. Amelioration d'inventaire - 30 pièces d'or")
		fmt.Println("16. Vendre un objet")
		fmt.Println("0. Retour")

		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			BuyItem(c, "Potion de vie", 3)

		case 2:
			BuyItem(c, "Potion de poison", 6)

		case 3:
			BuyItem(c, "Potion de mana", 6)
		case 4:
			BuyItem(c, "Livre de Sort : Boule de Feu", 25)
		case 5:
			BuyItem(c, "Livre de Sort : Soin", 25)
		case 6:
			BuyItem(c, "Livre de Sort : Régénération", 25)
		case 7:
			BuyItem(c, "Livre de Sort : Poison", 25)
		case 8:
			BuyItem(c, "Livre de Sort : Brûlure", 25)
		case 9:
			BuyItem(c, "Livre de Sort : Éclair", 25)
		case 10:
			BuyItem(c, "Livre de Sort : Barrière Sacrée", 25)
		case 11:
			BuyItem(c, "Fourrure de Loup", 4)
		case 12:
			BuyItem(c, "Peau de Troll", 7)
		case 13:
			BuyItem(c, "Cuir de Sanglier", 3)
		case 14:
			BuyItem(c, "Plume de Corbeau", 1)
		case 15:
			BuyItem(c, "Amelioration d'inventaire", 30)
		case 16:
			SellItem(c)

		case 0:
			fmt.Println("Retour au menu.")
			return

		default:
			fmt.Println("Choix invalide.")
		}

		if choice != 0 {
			WaitForEnter()
		}
	}
}

func BuyItem(c *Character, itemName string, price int) bool {
	if itemName == "Amelioration d'inventaire" && c.LimitInventoryUpgrade >= 3 {
		fmt.Println("Vous avez déjà utilisé les 3 améliorations d'inventaire disponibles.")
		return false
	}
	if alreadyKnowsSpell(c, itemName) {
		fmt.Println("Vous connaissez déjà ce sort.")
		return false
	}
	if c.Gold < price {
		fmt.Println()
		fmt.Println("Vous n'avez pas assez de pièces d'or.")
		return false
	}

	if InventoryCount(c)+1 > c.LimitInventory {
		IsInventoryFull(c)
		return false
	}

	c.Gold -= price
	c.AddOrMergeItem(itemName, 1)

	fmt.Println()
	fmt.Printf("Vous avez acheté : %s\n", itemName)
	fmt.Printf("Prix : %d pièces d'or\n", price)
	fmt.Printf("Pièces d'or restantes : %d\n", c.Gold)
	return true
}

// BuyItemSilent applies a purchase without writing to the console.
func BuyItemSilent(c *Character, itemName string, price int) bool {
	if itemName == "Amelioration d'inventaire" && c.LimitInventoryUpgrade >= 3 {
		return false
	}
	if alreadyKnowsSpell(c, itemName) {
		return false
	}
	if c.Gold < price {
		return false
	}
	if InventoryCount(c)+1 > c.LimitInventory {
		return false
	}
	c.Gold -= price
	c.AddOrMergeItem(itemName, 1)
	return true
}

// alreadyKnowsSpell reports whether itemName is a spell book for a skill the character already knows.
func alreadyKnowsSpell(c *Character, itemName string) bool {
	const spellPrefix = "Livre de Sort : "
	if !strings.HasPrefix(itemName, spellPrefix) {
		return false
	}
	spellName := strings.TrimPrefix(itemName, spellPrefix)
	for _, knownSkill := range c.Skill {
		if knownSkill.Name == spellName {
			return true
		}
	}
	return false
}

func SellItem(c *Character) {
	if len(c.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}

	ClearTerminal()
	fmt.Println("===== VENTE =====")
	for i, item := range c.Inventory {
		fmt.Printf("%d. %s x%d\n", i+1, item.Name, item.Quantity)
	}
	fmt.Println("0. Annuler")
	fmt.Print("Choisissez l'objet à vendre : ")

	var choice int
	fmt.Scanln(&choice)
	if choice == 0 {
		return
	}
	if choice < 1 || choice > len(c.Inventory) {
		fmt.Println("Choix invalide.")
		return
	}

	item := c.Inventory[choice-1]
	price := sellPrice(item.Name)
	c.Gold += price
	removeInventoryQuantity(c, choice-1, 1)
	fmt.Printf("Vous avez vendu 1 %s pour %d pièces d'or.\n", item.Name, price)
}

func sellPrice(itemName string) int {
	prices := map[string]int{
		"Potion de vie":                   1,
		"Potion de poison":                3,
		"Potion de mana":                  3,
		"Livre de Sort : Boule de Feu":    12,
		"Livre de Sort : Soin":            12,
		"Livre de Sort : Régénération":    12,
		"Livre de Sort : Poison":          12,
		"Livre de Sort : Brûlure":         12,
		"Livre de Sort : Éclair":          12,
		"Livre de Sort : Barrière Sacrée": 12,
		"Fourrure de Loup":                2,
		"Peau de Troll":                   3,
		"Cuir de Sanglier":                1,
		"Plume de Corbeau":                1,
		"Amelioration d'inventaire":       15,
		"Chapeau de l'aventurier":         8,
		"Tunique de l'aventurier":         12,
		"Bottes de l'aventurier":          10,
		"Casque de gobelin":               8,
		"Carapace de slime":               10,
		"Dague de l'assassin":             8,
		"Arc du chasseur":                 10,
		"Marteau du guerrier":             10,
		"Épée du chevalier":               12,
		"Lame de gobelin":                 8,
		"Bave de slime":                   6,
		"Capuche spectrale":               12,
		"Lame spectrale":                  12,
		"Casque de golem":                 15,
		"Marteau de golem":                15,
		"Peau de troll renforcée":         15,
		"Massue de troll":                 14,
		"Plumes du canard":                10,
		"Bec du canard":                   10,
		"Heaume squelette":                10,
		"Épée squelette":                  10,
		"Fourrure du loup":                8,
		"Crocs du loup":                   8,
		"Chapeau du sorcier":              10,
		"Bâton maudit":                    12,
		"Écailles de dragon":              25,
		"Griffe du dragon":                25,
	}
	if price, found := prices[itemName]; found {
		return price
	}
	return 1
}
