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
	viewStart = iota
	viewHeroSelect
	viewHeroConfirm
	viewNameInput
	viewClassSelect
	viewClassConfirm
	viewDashboard
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
	case "Samourai":
		return 115
	case "Clerc":
		return 105
	case "Barbare":
		return 180
	case "Invocateur":
		return 90
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
	selected      int
	heroSelected  int
	view          int
	message       string
	player        library.Character
	initialized   bool
	startSelected int
	classSelected int
	customName    string
}

func (g *Game) initPlayer() {
	if g.initialized || g.view != viewStart {
		return
	}
	g.view = viewStart
	g.message = "Choisissez votre mode de création."
}

func enterPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyNumpadEnter)
}

func heroChoices() []library.Character {
	return []library.Character{
		library.InitCharacter("Himiko Toga", "Assassin", 100),
		library.InitCharacter("Link", "Chevalier", 120),
		library.InitCharacter("Patrick Bouldefeu", "Mage", 80),
		library.InitCharacter("Musashi", "Samourai", 115),
		library.InitCharacter("Clara", "Clerc", 105),
		library.InitCharacter("Ragnar", "Barbare", 180),
		library.InitCharacter("Orion", "Invocateur", 90),
	}
}

func (g *Game) Update() error {
	g.initPlayer()

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		if g.view == viewStart {
			return ebiten.Termination
		}
		switch g.view {
		case viewHeroSelect, viewNameInput, viewClassSelect:
			g.view = viewStart
		case viewHeroConfirm:
			g.view = viewHeroSelect
		case viewClassConfirm:
			g.view = viewClassSelect
		default:
			g.view = viewDashboard
		}
		g.message = "Retour à l'écran précédent."
	}
	if g.view == viewStart {
		return g.updateStartSelection()
	}
	if g.view == viewHeroSelect {
		return g.updateHeroSelection()
	}
	if g.view == viewHeroConfirm {
		return g.updateHeroConfirmation()
	}
	if g.view == viewNameInput {
		return g.updateNameInput()
	}
	if g.view == viewClassSelect {
		return g.updateClassSelection()
	}
	if g.view == viewClassConfirm {
		return g.updateClassConfirmation()
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

	if enterPressed() {
		if err := g.handleChoice(g.selected); err != nil {
			return err
		}
	}

	for i := ebiten.Key1; i <= ebiten.Key8; i++ {
		if inpututil.IsKeyJustPressed(i) {
			idx := int(i - ebiten.Key1)
			g.selected = idx
			if err := g.handleChoice(idx); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *Game) updateStartSelection() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.startSelected = 1 - g.startSelected
	}
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.startSelected = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.startSelected = 1
	}
	if enterPressed() {
		if g.startSelected == 0 {
			g.view = viewHeroSelect
			g.message = "Choisissez un héros puis consultez sa confirmation."
		} else {
			g.customName = ""
			g.view = viewNameInput
			g.message = "Saisissez le nom de votre personnage."
		}
	}
	return nil
}

func (g *Game) updateHeroSelection() error {
	heroes := heroChoices()
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.heroSelected = (g.heroSelected + len(heroes) - 1) % len(heroes)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.heroSelected = (g.heroSelected + 1) % len(heroes)
	}
	for i := ebiten.Key1; i <= ebiten.Key7; i++ {
		if inpututil.IsKeyJustPressed(i) {
			g.heroSelected = int(i - ebiten.Key1)
		}
	}
	if enterPressed() {
		g.view = viewHeroConfirm
	}
	return nil
}

func (g *Game) updateHeroConfirmation() error {
	if inpututil.IsKeyJustPressed(ebiten.Key1) || enterPressed() {
		g.player = heroChoices()[g.heroSelected]
		g.initialized = true
		g.view = viewDashboard
		g.message = fmt.Sprintf("%s rejoint l'aventure.", g.player.Name)
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.view = viewHeroSelect
	}
	return nil
}

