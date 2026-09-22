package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"projet/src/library"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
	viewMerchant
	viewForge
	viewEnchanter
	viewCombat
	viewCombatSkills
	viewCombatInventory
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
				c3 := library.InitCharacter("Frieren", "Mage", 80)
				c4 := library.InitCharacter("Musashi", "Samourai", 115)
				c5 := library.InitCharacter("Elizabeth", "Clerc", 105)
				c6 := library.InitCharacter("Guts", "Barbare", 180)
				c7 := library.InitCharacter("Sung Jin Woo", "Invocateur", 90)

				fmt.Println("\nChoisissez votre héros :")
				fmt.Println("1. Himiko Toga (Assassin)")
				fmt.Println("2. Link (Chevalier)")
				fmt.Println("3. Frieren (Mage)")
				fmt.Println("4. Musashi (Samourai)")
				fmt.Println("5. Elizabeth (Clerc)")
				fmt.Println("6. Guts (Barbare)")
				fmt.Println("7. Sung Jin Woo (Invocateur)")
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
	selected                int
	heroSelected            int
	view                    int
	message                 string
	player                  library.Character
	initialized             bool
	startSelected           int
	classSelected           int
	customName              string
	enemy                   *library.Monster
	combatTurn              int
	combatMessage           string
	combatAdvancesRoom      bool
	combatStarted           bool
	merchantSelected        int
	skillSelected           int
	combatInventorySelected int
	combatResult            bool
	forgeSelected           int
	enchanterSelected       int
	inventorySelected       int
	combatAction            int
	confirmSelected         int
	frameCount              int
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
		library.InitCharacter("Frieren", "Mage", 80),
		library.InitCharacter("Musashi", "Samourai", 115),
		library.InitCharacter("Elizabeth", "Clerc", 105),
		library.InitCharacter("Guts", "Barbare", 180),
		library.InitCharacter("Sung Jin Woo", "Invocateur", 90),
	}
}

func (g *Game) Update() error {
	g.frameCount++
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
		case viewCombat, viewMerchant, viewForge, viewEnchanter:
			g.view = viewDashboard
		case viewCombatSkills:
			g.view = viewCombat
		case viewCombatInventory:
			g.view = viewCombat
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
	if g.view == viewCombat {
		return g.updateCombat()
	}
	if g.view == viewCombatSkills {
		return g.updateCombatSkills()
	}
	if g.view == viewCombatInventory {
		return g.updateCombatInventory()
	}
	if g.view == viewInventory {
		return g.updateInventory()
	}
	if g.view == viewMerchant {
		return g.updateMerchant()
	}
	if g.view == viewForge {
		return g.updateForge()
	}
	if g.view == viewEnchanter {
		return g.updateEnchanter()
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
		g.confirmSelected = 1 - g.confirmSelected
	}
	if enterPressed() {
		if g.confirmSelected == 0 {
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
	if enterPressed() {
		g.confirmSelected = 0
		g.view = viewHeroConfirm
	}
	return nil
}

func (g *Game) updateHeroConfirmation() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.confirmSelected = 1 - g.confirmSelected
	}
	if enterPressed() && g.confirmSelected == 0 {
		g.player = heroChoices()[g.heroSelected]
		g.initialized = true
		g.view = viewDashboard
		g.message = fmt.Sprintf("%s rejoint l'aventure.", g.player.Name)
	}
	if enterPressed() && g.confirmSelected == 1 {
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
	if enterPressed() {
		g.confirmSelected = 0
		g.view = viewClassConfirm
	}
	return nil
}

func (g *Game) updateClassConfirmation() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.confirmSelected = 1 - g.confirmSelected
	}
	if enterPressed() && g.confirmSelected == 0 {
		className := classChoices()[g.classSelected]
		g.player = library.InitCharacter(g.customName, className, getHPForClass(className))
		g.initialized = true
		g.view = viewDashboard
		g.message = fmt.Sprintf("%s rejoint l'aventure.", g.player.Name)
	}
	if enterPressed() && g.confirmSelected == 1 {
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
		g.startAdventureEncounter()

	case 2:
		g.enemy = makeMonster(library.InitGoblin("Gobelin d'entraînement", 40, 5))
		g.combatTurn = 1
		g.combatAdvancesRoom = false
		g.combatStarted = false
		g.combatResult = false
		g.combatAction = 0
		g.combatMessage = "Le combat d'entraînement commence."
		g.view = viewCombat

	case 3:
		g.view = viewInventory
		g.message = "Inventaire ouvert."

	case 4:
		g.view = viewMerchant
		g.message = ""

	case 5:
		g.view = viewForge
		g.message = "Le forgeron prépare son établi."

	case 6:
		g.view = viewEnchanter
		g.message = "L'enchanteur examine vos sorts."

	case 7:
		g.message = "Fermeture..."
		return ebiten.Termination
	}
	return nil
}

func makeMonster(monster library.Monster) *library.Monster {
	return &monster
}

func (g *Game) startAdventureEncounter() {
	room := g.player.LastClearedRoom + 1
	if room%10 == 0 {
		boss := library.InitGoblinLevel("Boss de l'étage", room+2)
		boss.MaxHP += 60 + dungeonFloorForUI(room)*25
		boss.CurrentHP = boss.MaxHP
		boss.Attack += 8 + dungeonFloorForUI(room)*3
		boss.XPReward = g.player.MaxXP - g.player.CurrentXP + 50
		g.enemy = makeMonster(boss)
	} else {
		switch rand.Intn(7) {
		case 0:
			g.enemy = makeMonster(library.InitGoblinLevel("Gobelin", room))
		case 1:
			g.enemy = makeMonster(library.InitSlimeLevel("Slime", room))
		case 2:
			g.enemy = makeMonster(library.InitGhostLevel("Fantôme", room))
		case 3:
			g.enemy = makeMonster(library.InitGolem())
		case 4:
			g.enemy = makeMonster(library.InitTroll())
		case 5:
			if room >= 3 {
				g.enemy = makeMonster(library.InitDuck())
			} else {
				g.enemy = makeMonster(library.InitWizard())
			}
		default:
			if room >= 5 {
				g.enemy = makeMonster(library.InitDragon())
			} else {
				g.enemy = makeMonster(library.InitWizard())
			}
		}
	}
	g.combatTurn = 1
	g.combatAdvancesRoom = true
	g.combatStarted = false
	g.combatResult = false
	g.combatAction = 0
	g.combatMessage = fmt.Sprintf("%s apparaît dans la salle %d.", g.enemy.Name, room)
	g.view = viewCombat
}

