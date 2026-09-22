package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"projet/src/library"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 960
	screenHeight = 600
	windowWidth  = 1440
	windowHeight = 900
)

var theme = struct {
	background color.RGBA
	panel      color.RGBA
	panelLight color.RGBA
	gold       color.RGBA
	goldBright color.RGBA
	text       color.RGBA
	muted      color.RGBA
	ruby       color.RGBA
}{
	background: color.RGBA{R: 15, G: 13, B: 18, A: 255},
	panel:      color.RGBA{R: 29, G: 25, B: 31, A: 245},
	panelLight: color.RGBA{R: 48, G: 39, B: 39, A: 245},
	gold:       color.RGBA{R: 157, G: 119, B: 52, A: 255},
	goldBright: color.RGBA{R: 230, G: 190, B: 93, A: 255},
	text:       color.RGBA{R: 238, G: 225, B: 193, A: 255},
	muted:      color.RGBA{R: 164, G: 148, B: 125, A: 255},
	ruby:       color.RGBA{R: 132, G: 48, B: 48, A: 255},
}

type screen int

const (
	startSelection screen = iota
	nameCreation
	classCreation
	predefinedSelection
	mainMenu
	characterSheet
	inventoryScreen
	merchantScreen
	forgeScreen
	enchanterScreen
	creditsScreen
	combatScreen
	combatSpellsScreen
	combatInventoryScreen
	messageScreen
)

type classChoice struct {
	name string
	hp   int
}

var classes = []classChoice{
	{"Guerrier", 150}, {"Mage", 80}, {"Archer", 110},
	{"Assassin", 100}, {"Chevalier", 120}, {"Samourai", 115},
	{"Clerc", 105}, {"Barbare", 180}, {"Invocateur", 90},
}

type predefinedHero struct {
	name  string
	class string
	hp    int
}

var predefinedHeroes = []predefinedHero{
	{"Himiko Toga", "Assassin", 100}, {"Link", "Chevalier", 120},
	{"Frieren", "Mage", 80}, {"Musashi", "Samourai", 115},
	{"Elizabeth", "Clerc", 105}, {"Guts", "Barbare", 180},
	{"Sung Jin Woo", "Invocateur", 90}, {"Colley", "Archer", 110},
}

type shopItem struct {
	name  string
	price int
}

var shopItems = []shopItem{
	{"Potion de vie", 3}, {"Potion de poison", 6}, {"Potion de mana", 6},
	{"Livre de Sort : Boule de Feu", 25}, {"Livre de Sort : Soin", 25},
	{"Livre de Sort : Régénération", 25}, {"Livre de Sort : Poison", 25},
	{"Livre de Sort : Brûlure", 25}, {"Livre de Sort : Éclair", 25},
	{"Livre de Sort : Barrière Sacrée", 25}, {"Fourrure de Loup", 4},
	{"Peau de Troll", 7}, {"Cuir de Sanglier", 3}, {"Plume de Corbeau", 1},
	{"Amelioration d'inventaire", 30},
}

var forgeItems = []string{
	"Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier",
}

type game struct {
	currentScreen       screen
	selected            int
	name                string
	player              library.Character
	message             string
	returnScreen        screen
	monster             library.Monster
	combatTurn          int
	combatAdventure     bool
	sellMode            bool
	combatLog           string
	random              *rand.Rand
	assassinExtraAction bool
	sprites             map[string]*ebiten.Image
	background          *ebiten.Image
	combatBackground    *ebiten.Image
	merchantBackground  *ebiten.Image
	forgeBackground     *ebiten.Image
	enchanterBackground *ebiten.Image
}

func newGame() *game {
	game := &game{
		currentScreen:       startSelection,
		random:              rand.New(rand.NewSource(time.Now().UnixNano())),
		sprites:             loadSprites(),
		background:          loadBackground(),
		combatBackground:    loadBackgroundFile("combat.png"),
		merchantBackground:  loadBackgroundFile("merchant.png"),
		forgeBackground:     loadBackgroundFile("forge.png"),
		enchanterBackground: loadBackgroundFile("enchanter.png"),
	}
	currentGame = game
	return game
}

var spriteFiles = map[string]string{
	"class:Guerrier": "class_guerrier.png", "class:Mage": "class_mage.png", "class:Archer": "class_archer.png",
	"class:Assassin": "class_assassin.png", "class:Chevalier": "class_chevalier.png", "class:Samourai": "class_samourai.png",
	"class:Clerc": "class_clerc.png", "class:Barbare": "class_barbare.png", "class:Invocateur": "class_invocateur.png",
	"hero:Himiko Toga": "hero_himiko_toga.png", "hero:Link": "hero_link.png",
	"hero:Frieren": "hero_frieren.png", "hero:Musashi": "hero_musashi.png",
	"hero:Elizabeth": "hero_elizabeth.png", "hero:Guts": "hero_guts.png", "hero:Sung Jin Woo": "hero_sung_jin_woo.png",
	"hero:Colley": "hero_colley.png",
	"goblin":      "enemy_goblin.png", "slime": "enemy_slime.png", "ghost": "enemy_ghost.png",
	"golem": "enemy_golem.png", "troll": "enemy_troll.png", "skeleton": "enemy_skeleton.png",
	"wolf": "enemy_wolf.png", "wizard": "enemy_wizard.png", "duck": "enemy_duck.png",
	"dragon": "enemy_dragon.png", "summon": "enemy_summon.png",
}

func loadSprites() map[string]*ebiten.Image {
	sprites := make(map[string]*ebiten.Image)
	for key, filename := range spriteFiles {
		for _, root := range []string{"frontend/assets/sprites", "assets/sprites"} {
			path := filepath.Join(root, filename)
			if _, err := os.Stat(path); err != nil {
				continue
			}
			image, _, err := ebitenutil.NewImageFromFile(path)
			if err == nil {
				sprites[key] = image
			}
			break
		}
	}
	return sprites
}

func loadBackground() *ebiten.Image {
	return loadBackgroundFile("dungeon.png")
}

func loadBackgroundFile(filename string) *ebiten.Image {
	for _, root := range []string{"frontend/assets/backgrounds", "assets/backgrounds"} {
		path := filepath.Join(root, filename)
		if _, err := os.Stat(path); err != nil {
			continue
		}
		background, _, err := ebitenutil.NewImageFromFile(path)
		if err == nil {
			return background
		}
	}
	return nil
}

func (g *game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		if g.currentScreen == startSelection {
			return ebiten.Termination
		}
		if g.currentScreen == predefinedSelection || g.currentScreen == nameCreation || g.currentScreen == classCreation {
			g.selected = 0
			g.currentScreen = startSelection
			return nil
		}
		if g.currentScreen == combatSpellsScreen || g.currentScreen == combatInventoryScreen {
			g.currentScreen = combatScreen
			return nil
		}
		g.currentScreen = mainMenu
		return nil
	}
	switch g.currentScreen {
	case startSelection:
		g.updateStartSelection()
	case nameCreation:
		g.updateNameCreation()
	case classCreation:
		g.updateClassCreation()
	case predefinedSelection:
		g.updatePredefinedSelection()
	case mainMenu:
		return g.updateMainMenu()
	case characterSheet:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.currentScreen = mainMenu
		}
	case inventoryScreen:
		g.updateInventory()
	case merchantScreen:
		g.updateMerchant()
	case forgeScreen:
		g.updateForge()
	case enchanterScreen:
		g.updateEnchanter()
	case creditsScreen:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.currentScreen = mainMenu
		}
	case combatScreen:
		g.updateCombatMenu()
	case combatSpellsScreen:
		g.updateCombatSpells()
	case combatInventoryScreen:
		g.updateCombatInventory()
	case messageScreen:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			if g.returnScreen == combatScreen && g.combatAdventure {
				g.startAdventure()
			} else {
				g.currentScreen = g.returnScreen
			}
		}
	}
	return nil
}

func (g *game) updateStartSelection() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.selected = 1 - g.selected
	}
	if !inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return
	}
	if g.selected == 0 {
		g.currentScreen = nameCreation
	} else {
		g.selected = 0
		g.currentScreen = predefinedSelection
	}
}

func (g *game) updatePredefinedSelection() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.selected = (g.selected + 1) % len(predefinedHeroes)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.selected = (g.selected + len(predefinedHeroes) - 1) % len(predefinedHeroes)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		hero := predefinedHeroes[g.selected]
		g.player = library.InitCharacter(hero.name, hero.class, hero.hp)
		g.selected = 0
		g.currentScreen = mainMenu
	}
}