func (g *Game) updateNameInput() error {
	for _, char := range ebiten.AppendInputChars(nil) {
		if len([]rune(g.customName)) < 20 {
			g.customName += string(char)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		runes := []rune(g.customName)
		if len(runes) > 0 {
			g.customName = string(runes[:len(runes)-1])
		}
	}
	if enterPressed() && len([]rune(g.customName)) > 0 {
		g.classSelected = 0
		g.view = viewClassSelect
	}
	return nil
}

func classChoices() []string {
	return []string{"Guerrier", "Mage", "Archer", "Assassin", "Chevalier", "Samourai", "Clerc", "Barbare", "Invocateur"}
}

func (g *Game) updateClassSelection() error {
	classes := classChoices()
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.classSelected = (g.classSelected + len(classes) - 1) % len(classes)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.classSelected = (g.classSelected + 1) % len(classes)
	}
	for i := ebiten.Key1; i <= ebiten.Key9; i++ {
		if inpututil.IsKeyJustPressed(i) {
			g.classSelected = int(i - ebiten.Key1)
		}
	}
	if enterPressed() {
		g.view = viewClassConfirm
	}
	return nil
}

func (g *Game) updateClassConfirmation() error {
	if inpututil.IsKeyJustPressed(ebiten.Key1) || enterPressed() {
		className := classChoices()[g.classSelected]
		g.player = library.InitCharacter(g.customName, className, getHPForClass(className))
		g.initialized = true
		g.view = viewDashboard
		g.message = fmt.Sprintf("%s rejoint l'aventure.", g.player.Name)
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.view = viewClassSelect
	}
	return nil
}

func (g *Game) handleChoice(idx int) error {
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
		return ebiten.Termination
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.initPlayer()
	screen.Fill(color.RGBA{R: 10, G: 14, B: 22, A: 255})
	if g.view == viewStart {
		g.drawStart(screen)
		return
	}
	if g.view == viewHeroSelect {
		g.drawHeroSelection(screen)
		return
	}
	if g.view == viewHeroConfirm {
		g.drawHeroConfirmation(screen)
		return
	}
	if g.view == viewNameInput {
		g.drawNameInput(screen)
		return
	}
	if g.view == viewClassSelect {
		g.drawClassSelection(screen)
		return
	}
	if g.view == viewClassConfirm {
		g.drawClassConfirmation(screen)
		return
	}

	if g.view == viewStats {
		g.drawStats(screen)
		return
	}
	if g.view == viewInventory {
		g.drawInventory(screen)
		return
	}

	g.drawDashboard(screen)
}

var (
	panelColor = color.RGBA{R: 22, G: 29, B: 43, A: 255}
	panelEdge  = color.RGBA{R: 48, G: 61, B: 83, A: 255}
	goldColor  = color.RGBA{R: 238, G: 184, B: 74, A: 255}
	blueColor  = color.RGBA{R: 76, G: 164, B: 235, A: 255}
	redColor   = color.RGBA{R: 226, G: 83, B: 91, A: 255}
)

func drawPanel(screen *ebiten.Image, x, y, width, height int) {
	ebitenutil.DrawRect(screen, float64(x+2), float64(y+2), float64(width), float64(height), color.RGBA{R: 5, G: 8, B: 14, A: 180})
	ebitenutil.DrawRect(screen, float64(x), float64(y), float64(width), float64(height), panelColor)
	ebitenutil.DrawRect(screen, float64(x), float64(y), float64(width), 1, panelEdge)
	ebitenutil.DrawRect(screen, float64(x), float64(y+height-1), float64(width), 1, panelEdge)
}

func drawBar(screen *ebiten.Image, x, y, width, height, value, maximum int, fill color.Color) {
	if maximum < 1 {
		maximum = 1
	}
	ratio := float64(value) / float64(maximum)
	if ratio < 0 {
		ratio = 0
	}
	if ratio > 1 {
		ratio = 1
	}
	ebitenutil.DrawRect(screen, float64(x), float64(y), float64(width), float64(height), color.RGBA{R: 8, G: 12, B: 20, A: 255})
	ebitenutil.DrawRect(screen, float64(x), float64(y), float64(width)*ratio, float64(height), fill)
}