func (g *Game) updateCombat() error {
	if g.enemy == nil {
		if g.combatResult && (enterPressed() || inpututil.IsKeyJustPressed(ebiten.Key1)) {
			g.view = viewDashboard
			g.combatResult = false
		}
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.combatAction = (g.combatAction + 3) % 4
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.combatAction = (g.combatAction + 1) % 4
	}
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.combatAction = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.combatAction = 1
	}
	if inpututil.IsKeyJustPressed(ebiten.Key3) {
		g.combatAction = 2
	}
	if inpututil.IsKeyJustPressed(ebiten.Key4) {
		g.combatAction = 3
	}
	if !g.combatStarted {
		library.StartCombatBonuses(&g.player)
		library.AssassinOpeningAttack(&g.player, g.enemy)
		g.combatStarted = true
		if g.enemy.CurrentHP <= 0 {
			drops := []string{}
			if g.combatAdvancesRoom {
				g.player.LastClearedRoom++
			}
			if g.combatAdvancesRoom {
				library.GainExperience(&g.player, g.enemy.XPReward)
				g.player.Gold += g.enemy.GoldReward
				drops = library.DropMonsterEquipment(&g.player, g.enemy)
			}
			g.message = "Victoire grâce à votre frappe d'ouverture."
			if len(drops) > 0 {
				g.message += fmt.Sprintf(" Loot : %v.", drops)
			}
			g.enemy = nil
			g.combatResult = true
			return nil
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) || (enterPressed() && g.combatAction == 1) {
		g.skillSelected = 0
		g.view = viewCombatSkills
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.Key3) || (enterPressed() && g.combatAction == 2) {
		g.combatInventorySelected = 0
		g.view = viewCombatInventory
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.Key1) || (enterPressed() && g.combatAction == 0) {
		library.ApplyCharacterEffects(&g.player)
		damage := library.BasicAttackDamage(&g.player)
		g.enemy.CurrentHP -= damage
		if g.enemy.CurrentHP < 0 {
			g.enemy.CurrentHP = 0
		}
		if g.player.Class == "Assassin" && g.enemy.CurrentHP > 0 {
			g.enemy.CurrentHP -= library.BasicAttackDamage(&g.player)
			if g.enemy.CurrentHP < 0 {
				g.enemy.CurrentHP = 0
			}
		}
		if g.player.Summon != nil && g.enemy.CurrentHP > 0 {
			library.SummonAttack(&g.player, g.enemy)
		}
		if g.enemy.CurrentHP <= 0 {
			drops := []string{}
			if g.combatAdvancesRoom {
				g.player.LastClearedRoom++
			}
			if g.combatAdvancesRoom {
				library.GainExperience(&g.player, g.enemy.XPReward)
				g.player.Gold += g.enemy.GoldReward
				drops = library.DropMonsterEquipment(&g.player, g.enemy)
			}
			g.combatMessage = fmt.Sprintf("Victoire ! +%d XP, +%d or.", g.enemy.XPReward, g.enemy.GoldReward)
			if len(drops) > 0 {
				g.combatMessage += fmt.Sprintf(" Loot : %v.", drops)
			}
			g.message = g.combatMessage
			g.enemy = nil
			g.combatResult = true
			return nil
		}
		library.ApplyMonsterEffects(g.enemy)
		if g.enemy.CurrentHP <= 0 {
			g.message = "L'ennemi succombe à ses effets."
			g.enemy = nil
			g.combatResult = true
			return nil
		}
		damage = library.MonsterPatternDamage(g.enemy, g.combatTurn)
		if library.CheckHolyBarrier(&g.player) || library.TryDodge(&g.player) {
			damage = 0
		}
		damage = library.BlockDamage(&g.player, damage)
		g.player.CurrentHP -= damage
		if g.player.CurrentHP < 0 {
			g.player.CurrentHP = 0
		}
		g.combatMessage = fmt.Sprintf("Vous infligez %d dégâts. %s riposte pour %d.", library.BasicAttackDamage(&g.player), g.enemy.Name, damage)
		g.combatTurn++
		library.TryCounterAttack(&g.player, g.enemy)
		if g.player.CurrentHP <= 0 {
			g.player.CurrentHP = g.player.MaxHP / 2
			g.enemy = nil
			g.combatMessage = "Défaite. Vous revenez au tableau de bord."
			g.message = g.combatMessage
			g.combatResult = true
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.Key4) || (enterPressed() && g.combatAction == 3) {
		g.enemy = nil
		g.combatMessage = "Vous quittez le combat."
		g.message = g.combatMessage
		g.combatResult = true
	}
	return nil
}

func (g *Game) updateCombatSkills() error {
	if g.enemy == nil {
		g.view = viewDashboard
		return nil
	}
	if len(g.player.Skill) == 0 {
		g.view = viewCombat
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.skillSelected = (g.skillSelected + len(g.player.Skill) - 1) % len(g.player.Skill)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.skillSelected = (g.skillSelected + 1) % len(g.player.Skill)
	}
	for key := ebiten.Key1; key <= ebiten.Key9; key++ {
		if inpututil.IsKeyJustPressed(key) {
			index := int(key - ebiten.Key1)
			if index < len(g.player.Skill) {
				g.skillSelected = index
			}
		}
	}
	if enterPressed() {
		if g.castSelectedSkill() {
			g.view = viewCombat
			g.combatTurn++
			g.enemyTurn()
		}
	}
	return nil
}

func (g *Game) updateCombatInventory() error {
	usable := make([]int, 0)
	for index, item := range g.player.Inventory {
		if item.Name == "Potion de vie" || item.Name == "Potion de mana" || item.Name == "Potion de poison" {
			usable = append(usable, index)
		}
	}
	if len(usable) == 0 {
		g.view = viewCombat
		g.combatMessage = "Aucune potion utilisable."
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.combatInventorySelected = (g.combatInventorySelected + len(usable) - 1) % len(usable)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.combatInventorySelected = (g.combatInventorySelected + 1) % len(usable)
	}
	if enterPressed() {
		index := usable[g.combatInventorySelected]
		switch g.player.Inventory[index].Name {
		case "Potion de vie":
			library.TakePot(&g.player, index)
			g.combatMessage = "Potion de vie utilisée."
		case "Potion de mana":
			library.TakeManaPot(&g.player, index)
			g.combatMessage = "Potion de mana utilisée."
		case "Potion de poison":
			library.PoisonPot(&g.player, g.enemy, index)
			g.combatMessage = "L'ennemi est empoisonné."
		}
		g.view = viewCombat
		g.combatTurn++
		g.enemyTurn()
	}
	return nil
}

func (g *Game) castSelectedSkill() bool {
	skill := g.player.Skill[g.skillSelected]
	if g.player.Mana < skill.ManaCost {
		g.combatMessage = "Mana insuffisante."
		return false
	}
	g.player.Mana -= skill.ManaCost
	switch skill.Name {
	case "Soin":
		heal := skill.HealAmount
		if g.player.Class == "Clerc" {
			heal = heal * 3 / 2
		}
		g.player.CurrentHP += heal
		if g.player.CurrentHP > g.player.MaxHP {
			g.player.CurrentHP = g.player.MaxHP
		}
		g.combatMessage = fmt.Sprintf("Soin : +%d PV.", heal)
	case "Régénération":
		g.player.Effects = append(g.player.Effects, library.Effect{Name: "Régénération", Value: skill.HealAmount, TurnsLeft: skill.EffectTurns})
		g.combatMessage = "Régénération active."
	case "Poison", "Brûlure":
		damage := library.SpellDamage(&g.player, skill.Damage)
		g.enemy.Effects = append(g.enemy.Effects, library.Effect{Name: skill.Name, Value: damage, TurnsLeft: skill.EffectTurns})
		g.combatMessage = fmt.Sprintf("%s appliqué.", skill.Name)
	case "Éclair":
		damage := library.SpellDamage(&g.player, skill.Damage)
		g.enemy.CurrentHP -= damage
		g.enemy.Initiative -= 15
		g.combatMessage = fmt.Sprintf("Éclair inflige %d dégâts.", damage)
	case "Barrière Sacrée":
		g.player.HolyBarrier = true
		g.combatMessage = "Barrière sacrée active."
	case "Invocation de soldat":
		if g.player.Summon != nil && g.player.Summon.CurrentHP > 0 {
			g.player.Mana += skill.ManaCost
			g.combatMessage = "Un soldat est déjà invoqué."
			return false
		}
		hp, attack := library.SummonStats(&g.player)
		g.player.Summon = &library.Monster{Name: "Soldat invoqué", Pattern: "summon", Level: g.player.Level, MaxHP: hp, CurrentHP: hp, Attack: attack}
		g.combatMessage = "Soldat invoqué."
	default:
		damage := library.SpellDamage(&g.player, skill.Damage)
		g.enemy.CurrentHP -= damage
		g.combatMessage = fmt.Sprintf("%s inflige %d dégâts.", skill.Name, damage)
	}
	if g.enemy.CurrentHP < 0 {
		g.enemy.CurrentHP = 0
	}
	if g.enemy.CurrentHP == 0 {
		drops := []string{}
		if g.combatAdvancesRoom {
			g.player.LastClearedRoom++
		}
		if g.combatAdvancesRoom {
			library.GainExperience(&g.player, g.enemy.XPReward)
			g.player.Gold += g.enemy.GoldReward
			drops = library.DropMonsterEquipment(&g.player, g.enemy)
		}
		g.message = fmt.Sprintf("Victoire ! +%d XP, +%d or.", g.enemy.XPReward, g.enemy.GoldReward)
		if len(drops) > 0 {
			g.message += fmt.Sprintf(" Loot : %v.", drops)
		}
		g.enemy = nil
		g.combatResult = true
	}
	return true
}

func (g *Game) enemyTurn() {
	if g.enemy == nil {
		return
	}
	library.ApplyMonsterEffects(g.enemy)
	if g.enemy.CurrentHP <= 0 {
		return
	}
	damage := library.MonsterPatternDamage(g.enemy, g.combatTurn)
	if library.CheckHolyBarrier(&g.player) || library.TryDodge(&g.player) {
		damage = 0
	}
	damage = library.BlockDamage(&g.player, damage)
	g.player.CurrentHP -= damage
	if g.player.CurrentHP < 0 {
		g.player.CurrentHP = 0
	}
	if g.player.CurrentHP <= 0 {
		g.player.CurrentHP = g.player.MaxHP / 2
		g.enemy = nil
		g.message = "Défaite. Vous revenez au tableau de bord."
		g.combatResult = true
	}
}

func (g *Game) updateMerchant() error {
	items := []struct {
		name  string
		price int
	}{
		{"Potion de vie", 3},
		{"Potion de poison", 6},
		{"Potion de mana", 6},
		{"Livre de Sort : Boule de Feu", 25},
		{"Livre de Sort : Soin", 25},
		{"Livre de Sort : Régénération", 25},
		{"Livre de Sort : Poison", 25},
		{"Livre de Sort : Brûlure", 25},
		{"Livre de Sort : Éclair", 25},
		{"Livre de Sort : Barrière Sacrée", 25},
		{"Fourrure de Loup", 4},
		{"Peau de Troll", 7},
		{"Cuir de Sanglier", 3},
		{"Plume de Corbeau", 1},
		{"Amelioration d'inventaire", 30},
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.merchantSelected = (g.merchantSelected + len(items) - 1) % len(items)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.merchantSelected = (g.merchantSelected + 1) % len(items)
	}
	if enterPressed() {
		item := items[g.merchantSelected]
		if library.BuyItem(&g.player, item.name, item.price) {
			g.message = fmt.Sprintf("Acheté : %s.", item.name)
		} else {
			g.message = "Achat impossible : or ou inventaire insuffisant."
		}
	}
	return nil
}

func (g *Game) updateInventory() error {
	if len(g.player.Inventory) == 0 {
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.inventorySelected = (g.inventorySelected + len(g.player.Inventory) - 1) % len(g.player.Inventory)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.inventorySelected = (g.inventorySelected + 1) % len(g.player.Inventory)
	}
	if enterPressed() && g.inventorySelected < len(g.player.Inventory) {
		g.useInventoryItem(g.inventorySelected)
		if len(g.player.Inventory) == 0 {
			g.inventorySelected = 0
		} else if g.inventorySelected >= len(g.player.Inventory) {
			g.inventorySelected = len(g.player.Inventory) - 1
		}
	}
	return nil
}

func (g *Game) useInventoryItem(index int) {
	if index < 0 || index >= len(g.player.Inventory) {
		return
	}
	item := g.player.Inventory[index]
	switch item.Name {
	case "Potion de vie":
		library.TakePot(&g.player, index)
		g.message = "Potion de vie utilisée."
	case "Potion de mana":
		library.TakeManaPot(&g.player, index)
		g.message = "Potion de mana utilisée."
	case "Livre de Sort : Boule de Feu", "Livre de Sort : Soin", "Livre de Sort : Régénération", "Livre de Sort : Poison", "Livre de Sort : Brûlure", "Livre de Sort : Éclair", "Livre de Sort : Barrière Sacrée":
		library.LearnSpellBook(&g.player, item.Name)
		consumeInventoryItem(&g.player, index)
		g.message = fmt.Sprintf("Livre appris : %s.", item.Name)
	default:
		if isFrontEquipment(item.Name) {
			library.EquipItemDirect(&g.player, library.ItemDisplayName(item), index)
			g.message = fmt.Sprintf("Équipement activé : %s.", item.Name)
		} else {
			g.message = "Cet objet n'est pas utilisable ici."
		}
	}
}

func isFrontEquipment(name string) bool {
	switch name {
	case "Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier", "Casque de gobelin", "Carapace de slime", "Capuche spectrale", "Casque de golem", "Peau de troll renforcée", "Plumes du canard", "Heaume squelette", "Fourrure du loup", "Chapeau du sorcier", "Écailles de dragon", "Dague de l'assassin", "Arc du chasseur", "Marteau du guerrier", "Épée du chevalier", "Lame de gobelin", "Bave de slime", "Lame spectrale", "Marteau de golem", "Massue de troll", "Bec du canard", "Épée squelette", "Crocs du loup", "Bâton maudit", "Griffe du dragon", "Tunique de gobelin", "Bottes de gobelin", "Tunique de slime", "Bottes de slime", "Tunique spectrale", "Bottes spectrales", "Tunique de golem", "Bottes de golem", "Tunique de troll", "Bottes de troll", "Tunique du canard", "Bottes du canard", "Tunique du squelette", "Bottes du squelette", "Tunique du loup", "Bottes du loup", "Tunique du sorcier", "Bottes du sorcier", "Tunique du dragon", "Bottes du dragon":
		return true
	default:
		return false
	}
}

func consumeInventoryItem(player *library.Character, index int) {
	player.Inventory[index].Quantity--
	if player.Inventory[index].Quantity <= 0 {
		player.Inventory = append(player.Inventory[:index], player.Inventory[index+1:]...)
	}
}

func (g *Game) updateForge() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.forgeSelected = (g.forgeSelected + 2) % 3
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.forgeSelected = (g.forgeSelected + 1) % 3
	}
	if enterPressed() {
		recettes := []string{"Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier"}
		library.FabriquerObjet(&g.player, recettes[g.forgeSelected])
		g.message = "Tentative de fabrication effectuée."
	}
	return nil
}