func (g *game) updateNameCreation() {
	for _, character := range ebiten.InputChars() {
		if unicode.IsLetter(character) || character == ' ' {
			g.name += string(character)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && len(g.name) > 0 {
		name := []rune(g.name)
		g.name = string(name[:len(name)-1])
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && strings.TrimSpace(g.name) != "" {
		g.name = titleName(g.name)
		g.selected = 0
		g.currentScreen = classCreation
	}
}

func titleName(name string) string {
	words := strings.Fields(strings.ToLower(name))
	for index, word := range words {
		letters := []rune(word)
		if len(letters) > 0 {
			letters[0] = unicode.ToUpper(letters[0])
		}
		words[index] = string(letters)
	}
	return strings.Join(words, " ")
}

func (g *game) updateClassCreation() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.selected = (g.selected + 1) % len(classes)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.selected = (g.selected + len(classes) - 1) % len(classes)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		choice := classes[g.selected]
		g.player = library.InitCharacter(g.name, choice.name, choice.hp)
		g.selected = 0
		g.currentScreen = mainMenu
	}
}

func (g *game) updateMainMenu() error {
	const menuLength = 9
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.selected = (g.selected + 1) % menuLength
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.selected = (g.selected + menuLength - 1) % menuLength
	}
	if !inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return nil
	}
	switch g.selected {
	case 0:
		g.currentScreen = characterSheet
	case 1:
		g.startAdventure()
	case 2:
		g.startTrainingFight()
	case 3:
		g.selected = 0
		g.currentScreen = inventoryScreen
	case 4:
		g.selected = 0
		g.currentScreen = merchantScreen
	case 5:
		g.selected = 0
		g.currentScreen = forgeScreen
	case 6:
		g.selected = 0
		g.currentScreen = enchanterScreen
	case 7:
		g.currentScreen = creditsScreen
	case 8:
		return ebiten.Termination
	}
	return nil
}

func (g *game) updateInventory() {
	if len(g.player.Inventory) == 0 {
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.selected = (g.selected + 1) % len(g.player.Inventory)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.selected = (g.selected + len(g.player.Inventory) - 1) % len(g.player.Inventory)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.useInventoryItem(g.selected)
	}
}

func (g *game) useInventoryItem(index int) {
	if index < 0 || index >= len(g.player.Inventory) {
		return
	}
	item := g.player.Inventory[index]
	switch item.Name {
	case "Potion de vie":
		library.TakePotSilent(&g.player, index)
	case "Potion de mana":
		library.TakeManaPotSilent(&g.player, index)
	case "Amelioration d'inventaire":
		if library.UpgradeInventorySlotSilent(&g.player) {
			g.player.Inventory[index].Quantity--
			if g.player.Inventory[index].Quantity <= 0 {
				g.player.Inventory = append(g.player.Inventory[:index], g.player.Inventory[index+1:]...)
			}
			g.message = fmt.Sprintf("Inventaire augmenté : %d emplacements.", g.player.LimitInventory)
		} else {
			g.message = "Les trois améliorations d'inventaire ont déjà été utilisées."
		}
		g.returnScreen = inventoryScreen
		g.currentScreen = messageScreen
	case "Livre de Sort : Boule de Feu", "Livre de Sort : Soin", "Livre de Sort : Régénération", "Livre de Sort : Poison", "Livre de Sort : Brûlure", "Livre de Sort : Éclair", "Livre de Sort : Barrière Sacrée":
		library.LearnSpellBookSilent(&g.player, item.Name)
		g.player.Inventory[index].Quantity--
		if g.player.Inventory[index].Quantity <= 0 {
			g.player.Inventory = append(g.player.Inventory[:index], g.player.Inventory[index+1:]...)
		}
		g.message = "Sort appris : " + item.Name
		g.returnScreen = inventoryScreen
		g.currentScreen = messageScreen
	case "Potion de poison":
		g.message = "La potion de poison s'utilise pendant un combat."
		g.returnScreen = inventoryScreen
		g.currentScreen = messageScreen
	default:
		if !isEquipmentItem(item.Name) {
			g.message = "Cet objet n'est pas utilisable ici."
		} else {
			library.EquipItem(&g.player, item.Name, index)
			g.message = "Équipement équipé : " + item.Name
		}
		g.returnScreen = inventoryScreen
		g.currentScreen = messageScreen
	}
	if g.selected >= len(g.player.Inventory) {
		g.selected = len(g.player.Inventory) - 1
	}
}

func isEquipmentItem(name string) bool {
	base := strings.TrimSpace(strings.Split(name, "[")[0])
	for _, item := range []string{
		"Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier",
		"Casque de gobelin", "Lame de gobelin", "Tunique de gobelin", "Bottes de gobelin",
		"Carapace de slime", "Bave de slime", "Tunique de slime", "Bottes de slime",
		"Capuche spectrale", "Lame spectrale", "Tunique spectrale", "Bottes spectrales",
		"Casque de golem", "Marteau de golem", "Tunique de golem", "Bottes de golem",
		"Peau de troll renforcée", "Massue de troll", "Tunique de troll", "Bottes de troll",
		"Plumes du canard", "Bec du canard", "Tunique du canard", "Bottes du canard",
		"Heaume squelette", "Épée squelette", "Tunique du squelette", "Bottes du squelette",
		"Fourrure du loup", "Crocs du loup", "Tunique du loup", "Bottes du loup",
		"Chapeau du sorcier", "Bâton maudit", "Tunique du sorcier", "Bottes du sorcier",
		"Écailles de dragon", "Griffe du dragon", "Tunique du dragon", "Bottes du dragon",
		"Dague de l'assassin", "Arc du chasseur", "Arc long", "Arc composite", "Arc elfique", "Marteau du guerrier", "Épée du chevalier",
	} {
		if base == item {
			return true
		}
	}
	return false
}

func (g *game) updateMerchant() {
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.sellMode = !g.sellMode
		g.selected = 0
	}
	if g.sellMode {
		if len(g.player.Inventory) == 0 {
			return
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
			g.selected = (g.selected + 1) % len(g.player.Inventory)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
			g.selected = (g.selected + len(g.player.Inventory) - 1) % len(g.player.Inventory)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.sellInventoryItem(g.selected)
		}
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.selected = (g.selected + 1) % len(shopItems)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.selected = (g.selected + len(shopItems) - 1) % len(shopItems)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		item := shopItems[g.selected]
		if item.name == "Amelioration d'inventaire" && g.player.LimitInventoryUpgrade >= 3 {
			g.message = "Vous avez déjà atteint le maximum de 3 améliorations."
			g.returnScreen = merchantScreen
			g.currentScreen = messageScreen
		} else if !library.BuyItemSilent(&g.player, item.name, item.price) {
			g.message = "Achat impossible : or insuffisant ou inventaire plein."
			g.returnScreen = merchantScreen
			g.currentScreen = messageScreen
		}
	}
}

func (g *game) sellInventoryItem(index int) {
	if index < 0 || index >= len(g.player.Inventory) {
		return
	}
	prices := map[string]int{
		"Potion de vie": 1, "Potion de poison": 3, "Potion de mana": 3,
		"Fourrure de Loup": 2, "Peau de Troll": 3, "Cuir de Sanglier": 1,
		"Plume de Corbeau": 1, "Amelioration d'inventaire": 15,
	}
	price := prices[g.player.Inventory[index].Name]
	if price == 0 {
		price = 1
	}
	g.player.Gold += price
	g.player.Inventory[index].Quantity--
	if g.player.Inventory[index].Quantity <= 0 {
		g.player.Inventory = append(g.player.Inventory[:index], g.player.Inventory[index+1:]...)
	}
	if g.selected >= len(g.player.Inventory) {
		g.selected = len(g.player.Inventory) - 1
	}
}

func (g *game) updateForge() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.selected = (g.selected + 1) % (len(forgeItems) + 1)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.selected = (g.selected + len(forgeItems)) % (len(forgeItems) + 1)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && g.selected < len(forgeItems) {
		library.FabriquerObjet(&g.player, forgeItems[g.selected])
	}
}

