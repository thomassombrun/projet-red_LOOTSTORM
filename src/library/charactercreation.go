package library

import "fmt"

func ChoisirClass() (string, int) {
	var choix int
	valide := false
	Class := ""
	maxHP := 0

	for !valide {
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
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			Class = "Guerrier"
			maxHP = 150
			valide = true
		case 2:
			Class = "Mage"
			maxHP = 80
			valide = true
		case 3:
			Class = "Archer"
			maxHP = 110
			valide = true
		case 4:
			Class = "Assassin"
			maxHP = 100
			valide = true
		case 5:
			Class = "Chevalier"
			maxHP = 120
			valide = true
		case 6:
			Class = "Samourai"
			maxHP = 115
			valide = true
		case 7:
			Class = "Clerc"
			maxHP = 105
			valide = true
		case 8:
			Class = "Barbare"
			maxHP = 180
			valide = true
		case 9:
			Class = "Invocateur"
			maxHP = 90
			valide = true
		default:
			fmt.Println("Choix invalide.")
		}
	}

	return Class, maxHP
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