func (g *Game) updateEnchanter() error {
	if len(g.player.Skill) == 0 {
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.enchanterSelected = (g.enchanterSelected + len(g.player.Skill) - 1) % len(g.player.Skill)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.enchanterSelected = (g.enchanterSelected + 1) % len(g.player.Skill)
	}
	if enterPressed() {
		if library.EnchantSkill(&g.player, g.enchanterSelected) {
			g.message = "Sort enchanté : dégâts +5."
		} else {
			g.message = "Enchantement impossible : or insuffisant ou sort non compatible."
		}
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
	if g.view == viewCombat {
		g.drawCombat(screen)
		return
	}
	if g.view == viewCombatSkills {
		g.drawCombatSkills(screen)
		return
	}
	if g.view == viewCombatInventory {
		g.drawCombatInventory(screen)
		return
	}
	if g.view == viewMerchant {
		g.drawMerchant(screen)
		return
	}
	if g.view == viewForge {
		g.drawForge(screen)
		return
	}
	if g.view == viewEnchanter {
		g.drawEnchanter(screen)
		return
	}

	g.drawDashboard(screen)
}

var (
	panelColor    = color.RGBA{R: 22, G: 29, B: 43, A: 235}
	panelColorTop = color.RGBA{R: 32, G: 41, B: 58, A: 235}
	panelEdge     = color.RGBA{R: 168, G: 138, B: 82, A: 255}
	goldColor     = color.RGBA{R: 238, G: 184, B: 74, A: 255}
	blueColor     = color.RGBA{R: 76, G: 164, B: 235, A: 255}
	redColor      = color.RGBA{R: 226, G: 83, B: 91, A: 255}
)

func drawPanel(screen *ebiten.Image, x, y, width, height int) {
	fx, fy, fw, fh := float32(x), float32(y), float32(width), float32(height)
	ebitenutil.DrawRect(screen, float64(x+4), float64(y+4), float64(width), float64(height), color.RGBA{R: 0, G: 0, B: 0, A: 150})
	ebitenutil.DrawRect(screen, float64(x), float64(y), float64(width), float64(height), panelColor)
	ebitenutil.DrawRect(screen, float64(x), float64(y), float64(width), float64(height)/3, panelColorTop)
	vector.StrokeRect(screen, fx, fy, fw, fh, 2, panelEdge, true)
	corner := float32(12)
	for _, c := range [][2]float32{{fx, fy}, {fx + fw, fy}, {fx, fy + fh}, {fx + fw, fy + fh}} {
		dx, dy := corner, corner
		if c[0] > fx {
			dx = -corner
		}
		if c[1] > fy {
			dy = -corner
		}
		vector.StrokeLine(screen, c[0], c[1], c[0]+dx, c[1], 3, goldColor, true)
		vector.StrokeLine(screen, c[0], c[1], c[0], c[1]+dy, 3, goldColor, true)
	}
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
	ebitenutil.DrawRect(screen, float64(x), float64(y), float64(width)*ratio, float64(height)/2, color.RGBA{R: 255, G: 255, B: 255, A: 60})
	vector.StrokeRect(screen, float32(x), float32(y), float32(width), float32(height), 1, color.RGBA{R: 10, G: 12, B: 16, A: 255}, true)
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
	g.drawDungeonBackdrop(screen)
	g.drawLogo(screen)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s  //  %s", g.player.Name, g.player.Class), 500, 28)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("NIVEAU %d", g.player.Level), 590, 43)

	drawPanel(screen, 28, 75, 300, 142)
	ebitenutil.DebugPrintAt(screen, "PROFIL DU HÉROS", 44, 91)
	ebitenutil.DebugPrintAt(screen, g.player.Name, 44, 111)
	ebitenutil.DebugPrintAt(screen, g.player.Class, 44, 126)
	drawCharacterSprite(screen, 264, 95, 4, g.player.Name, g.player.Class)
	drawBar(screen, 44, 151, 210, 10, g.player.CurrentHP, g.player.MaxHP, redColor)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("PV  %d / %d", g.player.CurrentHP, g.player.MaxHP), 44, 164)
	drawBar(screen, 44, 181, 210, 10, g.player.Mana, g.player.MaxMana, blueColor)
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

func brickShade(col, row int) color.RGBA {
	h := uint32(col*928371+row*123457) % 16
	base := uint8(30 + h)
	return color.RGBA{R: base, G: base + 5, B: base + 14, A: 255}
}

func drawTorch(screen *ebiten.Image, x, y int, frame int) {
	flicker := float32(math.Sin(float64(frame)*0.18+float64(x))) * 3
	fx, fy := float32(x), float32(y)
	vector.DrawFilledCircle(screen, fx, fy, 20+flicker, color.RGBA{R: 255, G: 130, B: 40, A: 35}, true)
	vector.DrawFilledCircle(screen, fx, fy, 12+flicker*0.6, color.RGBA{R: 255, G: 160, B: 60, A: 70}, true)
	ebitenutil.DrawRect(screen, float64(x-2), float64(y+4), 4, 16, color.RGBA{R: 44, G: 34, B: 26, A: 255})
	ebitenutil.DrawRect(screen, float64(x-6), float64(y), 12, 5, color.RGBA{R: 60, G: 46, B: 34, A: 255})
	vector.DrawFilledCircle(screen, fx, fy-2+flicker*0.2, 6, color.RGBA{R: 255, G: 110, B: 30, A: 255}, true)
	vector.DrawFilledCircle(screen, fx, fy-4+flicker*0.3, 3, color.RGBA{R: 255, G: 224, B: 130, A: 255}, true)
}

func (g *Game) drawDungeonBackdrop(screen *ebiten.Image) {
	const width, height = 700, 450
	const wallHeight = 270
	ebitenutil.DrawRect(screen, 0, 0, width, wallHeight, color.RGBA{R: 26, G: 24, B: 28, A: 255})
	brickW, brickH := 44, 24
	for row := 0; row*brickH < wallHeight; row++ {
		offset := 0
		if row%2 == 1 {
			offset = brickW / 2
		}
		for col := -1; col*brickW < width+brickW; col++ {
			bx := col*brickW + offset
			by := row * brickH
			ebitenutil.DrawRect(screen, float64(bx+1), float64(by+1), float64(brickW-2), float64(brickH-2), brickShade(col, row))
		}
	}
	ebitenutil.DrawRect(screen, 0, float64(wallHeight-6), width, 6, color.RGBA{R: 12, G: 12, B: 16, A: 200})

	ebitenutil.DrawRect(screen, 0, wallHeight, width, height-wallHeight, color.RGBA{R: 34, G: 27, B: 22, A: 255})
	for x := 0; x < width; x += 58 {
		shade := color.RGBA{R: 42, G: 33, B: 27, A: 255}
		if (x/58)%2 == 0 {
			shade = color.RGBA{R: 38, G: 30, B: 24, A: 255}
		}
		ebitenutil.DrawRect(screen, float64(x+1), float64(wallHeight+1), 56, float64(height-wallHeight-2), shade)
	}
	for y := wallHeight; y < height; y += 40 {
		ebitenutil.DrawRect(screen, 0, float64(y), width, 1, color.RGBA{R: 18, G: 14, B: 12, A: 150})
	}

	for _, tx := range []int{90, 350, 610} {
		drawTorch(screen, tx, 44, g.frameCount)
	}

	vignetteLayers := 5
	for i := 0; i < vignetteLayers; i++ {
		alpha := uint8(26 - i*5)
		inset := float64(i * 6)
		shadow := color.RGBA{R: 0, G: 0, B: 0, A: alpha}
		ebitenutil.DrawRect(screen, inset, inset, width-2*inset, 4, shadow)
		ebitenutil.DrawRect(screen, inset, height-inset-4, width-2*inset, 4, shadow)
		ebitenutil.DrawRect(screen, inset, inset, 4, height-2*inset, shadow)
		ebitenutil.DrawRect(screen, width-inset-4, inset, 4, height-2*inset, shadow)
	}
}

func (g *Game) drawCombat(screen *ebiten.Image) {
	g.drawDungeonBackdrop(screen)
	g.drawLogo(screen)
	drawPanel(screen, 28, 78, 300, 140)
	drawPanel(screen, 344, 78, 328, 140)
	drawPanel(screen, 28, 236, 644, 180)
	if g.enemy == nil {
		ebitenutil.DebugPrintAt(screen, "COMBAT TERMINÉ", 48, 100)
		ebitenutil.DebugPrintAt(screen, g.combatMessage, 48, 125)
		ebitenutil.DebugPrintAt(screen, "Entrée ou Échap : retour au tableau de bord", 48, 180)
		return
	}
	ebitenutil.DebugPrintAt(screen, "VOUS", 48, 95)
	ebitenutil.DebugPrintAt(screen, g.player.Name, 48, 113)
	drawCharacterSprite(screen, 246, 88, 5, g.player.Name, g.player.Class)
	drawBar(screen, 48, 140, 190, 14, g.player.CurrentHP, g.player.MaxHP, redColor)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("PV %d / %d", g.player.CurrentHP, g.player.MaxHP), 48, 160)
	ebitenutil.DebugPrintAt(screen, "ENNEMI", 364, 95)
	ebitenutil.DebugPrintAt(screen, g.enemy.Name, 364, 113)
	drawMonsterSprite(screen, 528, 88, 5, g.enemy.Pattern)
	drawBar(screen, 364, 140, 155, 14, g.enemy.CurrentHP, g.enemy.MaxHP, redColor)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("PV %d / %d", g.enemy.CurrentHP, g.enemy.MaxHP), 364, 160)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TOUR %d", g.combatTurn), 590, 95)
	ebitenutil.DebugPrintAt(screen, "ACTIONS", 48, 255)
	attackCursor, spellCursor, inventoryCursor, fleeCursor := "  ", "  ", "  ", "  "
	switch g.combatAction {
	case 0:
		attackCursor = ">>"
	case 1:
		spellCursor = ">>"
	case 2:
		inventoryCursor = ">>"
	case 3:
		fleeCursor = ">>"
	}
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s 1  Attaque de base", attackCursor), 48, 280)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s 2  Ouvrir le grimoire", spellCursor), 48, 300)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s 3  Inventaire", inventoryCursor), 48, 320)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s 4  Abandonner", fleeCursor), 48, 340)
	ebitenutil.DebugPrintAt(screen, g.combatMessage, 48, 360)
	ebitenutil.DebugPrintAt(screen, "Flèches : sélectionner     Entrée : valider     Échap : tableau de bord", 48, 390)
}