func (g *game) updateEnchanter() {
	if len(g.player.Skill) == 0 {
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.selected = (g.selected + 1) % len(g.player.Skill)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.selected = (g.selected + len(g.player.Skill) - 1) % len(g.player.Skill)
	}
	if !inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return
	}
	skill := g.player.Skill[g.selected]
	if skill.Damage <= 0 {
		g.message = "Ce sort ne peut pas être enchanté."
	} else if library.EnchantSkillSilent(&g.player, g.selected) {
		g.message = fmt.Sprintf("%s enchanté : dégâts +5.", skill.Name)
	} else {
		g.message = fmt.Sprintf("Or insuffisant. Coût : %d or.", library.EnchantmentCost(skill))
	}
	g.returnScreen = enchanterScreen
	g.currentScreen = messageScreen
}

func (g *game) startTrainingFight() {
	g.monster = library.InitGoblin("Gobelin d'entrainement", 40, 5)
	g.combatTurn = 1
	g.combatAdventure = false
	g.selected = 0
	g.assassinExtraAction = false
	library.StartCombatBonuses(&g.player)
	g.combatLog = g.passiveCombatMessage()
	if library.AssassinOpeningAttack(&g.player, &g.monster) {
		g.combatLog += " Frappe d'ouverture de l'Assassin !"
	}
	g.currentScreen = combatScreen
}

func (g *game) startAdventure() {
	room := g.player.LastClearedRoom + 1
	g.monster = g.adventureMonster(room)
	g.combatTurn = 1
	g.combatAdventure = true
	g.selected = 0
	g.assassinExtraAction = false
	library.StartCombatBonuses(&g.player)
	g.combatLog = g.passiveCombatMessage()
	if library.AssassinOpeningAttack(&g.player, &g.monster) {
		g.combatLog += " Frappe d'ouverture de l'Assassin !"
	}
	g.currentScreen = combatScreen
}

func (g *game) adventureMonster(room int) library.Monster {
	if room%10 == 0 {
		boss := library.InitGoblinLevel("Boss gobelin", room+2)
		boss.Name = "Boss gobelin"
		boss.MaxHP += 60 + room/10*25
		boss.CurrentHP = boss.MaxHP
		boss.Attack += 8 + room/10*3
		boss.XPReward = g.player.MaxXP - g.player.CurrentXP + 50
		boss.GoldReward = 120 + room/10*50
		return boss
	}
	if room >= 3 && g.random.Intn(100) < 5 {
		monster := library.InitDuck()
		monster.Level = room
		monster.MaxHP += (room - 1) * 8
		monster.CurrentHP = monster.MaxHP
		monster.Attack += room - 1
		monster.Initiative += room - 1
		monster.XPReward += (room - 1) * 15
		monster.GoldReward += (room - 1) * 8
		return monster
	}
	var monster library.Monster
	switch g.random.Intn(9) {
	case 0:
		monster = library.InitGoblinLevel("Gobelin", room)
	case 1:
		monster = library.InitSlimeLevel("Slime", room)
	case 2:
		monster = library.InitGhostLevel("Fantôme", room)
	case 3:
		monster = library.InitGolem()
	case 4:
		monster = library.InitTroll()
	case 5:
		monster = library.InitSkeleton()
	case 6:
		monster = library.InitWolf()
	case 7:
		monster = library.InitWizard()
	default:
		if room >= 5 {
			monster = library.InitDragon()
		} else {
			monster = library.InitGoblinLevel("Gobelin", room)
		}
	}
	monster.Level = room
	if room > 1 {
		monster.MaxHP += (room - 1) * 8
		monster.CurrentHP = monster.MaxHP
		monster.Attack += room - 1
		monster.Initiative += (room - 1)
		monster.XPReward += (room - 1) * 15
		monster.GoldReward += (room - 1) * 8
	}
	return monster
}

func (g *game) updateCombatMenu() {
	const actionCount = 4
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.selected = (g.selected + 1) % actionCount
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.selected = (g.selected + actionCount - 1) % actionCount
	}
	if !inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return
	}
	switch g.selected {
	case 0:
		g.performCombatAttack()
	case 1:
		g.selected = 0
		g.currentScreen = combatSpellsScreen
	case 2:
		g.selected = 0
		g.currentScreen = combatInventoryScreen
	case 3:
		g.currentScreen = mainMenu
	}
}

func (g *game) performCombatAttack() {
	if g.monster.CurrentHP <= 0 || g.player.CurrentHP <= 0 {
		return
	}
	damage := library.BasicAttackDamage(&g.player)
	g.monster.CurrentHP -= damage
	g.combatLog = fmt.Sprintf("Vous infligez %d dégâts à %s.", damage, g.monster.Name)
	if g.monster.CurrentHP <= 0 {
		g.monster.CurrentHP = 0
		if g.combatAdventure {
			reward := library.CollectCombatReward(&g.player, &g.monster)
			g.player.LastClearedRoom++
			g.message = formatReward(reward)
		} else {
			g.message = "VICTOIRE !\nCombat d'entraînement terminé.\nAucune récompense attribuée."
		}
		if g.combatAdventure {
			g.returnScreen = combatScreen
		} else {
			g.returnScreen = mainMenu
		}
		g.currentScreen = messageScreen
		return
	}
	if g.canTakeAssassinExtraAction() {
		return
	}
	g.assassinExtraAction = false
	g.enemyTurn()
	if g.player.CurrentHP == 0 {
		g.player.CurrentHP = g.player.MaxHP / 2
		g.message = "Défaite. Vous réapparaissez avec la moitié de vos PV."
		if g.combatAdventure {
			g.returnScreen = combatScreen
		} else {
			g.returnScreen = mainMenu
		}
		g.currentScreen = messageScreen
	}
}

func (g *game) canTakeAssassinExtraAction() bool {
	if g.player.Class != "Assassin" || g.assassinExtraAction || g.monster.CurrentHP <= 0 || g.combatTurn%2 == 0 {
		return false
	}
	g.assassinExtraAction = true
	g.combatLog += " Action supplémentaire de l'Assassin."
	return true
}

func (g *game) enemyTurn() {
	poisonDamage := library.ApplyPoisonSilent(&g.monster)
	if g.monster.CurrentHP <= 0 {
		g.combatLog += fmt.Sprintf(" Poison : %d dégâts.", poisonDamage)
		return
	}
	damage := library.MonsterPatternDamage(&g.monster, g.combatTurn)
	enemyLog := ""
	if library.TryDodge(&g.player) {
		enemyLog = "Esquive réussie : l'attaque ennemie est évitée."
	} else {
		blockedDamage := library.BlockDamage(&g.player, damage)
		g.player.CurrentHP -= blockedDamage
		enemyLog = fmt.Sprintf("Le %s inflige %d dégâts.", g.monster.Name, blockedDamage)
		if blockedDamage < damage {
			enemyLog += " Votre protection réduit les dégâts."
		}
	}
	g.combatLog += " " + enemyLog
	if library.TryCounterAttack(&g.player, &g.monster) {
		g.combatLog += " Contre-attaque du Samouraï !"
	}
	if g.player.CurrentHP < 0 {
		g.player.CurrentHP = 0
	}
	g.combatTurn++
}

func (g *game) updateCombatSpells() {
	if len(g.player.Skill) == 0 {
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.selected = (g.selected + 1) % len(g.player.Skill)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.selected = (g.selected + len(g.player.Skill) - 1) % len(g.player.Skill)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if g.useCombatSpell(g.selected) {
			g.afterCombatAction()
		}
	}
}