func (g *Game) drawLogo(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, 28, 21, 5, 24, goldColor)
	ebitenutil.DrawRect(screen, 38, 21, 5, 24, goldColor)
	ebitenutil.DrawRect(screen, 48, 21, 5, 24, goldColor)
	ebitenutil.DrawRect(screen, 33, 29, 15, 5, goldColor)
	ebitenutil.DebugPrintAt(screen, "LOOTSTORM", 66, 25)
	ebitenutil.DebugPrintAt(screen, "DUNGEON // ADVENTURE", 66, 40)
}

func (g *Game) drawDashboard(screen *ebiten.Image) {
	g.drawLogo(screen)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s  //  %s", g.player.Name, g.player.Class), 500, 28)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("NIVEAU %d", g.player.Level), 590, 43)

	drawPanel(screen, 28, 75, 300, 142)
	ebitenutil.DebugPrintAt(screen, "PROFIL DU HÉROS", 44, 91)
	ebitenutil.DebugPrintAt(screen, g.player.Name, 44, 111)
	ebitenutil.DebugPrintAt(screen, g.player.Class, 44, 126)
	drawBar(screen, 44, 151, 250, 10, g.player.CurrentHP, g.player.MaxHP, redColor)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("PV  %d / %d", g.player.CurrentHP, g.player.MaxHP), 44, 164)
	drawBar(screen, 44, 181, 250, 10, g.player.Mana, g.player.MaxMana, blueColor)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("MANA  %d / %d", g.player.Mana, g.player.MaxMana), 44, 194)

	drawPanel(screen, 344, 75, 328, 142)
	ebitenutil.DebugPrintAt(screen, "PROGRESSION", 360, 91)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("ÉTAGE %d   //   SALLE %d", dungeonFloorForUI(g.player.LastClearedRoom+1), g.player.LastClearedRoom+1), 360, 112)
	drawBar(screen, 360, 143, 280, 12, g.player.CurrentXP, g.player.MaxXP, goldColor)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("XP  %d / %d", g.player.CurrentXP, g.player.MaxXP), 360, 159)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("ATTAQUE  %d     INITIATIVE  %d", g.player.Attack+g.player.Equip.WeaponDamage, g.player.Initiative), 360, 181)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("OR  %d     INVENTAIRE  %d / %d", g.player.Gold, len(g.player.Inventory), g.player.LimitInventory), 360, 196)

	drawPanel(screen, 28, 238, 300, 180)
	ebitenutil.DebugPrintAt(screen, "NAVIGATION", 44, 254)
	for index, option := range menuOptions {
		cursor := "  "
		if index == g.selected {
			cursor = ">>"
		}
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s %d  %s", cursor, index+1, option), 44, 274+index*17)
	}

	drawPanel(screen, 344, 238, 328, 180)
	ebitenutil.DebugPrintAt(screen, "ÉQUIPEMENT ACTIF", 360, 254)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("ARME       %s", g.player.Equip.Weapon), 360, 278)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("CASQUE     %s", g.player.Equip.Helmet), 360, 299)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("PLASTRON   %s", g.player.Equip.Chestplate), 360, 320)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("BOTTES     %s", g.player.Equip.Boots), 360, 341)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("BONUS ARME  +%d ATQ", g.player.Equip.WeaponDamage), 360, 369)
	ebitenutil.DebugPrintAt(screen, g.message, 360, 395)
	ebitenutil.DebugPrintAt(screen, "FLÈCHES  naviguer     ENTRÉE  sélectionner     ÉCHAP  retour", 28, 432)
}

func dungeonFloorForUI(room int) int {
	if room < 1 {
		return 1
	}
	return (room-1)/10 + 1
}