func (g *Game) drawCombatSkills(screen *ebiten.Image) {
	g.drawDungeonBackdrop(screen)
	g.drawLogo(screen)
	drawPanel(screen, 28, 78, 644, 338)
	ebitenutil.DebugPrintAt(screen, "GRIMOIRE // SORTS DE COMBAT", 48, 98)
	if len(g.player.Skill) == 0 {
		ebitenutil.DebugPrintAt(screen, "Aucun sort connu.", 48, 145)
	} else {
		for index, skill := range g.player.Skill {
			cursor := "  "
			if index == g.skillSelected {
				cursor = ">>"
			}
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s %d  %-24s mana %d", cursor, index+1, skill.Name, skill.ManaCost), 48, 145+index*22)
		}
	}
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("MANA : %d / %d", g.player.Mana, g.player.MaxMana), 48, 350)
	ebitenutil.DebugPrintAt(screen, g.combatMessage, 48, 370)
	ebitenutil.DebugPrintAt(screen, "Flèches + Entrée : lancer     Échap : combat", 48, 395)
}

func (g *Game) drawCombatInventory(screen *ebiten.Image) {
	g.drawDungeonBackdrop(screen)
	g.drawLogo(screen)
	drawPanel(screen, 28, 78, 644, 338)
	ebitenutil.DebugPrintAt(screen, "INVENTAIRE // COMBAT", 48, 98)
	usable := make([]library.Item, 0)
	for _, item := range g.player.Inventory {
		if item.Name == "Potion de vie" || item.Name == "Potion de mana" || item.Name == "Potion de poison" {
			usable = append(usable, item)
		}
	}
	for index, item := range usable {
		cursor := "  "
		if index == g.combatInventorySelected {
			cursor = ">>"
		}
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s %d  %s x%d", cursor, index+1, item.Name, item.Quantity), 48, 145+index*24)
	}
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("PV %d / %d     MANA %d / %d", g.player.CurrentHP, g.player.MaxHP, g.player.Mana, g.player.MaxMana), 48, 300)
	ebitenutil.DebugPrintAt(screen, "Flèches : sélectionner     Entrée : utiliser     Échap : combat", 48, 380)
}