func (g *game) useCombatSpell(index int) bool {
	if index < 0 || index >= len(g.player.Skill) {
		return false
	}
	skill := g.player.Skill[index]
	if g.player.Mana < skill.ManaCost {
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
		g.combatLog = fmt.Sprintf("Soin : vous récupérez %d PV.", heal)
	case "Régénération":
		g.player.Effects = append(g.player.Effects, library.Effect{Name: skill.Name, Value: skill.HealAmount, TurnsLeft: skill.EffectTurns})
		g.combatLog = "Régénération active pour 3 tours."
	case "Poison", "Brûlure":
		damage := library.SpellDamage(&g.player, skill.Damage)
		g.monster.Effects = append(g.monster.Effects, library.Effect{Name: skill.Name, Value: damage, TurnsLeft: skill.EffectTurns})
		g.combatLog = fmt.Sprintf("%s appliqué : %d dégâts pendant %d tours.", skill.Name, damage, skill.EffectTurns)
	case "Barrière Sacrée":
		g.player.HolyBarrier = true
		g.combatLog = "Barrière sacrée activée."
	case "Invocation de soldat":
		if g.player.Summon != nil && g.player.Summon.CurrentHP > 0 {
			g.player.Mana += skill.ManaCost
			return false
		}
		summonHP, summonAttack := library.SummonStats(&g.player)
		g.player.Summon = &library.Monster{Name: "Soldat invoqué", Pattern: "summon", Level: g.player.Level, MaxHP: summonHP, CurrentHP: summonHP, Attack: summonAttack, Initiative: g.player.Initiative}
		g.combatLog = "Soldat invoqué."
	case "Éclair":
		damage := library.SpellDamage(&g.player, skill.Damage)
		g.monster.CurrentHP -= damage
		if g.monster.CurrentHP < 0 {
			g.monster.CurrentHP = 0
		}
		g.monster.Initiative -= 15
		if g.monster.Initiative < 0 {
			g.monster.Initiative = 0
		}
		g.combatLog = fmt.Sprintf("Éclair : %d dégâts, initiative ennemie réduite.", damage)
	default:
		damage := library.SpellDamage(&g.player, skill.Damage)
		g.monster.CurrentHP -= damage
		if g.monster.CurrentHP < 0 {
			g.monster.CurrentHP = 0
		}
		g.combatLog = fmt.Sprintf("%s inflige %d dégâts à %s.", skill.Name, damage, g.monster.Name)
	}
	return true
}

func (g *game) updateCombatInventory() {
	if len(g.player.Inventory) == 0 {
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.selected = (g.selected + 1) % len(g.player.Inventory)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.selected = (g.selected + len(g.player.Inventory) - 1) % len(g.player.Inventory)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		item := g.player.Inventory[g.selected]
		used := false
		if item.Name == "Potion de vie" {
			used = library.TakePotSilent(&g.player, g.selected)
			if used {
				g.combatLog = "Potion de vie utilisée."
			}
		}
		if item.Name == "Potion de mana" {
			used = library.TakeManaPotSilent(&g.player, g.selected)
			if used {
				g.combatLog = "Potion de mana utilisée."
			}
		}
		if item.Name == "Potion de poison" {
			used = library.PoisonPotSilent(&g.player, &g.monster, g.selected)
			if used {
				g.combatLog = "Potion de poison utilisée : le monstre subira 10 dégâts pendant 3 tours."
			}
		}
		if used {
			g.afterCombatAction()
		}
	}
}

func (g *game) afterCombatAction() {
	if g.monster.CurrentHP <= 0 {
		if g.combatAdventure {
			reward := library.CollectCombatReward(&g.player, &g.monster)
			g.player.LastClearedRoom++
			g.message = formatReward(reward)
		} else {
			g.message = "VICTOIRE !\nCombat d'entraînement terminé.\nAucune récompense attribuée."
		}
		if g.combatAdventure {
			g.returnScreen = combatScreen
		} else {
			g.returnScreen = mainMenu
		}
		g.currentScreen = messageScreen
		return
	}
	if g.canTakeAssassinExtraAction() {
		return
	}
	g.assassinExtraAction = false
	playerEffectLog := g.applyPlayerEffects()
	monsterEffectLog := g.applyMonsterEffects()
	if playerEffectLog != "" || monsterEffectLog != "" {
		g.combatLog += " " + playerEffectLog + " " + monsterEffectLog
	}
	if g.monster.CurrentHP <= 0 {
		g.afterCombatAction()
		return
	}
	if g.player.Summon != nil && g.player.Summon.CurrentHP > 0 {
		damage := g.player.Summon.Attack
		g.monster.CurrentHP -= damage
		if g.monster.CurrentHP < 0 {
			g.monster.CurrentHP = 0
		}
		g.combatLog += fmt.Sprintf(" Soldat invoqué : %d dégâts.", damage)
		if g.monster.CurrentHP <= 0 {
			g.afterCombatAction()
			return
		}
	}
	poisonDamage := library.ApplyPoisonSilent(&g.monster)
	if g.monster.CurrentHP <= 0 {
		g.combatLog += fmt.Sprintf(" Poison : %d dégâts.", poisonDamage)
		g.afterCombatAction()
		return
	}
	damage := library.MonsterPatternDamage(&g.monster, g.combatTurn)
	if g.player.HolyBarrier {
		g.player.HolyBarrier = false
		g.combatLog += " Barrière sacrée : l'attaque ennemie est bloquée."
	} else {
		if library.TryDodge(&g.player) {
			g.combatLog += " Esquive réussie : l'attaque ennemie est évitée."
		} else {
			blockedDamage := library.BlockDamage(&g.player, damage)
			g.player.CurrentHP -= blockedDamage
			g.combatLog += fmt.Sprintf(" Le %s inflige %d dégâts.", g.monster.Name, blockedDamage)
			if blockedDamage < damage {
				g.combatLog += " Votre protection réduit les dégâts."
			}
		}
	}
	if library.TryCounterAttack(&g.player, &g.monster) {
		g.combatLog += " Contre-attaque du Samouraï !"
	}
	if g.player.CurrentHP <= 0 {
		g.player.CurrentHP = g.player.MaxHP / 2
		g.message = "Défaite. Vous réapparaissez avec la moitié de vos PV."
		g.returnScreen = mainMenu
		g.currentScreen = messageScreen
		return
	}
	g.combatTurn++
	if g.player.Mana < g.player.MaxMana {
		g.player.Mana += 5
		if g.player.Mana > g.player.MaxMana {
			g.player.Mana = g.player.MaxMana
		}
	}
}

func (g *game) applyPlayerEffects() string {
	parts := []string{}
	for index := 0; index < len(g.player.Effects); {
		effect := &g.player.Effects[index]
		if effect.Name == "Régénération" {
			g.player.CurrentHP += effect.Value
			if g.player.CurrentHP > g.player.MaxHP {
				g.player.CurrentHP = g.player.MaxHP
			}
			parts = append(parts, fmt.Sprintf("Régénération : +%d PV.", effect.Value))
		}
		effect.TurnsLeft--
		if effect.TurnsLeft <= 0 {
			g.player.Effects = append(g.player.Effects[:index], g.player.Effects[index+1:]...)
			continue
		}
		index++
	}
	return strings.Join(parts, " ")
}

func (g *game) applyMonsterEffects() string {
	parts := []string{}
	for index := 0; index < len(g.monster.Effects); {
		effect := &g.monster.Effects[index]
		g.monster.CurrentHP -= effect.Value
		if g.monster.CurrentHP < 0 {
			g.monster.CurrentHP = 0
		}
		parts = append(parts, fmt.Sprintf("%s : %d dégâts.", effect.Name, effect.Value))
		effect.TurnsLeft--
		if effect.TurnsLeft <= 0 {
			g.monster.Effects = append(g.monster.Effects[:index], g.monster.Effects[index+1:]...)
			continue
		}
		index++
	}
	return strings.Join(parts, " ")
}

func (g *game) passiveCombatMessage() string {
	switch g.player.Class {
	case "Guerrier":
		if g.player.CombatWeaponBonus > 0 {
			return "Passif Guerrier : bonus d'attaque d'arme actif."
		}
		if g.player.CombatSpellBonus > 0 {
			return "Passif Guerrier : bonus de sort actif."
		}
		return "Passif Guerrier : bonus d'initiative actif."
	case "Mage":
		return "Passif Mage : les dégâts des sorts sont augmentés."
	case "Archer":
		return "Passif Archer : chance d'esquive active."
	case "Assassin":
		return "Passif Assassin : frappe d'ouverture renforcée et esquive active."
	case "Chevalier":
		return "Passif Chevalier : blocage amélioré actif."
	case "Samourai":
		return "Passif Samouraï : esquive et contre-attaque actives."
	case "Clerc":
		return "Passif Clerc : soins augmentés."
	case "Barbare":
		return "Passif Barbare : dégâts augmentés selon les PV perdus."
	case "Invocateur":
		return "Passif Invocateur : invocation évolutive disponible."
	default:
		return "Combat commencé."
	}
}

func formatReward(reward library.CombatReward) string {
	lines := []string{"VICTOIRE !", fmt.Sprintf("+%d XP    +%d or", reward.XP, reward.Gold)}
	if reward.NewLevels > 0 {
		lines = append(lines, fmt.Sprintf("Niveau gagné : +%d", reward.NewLevels))
	}
	if len(reward.Drops) == 0 {
		lines = append(lines, "Aucun équipement trouvé.")
	} else {
		lines = append(lines, "Butin :")
		lines = append(lines, reward.Drops...)
	}
	return strings.Join(lines, "\n")
}

