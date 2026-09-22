package library

import "fmt"

func ChoisirClass() (string, int) {
	for {
		fmt.Println("Choisissez votre classe :")
		fmt.Println("1. Guerrier (150 PV)")
		fmt.Println("2. Mage (80 PV)")
		fmt.Println("3. Archer (110 PV)")
		fmt.Println("4. Assassin (100 PV)")
		fmt.Println("5. Chevalier (120 PV)")
		fmt.Println("6. Samourai (115 PV)")
		fmt.Println("7. Clerc (105 PV)")
		fmt.Println("8. Barbare (180 PV)")
		fmt.Println("9. Invocateur (90 PV)")
		fmt.Print("Votre choix : ")
		var choice int
		fmt.Scanln(&choice)
		className, maxHP := selectedClass(choice)
		if className == "" {
			fmt.Println("Choix invalide.")
			continue
		}

		fmt.Printf("\n=== Résumé : %s ===\n", className)
		fmt.Printf("PV de base : %d\n", maxHP)
		fmt.Println(classAdvantages(className))
		fmt.Println("1. Valider cette classe")
		fmt.Println("2. Choisir une autre classe")
		fmt.Print("Votre choix : ")
		var confirmation int
		fmt.Scanln(&confirmation)
		if confirmation == 1 {
			return className, maxHP
		}
		if confirmation != 2 {
			fmt.Println("Choix invalide, la classe n'est pas validée.")
		}
	}
}

func selectedClass(choice int) (string, int) {
	classes := map[int]struct {
		name string
		hp   int
	}{
		1: {"Guerrier", 150}, 2: {"Mage", 80}, 3: {"Archer", 110},
		4: {"Assassin", 100}, 5: {"Chevalier", 120}, 6: {"Samourai", 115},
		7: {"Clerc", 105}, 8: {"Barbare", 180}, 9: {"Invocateur", 90},
	}
	classInfo, found := classes[choice]
	if !found {
		return "", 0
	}
	return classInfo.name, classInfo.hp
}

func classAdvantages(className string) string {
	advantages := map[string]string{
		"Guerrier":   "Atout aléatoire au début du combat : dégâts d'arme, dégâts de sort ou initiative.",
		"Mage":       "Les sorts infligent 1,5 fois plus de dégâts et la mana progresse davantage.",
		"Archer":     "Esquive et dégâts doublés avec un arc.",
		"Assassin":   "Frappe d'ouverture x2, double action, esquive et saignement avec une dague.",
		"Chevalier":  "Blocage amélioré des dégâts.",
		"Samourai":   "Esquive et chance de contre-attaque.",
		"Clerc":      "Soins augmentés.",
		"Barbare":    "Dégâts de base et d'arme augmentés selon les PV perdus.",
		"Invocateur": "Invocation d'un soldat qui évolue avec le niveau.",
	}
	return advantages[className]
}

func ClassAdvantages(className string) string {
	return classAdvantages(className)
}

func LireNomAvecEspaces() string {
	var mot1, mot2, mot3 string
	n, _ := fmt.Scanln(&mot1, &mot2, &mot3)

	nom := mot1
	if n >= 2 {
		nom = nom + " " + mot2
	}
	if n >= 3 {
		nom = nom + " " + mot3
	}
	return nom
}

func CharacterCreation() Character {
	var nom string
	valide := false

	for !valide {
		fmt.Print("Choisissez le nom de votre personnage : ")
		nom = LireNomAvecEspaces()

		valide = len(nom) > 0
		for _, r := range nom {
			if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == ' ') {
				valide = false
			}
		}
		if !valide {
			fmt.Println("Erreur : le nom ne doit contenir que des lettres.")
		}
	}

	nomFinal := ""
	debutMot := true
	for _, r := range nom {
		if r == ' ' {
			nomFinal += " "
			debutMot = true
		} else if debutMot {
			if r >= 'a' && r <= 'z' {
				r = r - 'a' + 'A'
			}
			nomFinal += string(r)
			debutMot = false
		} else {
			if r >= 'A' && r <= 'Z' {
				r = r - 'A' + 'a'
			}
			nomFinal += string(r)
		}
	}
	Class, maxHP := ChoisirClass()
	return InitCharacter(nomFinal, Class, maxHP)
}

func Main() {
	perso := CharacterCreation()
	fmt.Println("Personnage créé :", perso.Name)
}