func (g *Game) drawMerchant(screen *ebiten.Image) {
	g.drawDungeonBackdrop(screen)
	g.drawLogo(screen)
	drawPanel(screen, 28, 78, 644, 338)
	ebitenutil.DebugPrintAt(screen, "MARCHAND // COMPTOIR", 48, 98)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("OR DISPONIBLE : %d", g.player.Gold), 48, 122)
	merchantItems := []string{"Potion de vie - 3 or", "Potion de poison - 6 or", "Potion de mana - 6 or", "Livre de Sort : Boule de Feu - 25 or", "Livre de Sort : Soin - 25 or", "Livre de Sort : Régénération - 25 or", "Livre de Sort : Poison - 25 or", "Livre de Sort : Brûlure - 25 or", "Livre de Sort : Éclair - 25 or", "Livre de Sort : Barrière Sacrée - 25 or", "Fourrure de Loup - 4 or", "Peau de Troll - 7 or", "Cuir de Sanglier - 3 or", "Plume de Corbeau - 1 or", "Amelioration d'inventaire - 30 or"}
	for index, item := range merchantItems {
		cursor := "  "
		if index == g.merchantSelected {
			cursor = ">>"
		}
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s %02d  %s", cursor, index+1, item), 48, 145+index*15)
	}
	ebitenutil.DebugPrintAt(screen, g.message, 48, 365)
	ebitenutil.DebugPrintAt(screen, "Flèches : sélectionner     Entrée : acheter     Échap : retour", 48, 380)
}