func (g *game) Draw(screenImage *ebiten.Image) {
	screenImage.Fill(theme.background)
	drawBackdrop(screenImage)
	drawHeader(screenImage)
	switch g.currentScreen {
	case startSelection:
		g.drawStartSelection(screenImage)
	case nameCreation:
		g.drawNameCreation(screenImage)
	case classCreation:
		g.drawClassCreation(screenImage)
	case predefinedSelection:
		g.drawPredefinedSelection(screenImage)
	case mainMenu:
		g.drawMainMenu(screenImage)
	case characterSheet:
		g.drawCharacterSheet(screenImage)
	case inventoryScreen:
		g.drawInventory(screenImage)
	case merchantScreen:
		g.drawMerchant(screenImage)
	case forgeScreen:
		g.drawForge(screenImage)
	case enchanterScreen:
		g.drawEnchanter(screenImage)
	case creditsScreen:
		g.drawCredits(screenImage)
	case combatScreen:
		g.drawCombat(screenImage)
	case combatSpellsScreen:
		g.drawCombatSpells(screenImage)
	case combatInventoryScreen:
		g.drawCombatInventory(screenImage)
	case messageScreen:
		drawMessage(screenImage, g.message)
	}
}

func drawHeader(screenImage *ebiten.Image) {
	ebitenutil.DrawRect(screenImage, 0, 0, screenWidth, 86, theme.panel)
	ebitenutil.DrawRect(screenImage, 0, 83, screenWidth, 2, theme.gold)
	ebitenutil.DebugPrintAt(screenImage, "L O O T S T O R M", 38, 25)
	ebitenutil.DebugPrintAt(screenImage, "DUNGEON RUN  |  Echap: retour", 670, 34)
}

func drawPanel(screenImage *ebiten.Image, x, y, width, height float64, selected bool) {
	fill := theme.panel
	if selected {
		fill = theme.panelLight
	}
	ebitenutil.DrawRect(screenImage, x, y, width, height, fill)
	line := theme.gold
	if selected {
		line = theme.goldBright
	}
	ebitenutil.DrawRect(screenImage, x, y, width, 1, line)
	ebitenutil.DrawRect(screenImage, x, y+height-1, width, 1, line)
	ebitenutil.DrawRect(screenImage, x, y, 1, height, line)
	ebitenutil.DrawRect(screenImage, x+width-1, y, 1, height, line)
}

func drawChoiceRow(screenImage *ebiten.Image, x, y, width, height float64, selected bool) {
	if selected {
		ebitenutil.DrawRect(screenImage, x, y+height/2-1, 3, 2, theme.goldBright)
		ebitenutil.DrawRect(screenImage, x+12, y+height-3, width-24, 1, theme.gold)
	}
}

func drawBackdrop(screenImage *ebiten.Image) {
	if currentGame != nil {
		background := currentGame.backgroundForScreen()
		if background != nil {
			drawBackgroundImage(screenImage, background)
			ebitenutil.DrawRect(screenImage, 0, 86, screenWidth, screenHeight-86, color.RGBA{R: 0, G: 0, B: 0, A: 86})
			return
		}
	}
	if currentGame != nil && isCombatScreen(currentGame.currentScreen) && currentGame.combatBackground != nil {
		drawBackgroundImage(screenImage, currentGame.combatBackground)
		ebitenutil.DrawRect(screenImage, 0, 86, screenWidth, screenHeight-86, color.RGBA{R: 0, G: 0, B: 0, A: 82})
		return
	}
	if currentGame != nil && currentGame.background != nil {
		drawBackgroundImage(screenImage, currentGame.background)
		ebitenutil.DrawRect(screenImage, 0, 86, screenWidth, screenHeight-86, color.RGBA{R: 0, G: 0, B: 0, A: 92})
		return
	}
	ebitenutil.DrawRect(screenImage, 0, 86, screenWidth, screenHeight-86, theme.background)
	ebitenutil.DrawRect(screenImage, 0, 86, screenWidth, 8, color.RGBA{R: 45, G: 32, B: 34, A: 255})
	ebitenutil.DrawRect(screenImage, 300, 86, 360, 5, theme.gold)
	ebitenutil.DrawRect(screenImage, 320, 91, 320, 2, color.RGBA{R: 78, G: 55, B: 48, A: 255})
	for y := 112; y < 540; y += 72 {
		ebitenutil.DrawRect(screenImage, 0, float64(y), screenWidth, 1, color.RGBA{R: 35, G: 29, B: 35, A: 255})
		for x := 32; x < screenWidth; x += 128 {
			ebitenutil.DrawRect(screenImage, float64(x), float64(y), 1, 72, color.RGBA{R: 30, G: 25, B: 31, A: 255})
		}
	}
	drawDungeonArch(screenImage)
	for _, x := range []float64{18, 918} {
		ebitenutil.DrawRect(screenImage, x, 150, 24, 310, color.RGBA{R: 30, G: 24, B: 29, A: 255})
		ebitenutil.DrawRect(screenImage, x+5, 150, 14, 310, color.RGBA{R: 54, G: 42, B: 42, A: 255})
		ebitenutil.DrawRect(screenImage, x-5, 150, 34, 8, theme.gold)
		ebitenutil.DrawRect(screenImage, x-5, 452, 34, 8, theme.gold)
		ebitenutil.DrawRect(screenImage, x+8, 132, 8, 18, color.RGBA{R: 213, G: 91, B: 43, A: 255})
	}
	ebitenutil.DrawRect(screenImage, 0, 540, screenWidth, 60, color.RGBA{R: 10, G: 9, B: 13, A: 255})
	for x := 0; x < screenWidth; x += 48 {
		ebitenutil.DrawRect(screenImage, float64(x), 540, 1, 60, color.RGBA{R: 33, G: 27, B: 31, A: 255})
	}
}

func (g *game) backgroundForScreen() *ebiten.Image {
	switch g.currentScreen {
	case merchantScreen:
		if g.merchantBackground != nil {
			return g.merchantBackground
		}
	case forgeScreen:
		if g.forgeBackground != nil {
			return g.forgeBackground
		}
	case enchanterScreen:
		if g.enchanterBackground != nil {
			return g.enchanterBackground
		}
	case combatScreen, combatSpellsScreen, combatInventoryScreen:
		if g.combatBackground != nil {
			return g.combatBackground
		}
	}
	return g.background
}

func isCombatScreen(current screen) bool {
	return current == combatScreen || current == combatSpellsScreen || current == combatInventoryScreen
}

func drawBackgroundImage(screenImage, background *ebiten.Image) {
	width := float64(background.Bounds().Dx())
	height := float64(background.Bounds().Dy())
	if width == 0 || height == 0 {
		return
	}
	scaleX := float64(screenWidth) / width
	scaleY := float64(screenHeight-86) / height
	scale := scaleX
	if scaleY > scale {
		scale = scaleY
	}
	drawWidth := width * scale
	drawHeight := height * scale
	options := &ebiten.DrawImageOptions{}
	options.Filter = ebiten.FilterLinear
	options.GeoM.Scale(scale, scale)
	options.GeoM.Translate((float64(screenWidth)-drawWidth)/2, 86+(float64(screenHeight-86)-drawHeight)/2)
	screenImage.DrawImage(background, options)
}

func drawDungeonArch(screenImage *ebiten.Image) {
	stone := color.RGBA{R: 43, G: 35, B: 40, A: 255}
	shadow := color.RGBA{R: 22, G: 18, B: 24, A: 255}
	ebitenutil.DrawRect(screenImage, 330, 112, 300, 10, stone)
	ebitenutil.DrawRect(screenImage, 350, 122, 260, 8, stone)
	ebitenutil.DrawRect(screenImage, 370, 130, 220, 5, shadow)
	ebitenutil.DrawRect(screenImage, 330, 112, 10, 95, stone)
	ebitenutil.DrawRect(screenImage, 620, 112, 10, 95, stone)
	ebitenutil.DrawRect(screenImage, 350, 122, 5, 78, color.RGBA{R: 69, G: 52, B: 52, A: 255})
	ebitenutil.DrawRect(screenImage, 605, 122, 5, 78, color.RGBA{R: 69, G: 52, B: 52, A: 255})
	for _, point := range [][2]float64{{250, 270}, {285, 310}, {690, 280}, {735, 330}} {
		ebitenutil.DrawRect(screenImage, point[0], point[1], 16, 2, color.RGBA{R: 65, G: 48, B: 50, A: 255})
		ebitenutil.DrawRect(screenImage, point[0]+7, point[1]-8, 2, 18, color.RGBA{R: 65, G: 48, B: 50, A: 255})
	}
}