func (g *Game) drawStart(screen *ebiten.Image) {
	first := "> "
	second := "  "
	if g.startSelected == 1 {
		first, second = "  ", "> "
	}
	text := fmt.Sprintf("LOOTSTORM / NOUVELLE PARTIE\n\n%s1. Choisir un personnage prédéfini\n%s2. Créer mon personnage\n\nFlèches haut/bas, 1-2 puis Entrée\nÉchap : quitter", first, second)
	ebitenutil.DebugPrintAt(screen, text, 42, 40)
}

func (g *Game) drawHeroSelection(screen *ebiten.Image) {
	heroes := heroChoices()
	selectedHero := heroes[g.heroSelected]
	text := "LOOTSTORM / CHOIX DU HÉROS\n\n"
	for index, hero := range heroes {
		cursor := "  "
		if index == g.heroSelected {
			cursor = "> "
		}
		text += fmt.Sprintf("%s%d. %-20s %-12s PV %d\n", cursor, index+1, hero.Name, hero.Class, hero.MaxHP)
	}
	text += fmt.Sprintf("\n=== %s ===\nClasse : %s\nPV : %d\nAttaque : %d\nInitiative : %d\n\n%s\n\nFlèches haut/bas ou 1-7, Entrée : confirmer\nÉchap : quitter", selectedHero.Name, selectedHero.Class, selectedHero.MaxHP, selectedHero.Attack, selectedHero.Initiative, library.ClassAdvantages(selectedHero.Class))
	ebitenutil.DebugPrintAt(screen, text, 42, 40)
}

func (g *Game) drawHeroConfirmation(screen *ebiten.Image) {
	hero := heroChoices()[g.heroSelected]
	text := fmt.Sprintf("LOOTSTORM / CONFIRMATION\n\n%s\nClasse : %s\nPV : %d\nAttaque : %d\nInitiative : %d\n\n%s\n\n1. Valider ce personnage\n2. Choisir un autre personnage\n\nEntrée : valider", hero.Name, hero.Class, hero.MaxHP, hero.Attack, hero.Initiative, library.ClassAdvantages(hero.Class))
	ebitenutil.DebugPrintAt(screen, text, 42, 40)
}

func (g *Game) drawNameInput(screen *ebiten.Image) {
	text := fmt.Sprintf("LOOTSTORM / CRÉATION\n\nChoisissez le nom de votre personnage :\n\n> %s_\n\nEntrée : continuer\nRetour arrière : effacer\nÉchap : retour", g.customName)
	ebitenutil.DebugPrintAt(screen, text, 42, 40)
}

func (g *Game) drawClassSelection(screen *ebiten.Image) {
	classes := classChoices()
	text := fmt.Sprintf("LOOTSTORM / CHOIX DE CLASSE\n\nNom : %s\n\n", g.customName)
	for index, className := range classes {
		cursor := "  "
		if index == g.classSelected {
			cursor = "> "
		}
		text += fmt.Sprintf("%s%d. %-12s PV %d\n", cursor, index+1, className, getHPForClass(className))
	}
	text += fmt.Sprintf("\nAvantage : %s\n\nFlèches haut/bas ou 1-9, Entrée : continuer", library.ClassAdvantages(classes[g.classSelected]))
	ebitenutil.DebugPrintAt(screen, text, 42, 40)
}

func (g *Game) drawClassConfirmation(screen *ebiten.Image) {
	className := classChoices()[g.classSelected]
	text := fmt.Sprintf("LOOTSTORM / CONFIRMATION DE CLASSE\n\nNom : %s\nClasse : %s\nPV de base : %d\n\n%s\n\n1. Valider cette classe\n2. Choisir une autre classe\n\nEntrée : valider", g.customName, className, getHPForClass(className), library.ClassAdvantages(className))
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
	return 700, 450
}

func main() {
	game := &Game{selected: 0}
	ebiten.SetWindowSize(1400, 900)
	ebiten.SetWindowTitle("Lootstorm")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