func (g *Game) drawForge(screen *ebiten.Image) {
	g.drawDungeonBackdrop(screen)
	g.drawLogo(screen)
	drawPanel(screen, 28, 78, 644, 338)
	ebitenutil.DebugPrintAt(screen, "FORGERON // ÉTABLI", 48, 98)
	forgeCursor := func(index int) string {
		if index == g.forgeSelected {
			return ">>"
		}
		return "  "
	}
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s 1  Chapeau de l'aventurier       5 or", forgeCursor(0)), 48, 145)
	ebitenutil.DebugPrintAt(screen, "   Requis : 1 Plume de Corbeau + 1 Cuir de Sanglier", 48, 162)
	drawResourceSprite(screen, 560, 143, 3, "Plume de Corbeau")
	drawResourceSprite(screen, 585, 143, 3, "Cuir de Sanglier")
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s 2  Tunique de l'aventurier       5 or", forgeCursor(1)), 48, 187)
	ebitenutil.DebugPrintAt(screen, "   Requis : 2 Fourrures de Loup + 1 Peau de Troll", 48, 204)
	drawResourceSprite(screen, 560, 185, 3, "Fourrure de Loup")
	drawResourceSprite(screen, 585, 185, 3, "Peau de Troll")
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s 3  Bottes de l'aventurier        5 or", forgeCursor(2)), 48, 229)
	ebitenutil.DebugPrintAt(screen, "   Requis : 1 Fourrure de Loup + 1 Cuir de Sanglier", 48, 246)
	drawResourceSprite(screen, 560, 227, 3, "Fourrure de Loup")
	drawResourceSprite(screen, 585, 227, 3, "Cuir de Sanglier")
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("OR : %d", g.player.Gold), 48, 275)
	ebitenutil.DebugPrintAt(screen, g.message, 48, 300)
	ebitenutil.DebugPrintAt(screen, "Flèches : sélectionner     Entrée : fabriquer     Échap : retour", 48, 380)
}