func drawHeroSprite(screenImage *ebiten.Image, x, y int, name, class string) {
	if currentGame != nil {
		if sprite := currentGame.sprites["hero:"+name]; sprite != nil {
			drawImageSprite(screenImage, sprite, x, y, 160)
			return
		}
		if sprite := currentGame.sprites["class:"+class]; sprite != nil {
			drawImageSprite(screenImage, sprite, x, y, 160)
			return
		}
	}
	skin := color.RGBA{R: 221, G: 166, B: 119, A: 255}
	coat := color.RGBA{R: 48, G: 130, B: 128, A: 255}
	if class == "Mage" || class == "Clerc" {
		coat = color.RGBA{R: 92, G: 76, B: 157, A: 255}
	}
	if class == "Barbare" || class == "Guerrier" {
		coat = color.RGBA{R: 164, G: 69, B: 62, A: 255}
	}
	ebitenutil.DrawRect(screenImage, float64(x+18), float64(y), 28, 28, skin)
	ebitenutil.DrawRect(screenImage, float64(x+12), float64(y+28), 40, 48, coat)
	ebitenutil.DrawRect(screenImage, float64(x+10), float64(y+76), 15, 28, color.RGBA{R: 53, G: 64, B: 80, A: 255})
	ebitenutil.DrawRect(screenImage, float64(x+39), float64(y+76), 15, 28, color.RGBA{R: 53, G: 64, B: 80, A: 255})
	ebitenutil.DrawRect(screenImage, float64(x+4), float64(y+36), 8, 32, skin)
	ebitenutil.DrawRect(screenImage, float64(x+52), float64(y+36), 8, 32, skin)
	ebitenutil.DrawRect(screenImage, float64(x+25), float64(y+10), 5, 5, color.RGBA{R: 15, G: 20, B: 28, A: 255})
	if class == "Mage" {
		ebitenutil.DrawRect(screenImage, float64(x+8), float64(y-8), 48, 8, color.RGBA{R: 62, G: 42, B: 92, A: 255})
	}
}

func drawMonsterSprite(screenImage *ebiten.Image, x, y int, pattern string) {
	if currentGame != nil {
		if sprite := currentGame.sprites[pattern]; sprite != nil {
			drawImageSprite(screenImage, sprite, x, y, 160)
			return
		}
	}
	body := color.RGBA{R: 82, G: 151, B: 92, A: 255}
	switch pattern {
	case "slime":
		body = color.RGBA{R: 71, G: 177, B: 169, A: 255}
	case "ghost":
		body = color.RGBA{R: 150, G: 160, B: 190, A: 255}
	case "dragon":
		body = color.RGBA{R: 171, G: 70, B: 61, A: 255}
	case "golem":
		body = color.RGBA{R: 112, G: 111, B: 116, A: 255}
	case "troll":
		body = color.RGBA{R: 125, G: 91, B: 62, A: 255}
	}
	ebitenutil.DrawRect(screenImage, float64(x+12), float64(y+22), 70, 62, body)
	ebitenutil.DrawRect(screenImage, float64(x+22), float64(y), 50, 35, body)
	ebitenutil.DrawRect(screenImage, float64(x+25), float64(y+12), 7, 7, color.RGBA{R: 235, G: 220, B: 102, A: 255})
	ebitenutil.DrawRect(screenImage, float64(x+60), float64(y+12), 7, 7, color.RGBA{R: 235, G: 220, B: 102, A: 255})
	ebitenutil.DrawRect(screenImage, float64(x+4), float64(y+38), 12, 34, body)
	ebitenutil.DrawRect(screenImage, float64(x+78), float64(y+38), 12, 34, body)
}

var currentGame *game

func drawImageSprite(screenImage, sprite *ebiten.Image, x, y, size int) {
	width := float64(sprite.Bounds().Dx())
	height := float64(sprite.Bounds().Dy())
	if width == 0 || height == 0 {
		return
	}
	scale := float64(size) / width
	if height*scale > float64(size) {
		scale = float64(size) / height
	}
	options := &ebiten.DrawImageOptions{}
	options.Filter = ebiten.FilterLinear
	options.GeoM.Scale(scale, scale)
	options.GeoM.Translate(float64(x)+(float64(size)-width*scale)/2, float64(y)+(float64(size)-height*scale)/2)
	screenImage.DrawImage(sprite, options)
}

func (g *game) drawNameCreation(screenImage *ebiten.Image) {
	ebitenutil.DebugPrintAt(screenImage, "CREATION DE VOTRE HEROS", 52, 140)
	ebitenutil.DebugPrintAt(screenImage, "Entrez votre nom puis appuyez sur Entree", 52, 180)
	drawPanel(screenImage, 52, 220, 620, 48, false)
	ebitenutil.DebugPrintAt(screenImage, g.name+"_", 70, 238)
}

func (g *game) drawStartSelection(screenImage *ebiten.Image) {
	ebitenutil.DebugPrintAt(screenImage, "BIENVENUE DANS LOOTSTORM", 52, 150)
	ebitenutil.DebugPrintAt(screenImage, "Choisissez votre façon de commencer", 52, 185)
	items := []string{"Créer mon propre personnage", "Choisir un héros prédéfini"}
	for index, item := range items {
		y := 240 + index*52
		drawChoiceRow(screenImage, 52, float64(y), 500, 38, index == g.selected)
		prefix := "  "
		if index == g.selected {
			prefix = "> "
		}
		ebitenutil.DebugPrintAt(screenImage, prefix+item, 70, y+12)
	}
}

func (g *game) drawPredefinedSelection(screenImage *ebiten.Image) {
	ebitenutil.DebugPrintAt(screenImage, "HEROS PREDEFINIS", 52, 124)
	ebitenutil.DebugPrintAt(screenImage, "Fleches haut/bas   Entree: confirmer", 52, 152)
	for index, hero := range predefinedHeroes {
		y := 195 + index*44
		drawChoiceRow(screenImage, 52, float64(y), 540, 36, index == g.selected)
		prefix := "  "
		if index == g.selected {
			prefix = "> "
		}
		ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("%s%-22s %-12s %d PV", prefix, hero.name, hero.class, hero.hp), 68, y+11)
	}
	selectedHero := predefinedHeroes[g.selected]
	drawPanel(screenImage, 650, 190, 250, 300, true)
	ebitenutil.DebugPrintAt(screenImage, "APERÇU", 680, 215)
	drawHeroSprite(screenImage, 710, 245, selectedHero.name, selectedHero.class)
	ebitenutil.DebugPrintAt(screenImage, selectedHero.name, 680, 405)
	ebitenutil.DebugPrintAt(screenImage, "Classe: "+selectedHero.class, 680, 430)
	ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("PV max: %d", selectedHero.hp), 680, 455)
}

func (g *game) drawClassCreation(screenImage *ebiten.Image) {
	ebitenutil.DebugPrintAt(screenImage, "CHOISISSEZ UNE CLASSE POUR "+strings.ToUpper(g.name), 52, 124)
	ebitenutil.DebugPrintAt(screenImage, "Fleches haut/bas   Entree: confirmer", 52, 152)
	for index, choice := range classes {
		y := 190 + index*39
		selected := index == g.selected
		drawChoiceRow(screenImage, 52, float64(y), 430, 32, selected)
		prefix := "  "
		if selected {
			prefix = "> "
		}
		ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("%s%-14s %3d PV", prefix, choice.name, choice.hp), 68, y+10)
	}
	selectedClass := classes[g.selected]
	drawPanel(screenImage, 560, 185, 340, 315, true)
	ebitenutil.DebugPrintAt(screenImage, "APERÇU DE LA CLASSE", 590, 212)
	drawHeroSprite(screenImage, 665, 240, "", selectedClass.name)
	ebitenutil.DebugPrintAt(screenImage, selectedClass.name, 590, 405)
	ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("PV max: %d", selectedClass.hp), 590, 432)
	for index, line := range wrapText(passiveDescription(selectedClass.name), 34) {
		ebitenutil.DebugPrintAt(screenImage, line, 590, 465+index*18)
	}
}

