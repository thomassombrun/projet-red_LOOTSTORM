package main

import (
	"fmt"
	"image/color"
	"projet/src/library"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	viewDashboard = iota
	viewStats
	viewInventory
)

func normalizeClass(class string) string {
	switch class {
	case "Guerrier", "guerrier":
		return "Guerrier"
	case "Mage", "mage":
		return "Mage"
	case "Archer", "archer":
		return "Archer"
	case "Assassin", "assassin":
		return "Assassin"
	case "Chevalier", "chevalier":
		return "Chevalier"
	default:
		return class
	}
}

func getHPForClass(class string) int {
	switch class {
	case "Guerrier":
		return 150
	case "Mage":
		return 80
	case "Archer":
		return 110
	case "Assassin":
		return 100
	case "Chevalier":
		return 120
	default:
		return 100
	}
}

func maincli() {
	fmt.Println("========================================")
	fmt.Println("       BIENVENUE DANS LOOTSTORM         ")
	fmt.Println("========================================")

	var player library.Character
	selected := false

	for !selected {
		fmt.Println("\nChoisissez une option :")
		fmt.Println("1. Choisir un héros prédéfini")
		fmt.Println("2. Créer votre propre personnage")
		fmt.Print("Votre choix : ")

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("Saisie invalide.")
			continue
		}

		switch choice {
		case 1:
			for {
				c1 := library.InitCharacter("Himiko Toga", "Assassin", 100)
				c2 := library.InitCharacter("Link", "Chevalier", 120)
				c3 := library.InitCharacter("Patrick Bouldefeu", "Mage", 80)
				c4 := library.InitCharacter("Musashi", "Samourai", 115)
				c5 := library.InitCharacter("Clara", "Clerc", 105)
				c6 := library.InitCharacter("Ragnar", "Barbare", 180)
				c7 := library.InitCharacter("Orion", "Invocateur", 90)

				fmt.Println("\nChoisissez votre héros :")
				fmt.Println("1. Himiko Toga (Assassin)")
				fmt.Println("2. Link (Chevalier)")
				fmt.Println("3. Patrick Bouldefeu (Mage)")
				fmt.Println("4. Musashi (Samourai)")
				fmt.Println("5. Clara (Clerc)")
				fmt.Println("6. Ragnar (Barbare)")
				fmt.Println("7. Orion (Invocateur)")
				fmt.Println("0. Retour")
				fmt.Print("Votre choix : ")

				var heroChoice int
				if _, err := fmt.Scanln(&heroChoice); err != nil {
					fmt.Println("Choix invalide.")
					continue
				}

				if heroChoice >= 1 && heroChoice <= 7 {
					selectedHero := []library.Character{c1, c2, c3, c4, c5, c6, c7}[heroChoice-1]
					fmt.Printf("\n=== Résumé de %s ===\n", selectedHero.Name)
					fmt.Printf("Classe : %s | PV : %d | Attaque : %d | Initiative : %d\n", selectedHero.Class, selectedHero.MaxHP, selectedHero.Attack, selectedHero.Initiative)
					fmt.Println(library.ClassAdvantages(selectedHero.Class))
					fmt.Println("1. Choisir ce personnage")
					fmt.Println("2. Voir un autre personnage")
					fmt.Print("Votre choix : ")
					var heroConfirmation int
					fmt.Scanln(&heroConfirmation)
					if heroConfirmation != 1 {
						continue
					}
				}

				switch heroChoice {
				case 1:
					player = c1
					selected = true
					break
				case 2:
					player = c2
					selected = true
					break
				case 3:
					player = c3
					selected = true
					break
				case 4:
					player = c4
					selected = true
					break
				case 5:
					player = c5
					selected = true
					break
				case 6:
					player = c6
					selected = true
					break
				case 7:
					player = c7
					selected = true
					break
				case 0:
					fmt.Println("Retour au menu principal.")
					selected = false
					break
				default:
					library.ClearTerminal()

					fmt.Println("Choix invalide.")
					continue
				}

				if selected {
					break
				}
				if heroChoice == 0 {
					break
				}
			}
			if selected {
				break
			}

		case 2:
			player = library.CharacterCreation()
			selected = true

		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}

	fmt.Printf("\nC'est parti, %s (%s) entre dans la légende !\n", player.Name, player.Class)

	for {
		library.ClearTerminal()

		fmt.Println()
		fmt.Println("===== MENU PRINCIPAL =====")
		fmt.Println("1. Afficher mes statistiques")
		fmt.Println("2. Aventure")
		fmt.Println("3. Combat d'entraînement")
		fmt.Println("4. Ouvrir l'inventaire")
		fmt.Println("5. Marchand")
		fmt.Println("6. Forgeron")
		fmt.Println("7. Enchanteur")
		fmt.Println("8. Quitter le jeu")
		fmt.Print("Votre choix : ")

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("Saisie invalide, veuillez réessayer.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println()
			library.DisplayInfo(&player)

		case 2:
			library.StartAdventure(&player)

		case 3:
			library.TrainingFight(&player)

		case 4:
			library.AccessInventory(&player, nil)

		case 5:
			library.MerchantMenu(&player)

		case 6:
			library.ForgeronMenu(&player)

		case 7:
			library.EnchanterMenu(&player)

		case 8:
			fmt.Println("Merci d'avoir joué à Lootstorm ! À bientôt.")
			return

		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}

var menuOptions = []string{
	"Afficher mes statistiques",
	"Aventure",
	"Combat d'entraînement",
	"Ouvrir l'inventaire",
	"Marchand",
	"Forgeron",
	"Enchanteur",
	"Quitter le jeu",
}

type Game struct {
	selected    int
	view        int
	message     string
	player      library.Character
	initialized bool
}

func (g *Game) initPlayer() {
	if g.initialized {
		return
	}
	g.player = library.InitCharacter("Himiko Toga", "Assassin", 100)
	g.message = "Bienvenue dans le donjon."
	g.initialized = true
}

func (g *Game) Update() error {
	g.initPlayer()

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.view = viewDashboard
		g.message = "Retour au tableau de bord."
	}
	if g.view != viewDashboard {
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.selected--
		if g.selected < 0 {
			g.selected = len(menuOptions) - 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.selected++
		if g.selected >= len(menuOptions) {
			g.selected = 0
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.handleChoice(g.selected)
	}

	for i := ebiten.Key1; i <= ebiten.Key8; i++ {
		if inpututil.IsKeyJustPressed(i) {
			idx := int(i - ebiten.Key1)
			g.selected = idx
			g.handleChoice(idx)
		}
	}

	return nil
}

func (g *Game) handleChoice(idx int) {
	switch idx {
	case 0:
		g.view = viewStats
		g.message = "Fiche personnage ouverte."

	case 1:
		g.message = "Le mode aventure sera connecté au prochain écran."

	case 2:
		g.message = "Le combat d'entraînement sera connecté au prochain écran."

	case 3:
		g.view = viewInventory
		g.message = "Inventaire ouvert."

	case 4:
		g.message = "Marchand (à brancher sur library.MerchantMenu)"

	case 5:
		g.message = "Forgeron (à brancher sur library.ForgeronMenu)"

	case 6:
		g.message = "Enchanteur (à brancher sur library.EnchanterMenu)"

	case 7:
		g.message = "Fermeture..."

	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.initPlayer()
	screen.Fill(color.RGBA{R: 16, G: 20, B: 28, A: 255})

	if g.view == viewStats {
		g.drawStats(screen)
		return
	}
	if g.view == viewInventory {
		g.drawInventory(screen)
		return
	}

	text := "LOOTSTORM\n\n"
	text += fmt.Sprintf("Héros : %s\nClasse : %s   Niveau : %d\n\n", g.player.Name, g.player.Class, g.player.Level)

	for i, opt := range menuOptions {
		cursor := "  "
		if i == g.selected {
			cursor = "> "
		}
		text += fmt.Sprintf("%s%d. %s\n", cursor, i+1, opt)
	}

	text += "\nFlèches haut/bas + Entrée, ou touches 1-8\nÉchap : retour au tableau de bord\n\n"
	text += g.message

	ebitenutil.DebugPrintAt(screen, text, 42, 40)
}

func (g *Game) drawStats(screen *ebiten.Image) {
	text := "LOOTSTORM / PERSONNAGE\n\n"
	text += fmt.Sprintf("Nom          %s\nClasse       %s\nNiveau       %d\n\n", g.player.Name, g.player.Class, g.player.Level)
	text += fmt.Sprintf("PV           %d / %d\nMana         %d / %d\nAttaque      %d\nInitiative   %d\nXP           %d / %d\nSalle        %d\n\n", g.player.CurrentHP, g.player.MaxHP, g.player.Mana, g.player.MaxMana, g.player.Attack+g.player.Equip.WeaponDamage, g.player.Initiative, g.player.CurrentXP, g.player.MaxXP, g.player.LastClearedRoom)
	text += fmt.Sprintf("Arme         %s (+%d ATQ)\n\n%s\n\nÉchap : retour", g.player.Equip.Weapon, g.player.Equip.WeaponDamage, library.ClassAdvantages(g.player.Class))
	ebitenutil.DebugPrintAt(screen, text, 42, 40)
}

func (g *Game) drawInventory(screen *ebiten.Image) {
	text := "LOOTSTORM / INVENTAIRE\n\n"
	if len(g.player.Inventory) == 0 {
		text += "Inventaire vide.\n"
	} else {
		for index, item := range g.player.Inventory {
			text += fmt.Sprintf("%d. %s x%d\n", index+1, item.Name, item.Quantity)
		}
	}
	text += fmt.Sprintf("\nEmplacements : %d / %d\nOr : %d\n\nÉchap : retour", len(g.player.Inventory), g.player.LimitInventory, g.player.Gold)
	ebitenutil.DebugPrintAt(screen, text, 42, 40)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	game := &Game{selected: 0}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Lootstorm")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