func (g *Game) drawEnchanter(screen *ebiten.Image) {
	g.drawDungeonBackdrop(screen)
	g.drawLogo(screen)
	drawPanel(screen, 28, 78, 644, 338)
	ebitenutil.DebugPrintAt(screen, "ENCHANTEUR // AUTEL", 48, 98)
	if len(g.player.Skill) == 0 {
		ebitenutil.DebugPrintAt(screen, "Aucun sort connu.", 48, 145)
	} else {
		for index, skill := range g.player.Skill {
			cost := 20 + skill.Damage*2
			cursor := "  "
			if index == g.enchanterSelected {
				cursor = ">>"
			}
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s %d  %-24s dégâts %d  coût %d or", cursor, index+1, skill.Name, skill.Damage, cost), 48, 145+index*22)
		}
	}
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("OR : %d", g.player.Gold), 48, 280)
	ebitenutil.DebugPrintAt(screen, g.message, 48, 310)
	ebitenutil.DebugPrintAt(screen, "Flèches : sélectionner     Entrée : enchanter     Échap : retour", 48, 380)
}

func dungeonFloorForUI(room int) int {
	if room < 1 {
		return 1
	}
	return (room-1)/10 + 1
}

func (g *Game) drawStart(screen *ebiten.Image) {
	g.drawDungeonBackdrop(screen)
	g.drawLogo(screen)
	drawPanel(screen, 28, 78, 644, 338)
	first := "> "
	second := "  "
	if g.confirmSelected == 1 {
		first, second = "  ", "> "
	}
	text := fmt.Sprintf("LOOTSTORM / NOUVELLE PARTIE\n\n%sChoisir un personnage prédéfini\n%sCréer mon personnage\n\nFlèches haut/bas + Entrée\nÉchap : quitter", first, second)
	ebitenutil.DebugPrintAt(screen, text, 48, 98)
}

func (g *Game) drawHeroSelection(screen *ebiten.Image) {
	g.drawDungeonBackdrop(screen)
	g.drawLogo(screen)
	drawPanel(screen, 28, 78, 644, 338)
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
	text += fmt.Sprintf("\n=== %s ===\nClasse : %s\nPV : %d\nAttaque : %d\nInitiative : %d\n\n%s\n\nFlèches haut/bas + Entrée : sélectionner\nÉchap : quitter", selectedHero.Name, selectedHero.Class, selectedHero.MaxHP, selectedHero.Attack, selectedHero.Initiative, library.ClassAdvantages(selectedHero.Class))
	ebitenutil.DebugPrintAt(screen, text, 48, 98)
	drawCharacterSprite(screen, 510, 98, 8, selectedHero.Name, selectedHero.Class)
}

func (g *Game) drawHeroConfirmation(screen *ebiten.Image) {
	g.drawDungeonBackdrop(screen)
	g.drawLogo(screen)
	drawPanel(screen, 28, 78, 644, 338)
	hero := heroChoices()[g.heroSelected]
	text := fmt.Sprintf("LOOTSTORM / CONFIRMATION\n\n%s\nClasse : %s\nPV : %d\nAttaque : %d\nInitiative : %d\n\n%s\n\nFlèches : valider ou revenir\nEntrée : confirmer le choix", hero.Name, hero.Class, hero.MaxHP, hero.Attack, hero.Initiative, library.ClassAdvantages(hero.Class))
	ebitenutil.DebugPrintAt(screen, text, 48, 98)
	drawCharacterSprite(screen, 510, 98, 8, hero.Name, hero.Class)
}

func (g *Game) drawNameInput(screen *ebiten.Image) {
	g.drawDungeonBackdrop(screen)
	g.drawLogo(screen)
	drawPanel(screen, 28, 78, 644, 338)
	text := fmt.Sprintf("LOOTSTORM / CRÉATION\n\nChoisissez le nom de votre personnage :\n\n> %s_\n\nEntrée : continuer\nRetour arrière : effacer\nÉchap : retour", g.customName)
	ebitenutil.DebugPrintAt(screen, text, 48, 98)
}

func (g *Game) drawClassSelection(screen *ebiten.Image) {
	g.drawDungeonBackdrop(screen)
	g.drawLogo(screen)
	drawPanel(screen, 28, 78, 644, 338)
	classes := classChoices()
	text := fmt.Sprintf("LOOTSTORM / CHOIX DE CLASSE\n\nNom : %s\n\n", g.customName)
	for index, className := range classes {
		cursor := "  "
		if index == g.classSelected {
			cursor = "> "
		}
		text += fmt.Sprintf("%s%d. %-12s PV %d\n", cursor, index+1, className, getHPForClass(className))
	}
	text += fmt.Sprintf("\nAvantage : %s\n\nFlèches haut/bas + Entrée : sélectionner", library.ClassAdvantages(classes[g.classSelected]))
	ebitenutil.DebugPrintAt(screen, text, 48, 98)
	drawCharacterSprite(screen, 510, 98, 8, g.customName, classes[g.classSelected])
}

func (g *Game) drawClassConfirmation(screen *ebiten.Image) {
	g.drawDungeonBackdrop(screen)
	g.drawLogo(screen)
	drawPanel(screen, 28, 78, 644, 338)
	className := classChoices()[g.classSelected]
	text := fmt.Sprintf("LOOTSTORM / CONFIRMATION DE CLASSE\n\nNom : %s\nClasse : %s\nPV de base : %d\n\n%s\n\nFlèches : valider ou revenir\nEntrée : confirmer la classe", g.customName, className, getHPForClass(className), library.ClassAdvantages(className))
	ebitenutil.DebugPrintAt(screen, text, 48, 98)
	drawCharacterSprite(screen, 510, 98, 8, g.customName, className)
}