func (g *game) drawMainMenu(screenImage *ebiten.Image) {
	ebitenutil.DebugPrintAt(screenImage, "CAMP DE BASE", 52, 124)
	ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("%s  |  %s", g.player.Name, g.player.Class), 52, 152)
	items := []string{"Mes statistiques", "Aventure", "Combat d'entrainement", "Inventaire", "Marchand", "Forgeron", "Enchanteur", "Qui sont-ils ?", "Quitter"}
	for index, item := range items {
		y := 184 + index*39
		drawChoiceRow(screenImage, 52, float64(y), 430, 36, index == g.selected)
		prefix := "  "
		if index == g.selected {
			prefix = "> "
		}
		ebitenutil.DebugPrintAt(screenImage, prefix+item, 68, y+11)
	}
	drawPanel(screenImage, 570, 198, 330, 220, false)
	ebitenutil.DebugPrintAt(screenImage, "ETAT DU HEROS", 596, 225)
	drawBar(screenImage, "PV", g.player.CurrentHP, g.player.MaxHP, 596, 265, color.RGBA{R: 217, G: 92, B: 85, A: 255})
	drawBar(screenImage, "MANA", g.player.Mana, g.player.MaxMana, 596, 310, color.RGBA{R: 83, G: 151, B: 218, A: 255})
	drawExperienceBar(screenImage, g.player, 596, 355)
	ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("Niveau %d     Or %d", g.player.Level, g.player.Gold), 596, 395)
}

func drawBar(screenImage *ebiten.Image, label string, value, maximum int, x, y int, fill color.Color) {
	ratio := 0.0
	if maximum > 0 {
		ratio = float64(value) / float64(maximum)
	}
	ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("%s  %d/%d", label, value, maximum), x, y-18)
	ebitenutil.DrawRect(screenImage, float64(x), float64(y), 220, 12, color.RGBA{R: 45, G: 52, B: 66, A: 255})
	ebitenutil.DrawRect(screenImage, float64(x), float64(y), 220*ratio, 12, fill)
}

func drawExperienceBar(screenImage *ebiten.Image, player library.Character, x, y int) {
	ratio := 0.0
	if player.MaxXP > 0 {
		ratio = float64(player.CurrentXP) / float64(player.MaxXP)
	}
	ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("XP  %d/%d", player.CurrentXP, player.MaxXP), x, y-18)
	ebitenutil.DrawRect(screenImage, float64(x), float64(y), 220, 12, color.RGBA{R: 45, G: 52, B: 66, A: 255})
	ebitenutil.DrawRect(screenImage, float64(x), float64(y), 220*ratio, 12, color.RGBA{R: 224, G: 174, B: 67, A: 255})
}

func (g *game) drawCharacterSheet(screenImage *ebiten.Image) {
	ebitenutil.DebugPrintAt(screenImage, "FICHE DU PERSONNAGE", 52, 124)
	ebitenutil.DebugPrintAt(screenImage, "Entree: retour", 52, 152)
	drawPanel(screenImage, 40, 185, 450, 375, false)
	drawPanel(screenImage, 510, 185, 410, 375, false)
	bonusAttack := g.player.Equip.WeaponDamage
	bonusHP := g.player.Equip.HelmetHP + g.player.Equip.ChestplateHP + g.player.Equip.BootsHP
	lines := []string{"Nom: " + g.player.Name, "Classe: " + g.player.Class, fmt.Sprintf("Niveau: %d", g.player.Level), fmt.Sprintf("Attaque: %d (+%d eq.) = %d", g.player.Attack, bonusAttack, g.player.Attack+bonusAttack), fmt.Sprintf("PV: %d / %d (+%d eq.)", g.player.CurrentHP, g.player.MaxHP, bonusHP), fmt.Sprintf("Mana: %d / %d", g.player.Mana, g.player.MaxMana), fmt.Sprintf("Initiative: %d", g.player.Initiative), fmt.Sprintf("XP: %d / %d", g.player.CurrentXP, g.player.MaxXP), fmt.Sprintf("Or: %d", g.player.Gold), fmt.Sprintf("Inventaire: %d / %d", len(g.player.Inventory), g.player.LimitInventory)}
	for index, line := range lines {
		for lineIndex, wrappedLine := range wrapText(line, 38) {
			ebitenutil.DebugPrintAt(screenImage, wrappedLine, 70, 212+index*28+lineIndex*16)
		}
	}
	drawExperienceBar(screenImage, g.player, 70, 520)
	ebitenutil.DebugPrintAt(screenImage, "EQUIPEMENT", 540, 212)
	ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("Casque: %s (+%d PV)", g.player.Equip.Helmet, g.player.Equip.HelmetHP), 540, 250)
	ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("Plastron: %s (+%d PV)", g.player.Equip.Chestplate, g.player.Equip.ChestplateHP), 540, 282)
	ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("Bottes: %s (+%d PV)", g.player.Equip.Boots, g.player.Equip.BootsHP), 540, 314)
	ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("Arme: %s (+%d ATQ)", g.player.Equip.Weapon, g.player.Equip.WeaponDamage), 540, 346)
	ebitenutil.DebugPrintAt(screenImage, "PASSIF", 540, 394)
	for index, line := range wrapText(passiveDescription(g.player.Class), 36) {
		ebitenutil.DebugPrintAt(screenImage, line, 540, 426+index*18)
	}
}

func passiveDescription(class string) string {
	passives := map[string]string{
		"Guerrier":   "Atout aleatoire : arme, sort ou initiative en combat.",
		"Mage":       "Sorts infligeant 1,5 fois leurs degats et progression de mana augmentee.",
		"Archer":     "20% esquive et degats doubles avec l'Arc du chasseur.",
		"Assassin":   "Ouverture x1,5, 22% esquive, saignement avec la Dague et action bonus un tour sur deux.",
		"Chevalier":  "35% de chance de reduire les degats de moitie.",
		"Samourai":   "15% esquive et 25% de chance de contre-attaque.",
		"Clerc":      "Soins augmentes de 50%.",
		"Barbare":    "Degats augmentes selon les PV manquants.",
		"Invocateur": "Invocation d'un soldat qui evolue avec le niveau.",
	}
	return passives[class]
}

func (g *game) drawInventory(screenImage *ebiten.Image) {
	ebitenutil.DebugPrintAt(screenImage, "INVENTAIRE", 52, 124)
	ebitenutil.DebugPrintAt(screenImage, "Entree: utiliser/equiper   Fleches: naviguer", 52, 152)
	if len(g.player.Inventory) == 0 {
		ebitenutil.DebugPrintAt(screenImage, "Votre inventaire est vide.", 70, 215)
		return
	}
	for index, item := range g.player.Inventory {
		y := 200 + index*34
		if y > 550 {
			break
		}
		if index == g.selected {
			drawChoiceRow(screenImage, 52, float64(y-5), 540, 30, true)
		}
		prefix := " "
		if index == g.selected {
			prefix = ">"
		}
		ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("%s %-30s %-14s x%d", prefix, library.ItemDisplayName(item), library.ItemStatSummary(item.Name), item.Quantity), 68, y+5)
	}
}

func (g *game) drawMerchant(screenImage *ebiten.Image) {
	if g.sellMode {
		ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("VENTE   Or: %d", g.player.Gold), 52, 124)
		ebitenutil.DebugPrintAt(screenImage, "Entree: vendre   Fleches: naviguer   S: achats   Echap: retour", 52, 152)
		for index, item := range g.player.Inventory {
			y := 190 + index*30
			if index == g.selected {
				drawChoiceRow(screenImage, 52, float64(y-3), 700, 27, true)
			}
			prefix := " "
			if index == g.selected {
				prefix = ">"
			}
			ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("%s %-42s x%d", prefix, library.ItemDisplayName(item), item.Quantity), 68, y+5)
		}
		return
	}
	ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("MARCHAND   Or: %d", g.player.Gold), 52, 124)
	ebitenutil.DebugPrintAt(screenImage, "Entree: acheter   Fleches: naviguer   S: ventes   Echap: retour", 52, 152)
	for index, item := range shopItems {
		y := 190 + index*25
		if index == g.selected {
			drawChoiceRow(screenImage, 52, float64(y-3), 700, 23, true)
		}
		prefix := " "
		if index == g.selected {
			prefix = ">"
		}
		ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("%s %-42s %d or", prefix, item.name, item.price), 68, y+4)
	}
}

func (g *game) drawForge(screenImage *ebiten.Image) {
	ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("FORGERON   Or: %d", g.player.Gold), 52, 124)
	ebitenutil.DebugPrintAt(screenImage, "Entree: fabriquer   Fleches: naviguer   Echap: retour", 52, 152)
	for index, item := range forgeItems {
		y := 220 + index*48
		if index == g.selected {
			drawChoiceRow(screenImage, 52, float64(y-5), 580, 34, true)
		}
		prefix := " "
		if index == g.selected {
			prefix = ">"
		}
		ebitenutil.DebugPrintAt(screenImage, prefix+" "+item+" (5 or)", 68, y+8)
	}
	materials := []string{"Chapeau: 1 Plume de Corbeau + 1 Cuir de Sanglier", "Tunique: 2 Fourrures de Loup + 1 Peau de Troll", "Bottes: 1 Fourrure de Loup + 1 Cuir de Sanglier"}
	ebitenutil.DebugPrintAt(screenImage, "MATERIAUX REQUIS", 600, 220)
	for index, material := range materials {
		ebitenutil.DebugPrintAt(screenImage, material, 600, 250+index*34)
	}
}

func (g *game) drawEnchanter(screenImage *ebiten.Image) {
	ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("ENCHANTEUR   Or: %d", g.player.Gold), 52, 124)
	ebitenutil.DebugPrintAt(screenImage, "Entree: enchanter   Fleches: naviguer   Echap: retour", 52, 152)
	if len(g.player.Skill) == 0 {
		ebitenutil.DebugPrintAt(screenImage, "Vous ne connaissez aucun sort.", 70, 220)
		return
	}
	for index, skill := range g.player.Skill {
		y := 205 + index*38
		if index == g.selected {
			drawChoiceRow(screenImage, 52, float64(y-5), 650, 30, true)
		}
		cost := "Non enchantable"
		if skill.Damage > 0 {
			cost = fmt.Sprintf("+5 dégâts | %d or", library.EnchantmentCost(skill))
		}
		ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("%s %-24s %s", selectionPrefix(index == g.selected), skill.Name, cost), 68, y+5)
	}
}

func (g *game) drawCredits(screenImage *ebiten.Image) {
	ebitenutil.DebugPrintAt(screenImage, "QUI SONT-ILS ?", 52, 124)
	ebitenutil.DebugPrintAt(screenImage, "Entree ou Echap : retour", 52, 152)
	ebitenutil.DrawRect(screenImage, 150, 190, 660, 300, color.RGBA{R: 18, G: 16, B: 22, A: 235})
	ebitenutil.DrawRect(screenImage, 150, 190, 660, 3, theme.goldBright)
	ebitenutil.DrawRect(screenImage, 480, 230, 1, 210, theme.gold)
	ebitenutil.DebugPrintAt(screenImage, "UNE BANDE QUI FAIT VIBRER LES SALLES", 215, 225)
	ebitenutil.DebugPrintAt(screenImage, "ABBA", 270, 300)
	ebitenutil.DebugPrintAt(screenImage, "MAMMA MIA !", 245, 345)
	ebitenutil.DebugPrintAt(screenImage, "LE MAITRE DE LA MISE EN SCENE", 525, 225)
	ebitenutil.DebugPrintAt(screenImage, "Steven", 555, 295)
	ebitenutil.DebugPrintAt(screenImage, "Spielberg", 535, 330)
	ebitenutil.DebugPrintAt(screenImage, "Appuyez sur Entree pour revenir", 330, 455)
}

func (g *game) drawCombat(screenImage *ebiten.Image) {
	title := "COMBAT D'ENTRAINEMENT"
	if g.combatAdventure {
		title = fmt.Sprintf("AVENTURE - SALLE %d", g.player.LastClearedRoom+1)
	}
	ebitenutil.DebugPrintAt(screenImage, title, 52, 124)
	ebitenutil.DebugPrintAt(screenImage, "Fleches haut/bas: choisir   Entree: confirmer", 52, 152)
	ebitenutil.DebugPrintAt(screenImage, g.player.Name, 100, 170)
	ebitenutil.DebugPrintAt(screenImage, g.monster.Name, 700, 170)
	drawHeroSprite(screenImage, 120, 180, g.player.Name, g.player.Class)
	drawMonsterSprite(screenImage, 680, 180, g.monster.Pattern)
	if g.player.Summon != nil && g.player.Summon.CurrentHP > 0 {
		ebitenutil.DebugPrintAt(screenImage, "SOLDAT INVOQUE", 395, 165)
		drawMonsterSprite(screenImage, 400, 180, "summon")
		ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("Niv. %d | %d PV | %d ATQ", g.player.Summon.Level, g.player.Summon.CurrentHP, g.player.Summon.Attack), 385, 345)
		ebitenutil.DebugPrintAt(screenImage, "VS", 570, 345)
	} else {
		ebitenutil.DebugPrintAt(screenImage, "VS", 450, 225)
	}
	drawBar(screenImage, "PV", g.player.CurrentHP, g.player.MaxHP, 100, 365, color.RGBA{R: 217, G: 92, B: 85, A: 255})
	drawBar(screenImage, "MANA", g.player.Mana, g.player.MaxMana, 100, 410, color.RGBA{R: 83, G: 151, B: 218, A: 255})
	drawExperienceBar(screenImage, g.player, 100, 448)
	drawBar(screenImage, "PV", g.monster.CurrentHP, g.monster.MaxHP, 600, 365, color.RGBA{R: 217, G: 92, B: 85, A: 255})
	drawPanel(screenImage, 410, 135, 100, 28, false)
	ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("Tour %d", g.combatTurn), 430, 145)
	for index, line := range wrapText(g.combatLog, 52) {
		ebitenutil.DebugPrintAt(screenImage, line, 80, 465+index*18)
	}
	ebitenutil.DebugPrintAt(screenImage, "ACTIONS", 650, 425)
	items := []string{"Attaque de base", "Sorts", "Inventaire", "Abandonner"}
	for index, item := range items {
		y := 445 + index*24
		drawChoiceRow(screenImage, 630, float64(y-5), 250, 22, index == g.selected)
		ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("%s %s", selectionPrefix(index == g.selected), item), 650, y)
	}
}

func wrapText(text string, width int) []string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return nil
	}
	lines := []string{words[0]}
	for _, word := range words[1:] {
		last := len(lines) - 1
		if len(lines[last])+1+len(word) <= width {
			lines[last] += " " + word
		} else {
			lines = append(lines, word)
		}
	}
	return lines
}

func selectionPrefix(selected bool) string {
	if selected {
		return ">"
	}
	return " "
}

func (g *game) drawCombatSpells(screenImage *ebiten.Image) {
	g.drawCombat(screenImage)
	drawPanel(screenImage, 250, 120, 470, 360, false)
	ebitenutil.DebugPrintAt(screenImage, "CHOISIR UN SORT", 285, 155)
	for index, skill := range g.player.Skill {
		y := 200 + index*38
		if index == g.selected {
			drawChoiceRow(screenImage, 275, float64(y-5), 390, 30, true)
		}
		ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("%s %-24s Mana: %d", selectionPrefix(index == g.selected), skill.Name, skill.ManaCost), 290, y+5)
	}
}

func (g *game) drawCombatInventory(screenImage *ebiten.Image) {
	g.drawCombat(screenImage)
	drawPanel(screenImage, 250, 120, 470, 360, false)
	ebitenutil.DebugPrintAt(screenImage, "CHOISIR UN OBJET", 285, 155)
	for index, item := range g.player.Inventory {
		y := 200 + index*34
		if index == g.selected {
			drawChoiceRow(screenImage, 275, float64(y-5), 390, 28, true)
		}
		ebitenutil.DebugPrintAt(screenImage, fmt.Sprintf("%s %-28s x%d", selectionPrefix(index == g.selected), library.ItemDisplayName(item), item.Quantity), 290, y+5)
	}
}

func drawMessage(screenImage *ebiten.Image, message string) {
	drawPanel(screenImage, 110, 210, 740, 150, false)
	for index, line := range strings.Split(message, "\n") {
		ebitenutil.DebugPrintAt(screenImage, line, 145, 235+index*20)
	}
	ebitenutil.DebugPrintAt(screenImage, "Appuyez sur Entree pour continuer", 290, 320)
}

func (g *game) Layout(_, _ int) (int, int) { return screenWidth, screenHeight }

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Lootstorm")
	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(newGame()); err != nil && err != ebiten.Termination {
		panic(err)
	}
}