func (g *Game) drawStats(screen *ebiten.Image) {
	g.drawDungeonBackdrop(screen)
	g.drawLogo(screen)
	p := g.player

	drawPanel(screen, 28, 78, 300, 142)
	ebitenutil.DebugPrintAt(screen, "IDENTITÉ", 44, 94)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Nom     %s", p.Name), 44, 114)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Classe  %s", p.Class), 44, 129)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Niveau  %d", p.Level), 44, 144)
	drawCharacterSprite(screen, 216, 96, 4, p.Name, p.Class)

	drawPanel(screen, 344, 78, 328, 142)
	ebitenutil.DebugPrintAt(screen, "PROGRESSION", 360, 94)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Étage %d   //   Salle %d", dungeonFloorForUI(p.LastClearedRoom+1), p.LastClearedRoom+1), 360, 114)
	drawBar(screen, 360, 128, 280, 12, p.CurrentXP, p.MaxXP, goldColor)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("XP  %d / %d", p.CurrentXP, p.MaxXP), 360, 148)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Inventaire  %d / %d     Améliorations  %d / 3", len(p.Inventory), p.LimitInventory, p.LimitInventoryUpgrade), 360, 168)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Or  %d", p.Gold), 360, 188)

	drawPanel(screen, 28, 228, 300, 180)
	ebitenutil.DebugPrintAt(screen, "STATISTIQUES DE COMBAT", 44, 244)
	drawBar(screen, 44, 258, 250, 12, p.CurrentHP, p.MaxHP, redColor)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("PV  %d / %d", p.CurrentHP, p.MaxHP), 44, 278)
	drawBar(screen, 44, 292, 250, 12, p.Mana, p.MaxMana, blueColor)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("MANA  %d / %d", p.Mana, p.MaxMana), 44, 312)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("ATTAQUE      %d", p.Attack+p.Equip.WeaponDamage), 44, 334)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("INITIATIVE   %d", p.Initiative), 44, 349)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("SLOTS LIBRES %d", p.LimitInventory-len(p.Inventory)), 44, 364)

	drawPanel(screen, 344, 228, 328, 180)
	ebitenutil.DebugPrintAt(screen, "ÉQUIPEMENT ACTIF", 360, 244)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("ARME       %s (+%d ATQ)", p.Equip.Weapon, p.Equip.WeaponDamage), 360, 266)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("CASQUE     %s%s", p.Equip.Helmet, library.ItemStatSummary(p.Equip.Helmet)), 360, 286)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("PLASTRON   %s%s", p.Equip.Chestplate, library.ItemStatSummary(p.Equip.Chestplate)), 360, 306)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("BOTTES     %s%s", p.Equip.Boots, library.ItemStatSummary(p.Equip.Boots)), 360, 326)
	ebitenutil.DebugPrintAt(screen, wrapText(library.ClassAdvantages(p.Class), 42), 360, 350)

	ebitenutil.DebugPrintAt(screen, "Échap : retour", 28, 432)
}

// wrapText breaks text into lines of at most maxChars, splitting on spaces so
// long sentences (like class advantage descriptions) don't overflow a panel.
func wrapText(text string, maxChars int) string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return text
	}
	var lines []string
	line := words[0]
	for _, word := range words[1:] {
		if len(line)+1+len(word) > maxChars {
			lines = append(lines, line)
			line = word
			continue
		}
		line += " " + word
	}
	lines = append(lines, line)
	return strings.Join(lines, "\n")
}

func (g *Game) drawInventory(screen *ebiten.Image) {
	g.drawDungeonBackdrop(screen)
	g.drawLogo(screen)
	drawPanel(screen, 28, 78, 644, 338)
	ebitenutil.DebugPrintAt(screen, "INVENTAIRE // SACOCHE", 48, 98)
	if len(g.player.Inventory) == 0 {
		ebitenutil.DebugPrintAt(screen, "Inventaire vide.", 48, 145)
	} else {
		for index, item := range g.player.Inventory {
			if index >= 15 {
				break
			}
			cursor := "  "
			if index == g.inventorySelected {
				cursor = ">>"
			}
			drawItemSprite(screen, 365, 127+index*17, 2, item.Name)
			bonus := library.ItemStatSummary(library.ItemDisplayName(item))
			label := library.ItemDisplayName(item)
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s %02d  %-28s%s x%d", cursor, index+1, label, bonus, item.Quantity), 48, 130+index*17)
		}
	}
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Emplacements : %d / %d     Or : %d", len(g.player.Inventory), g.player.LimitInventory, g.player.Gold), 48, 385)
	ebitenutil.DebugPrintAt(screen, "Flèches : sélectionner     Entrée : utiliser / équiper     Échap : retour", 48, 402)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 700, 450
}

// DrawFinalScreen upscales the 700x450 canvas to the window. It snaps to a
// crisp nearest-neighbor blit whenever the scale is (close to) a whole
// number — which is the common case in windowed mode — and falls back to
// smooth linear filtering for the fractional scales fullscreen produces on
// arbitrary monitor resolutions, avoiding jagged edges.
func (g *Game) DrawFinalScreen(screen ebiten.FinalScreen, offscreen *ebiten.Image, geoM ebiten.GeoM) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM = geoM
	scale := geoM.Element(0, 0)
	if math.Abs(scale-math.Round(scale)) < 0.02 {
		op.Filter = ebiten.FilterNearest
	} else {
		op.Filter = ebiten.FilterLinear
	}
	screen.DrawImage(offscreen, op)
}

func main() {
	game := &Game{selected: 0}
	// Size the (bordered, non-fullscreen) window to the largest exact
	// multiple of the 700x450 canvas that fits the screen, so the final
	// upscale is always an integer ratio and stays pixel-crisp.
	windowWidth, windowHeight := 1400, 900
	if monitor := ebiten.Monitor(); monitor != nil {
		monitorWidth, monitorHeight := monitor.Size()
		scale := monitorWidth / 700
		if alt := monitorHeight / 450; alt < scale {
			scale = alt
		}
		if scale < 1 {
			scale = 1
		}
		if scale > 2 {
			scale-- // leave room for the title bar and taskbar
		}
		windowWidth, windowHeight = 700*scale, 450*scale
	}
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Lootstorm")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
