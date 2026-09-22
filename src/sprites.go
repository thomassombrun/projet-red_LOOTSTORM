package main

import (
	"bytes"
	"embed"
	"image"
	"image/color"
	_ "image/png"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "golang.org/x/image/webp"
)

// spriteAssets embeds everything under assets/ so new hand-drawn sprites are
// picked up automatically just by dropping a correctly named PNG in the
// matching subfolder (characters/, monsters/, items/) — no code change needed.
//
//go:embed all:assets
var spriteAssets embed.FS

var imageCache = map[string]*ebiten.Image{}

// slugify turns a display name ("Himiko Toga", "Potion de vie") into the
// lowercase, underscore-separated filename it must be saved as (himiko_toga.png).
func slugify(name string) string {
	var builder strings.Builder
	lastUnderscore := false
	for _, r := range strings.ToLower(name) {
		switch {
		case r >= 'a' && r <= 'z' || r >= '0' && r <= '9':
			builder.WriteRune(r)
			lastUnderscore = false
		case strings.ContainsRune("éèêë", r):
			builder.WriteRune('e')
			lastUnderscore = false
		case strings.ContainsRune("àâ", r):
			builder.WriteRune('a')
			lastUnderscore = false
		case strings.ContainsRune("îï", r):
			builder.WriteRune('i')
			lastUnderscore = false
		case strings.ContainsRune("ôö", r):
			builder.WriteRune('o')
			lastUnderscore = false
		case strings.ContainsRune("ùûü", r):
			builder.WriteRune('u')
			lastUnderscore = false
		case r == 'ç':
			builder.WriteRune('c')
			lastUnderscore = false
		default:
			if !lastUnderscore && builder.Len() > 0 {
				builder.WriteRune('_')
				lastUnderscore = true
			}
		}
	}
	return strings.Trim(builder.String(), "_")
}

// loadEmbeddedImage returns the decoded image at path, or nil if it doesn't
// exist. Misses are cached too so missing files aren't re-read every frame.
func loadEmbeddedImage(path string) *ebiten.Image {
	if img, cached := imageCache[path]; cached {
		return img
	}
	data, err := spriteAssets.ReadFile(path)
	if err != nil {
		imageCache[path] = nil
		return nil
	}
	decoded, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		imageCache[path] = nil
		return nil
	}
	img := ebiten.NewImageFromImage(decoded)
	imageCache[path] = img
	return img
}

func drawScaledImage(screen *ebiten.Image, img *ebiten.Image, x, y, targetSize int) {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	if width == 0 || height == 0 {
		return
	}
	factor := float64(targetSize) / float64(width)
	if height > width {
		factor = float64(targetSize) / float64(height)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(factor, factor)
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img, op)
}

var spritePalette = map[rune]color.RGBA{
	'.': {R: 0, G: 0, B: 0, A: 0},
	'K': {R: 24, G: 28, B: 38, A: 255},
	'S': {R: 229, G: 184, B: 105, A: 255},
	'H': {R: 194, G: 148, B: 86, A: 255},
	'R': {R: 205, G: 67, B: 74, A: 255},
	'M': {R: 140, G: 40, B: 50, A: 255},
	'B': {R: 68, G: 137, B: 211, A: 255},
	'N': {R: 40, G: 90, B: 150, A: 255},
	'G': {R: 68, G: 182, B: 111, A: 255},
	'F': {R: 40, G: 130, B: 80, A: 255},
	'P': {R: 152, G: 92, B: 211, A: 255},
	'U': {R: 100, G: 60, B: 160, A: 255},
	'W': {R: 224, G: 230, B: 238, A: 255},
	'O': {R: 226, G: 127, B: 54, A: 255},
	'D': {R: 170, G: 90, B: 35, A: 255},
	'C': {R: 117, G: 79, B: 52, A: 255},
	'J': {R: 80, G: 55, B: 35, A: 255},
	'Y': {R: 240, G: 200, B: 80, A: 255},
	'Z': {R: 180, G: 140, B: 40, A: 255},
	'A': {R: 180, G: 186, B: 196, A: 255},
	'T': {R: 110, G: 116, B: 128, A: 255},
	'E': {R: 225, G: 220, B: 200, A: 255},
	'V': {R: 140, G: 220, B: 220, A: 255},
	'X': {R: 80, G: 160, B: 160, A: 255},
}

func drawPixelSprite(screen *ebiten.Image, x, y, scale int, pattern string) {
	lines := strings.Split(strings.TrimSpace(pattern), "\n")
	for row, line := range lines {
		for column, symbol := range line {
			color, exists := spritePalette[symbol]
			if !exists || color.A == 0 {
				continue
			}
			ebitenutil.DrawRect(screen, float64(x+column*scale), float64(y+row*scale), float64(scale), float64(scale), color)
		}
	}
}

// recolor swaps the '1' (base tone) and '2' (shadow tone) placeholders of a
// shared silhouette for the two colors of a specific palette/family.
func recolor(pattern string, base, shadow rune) string {
	return strings.NewReplacer("1", string(base), "2", string(shadow)).Replace(pattern)
}

func drawCharacterSprite(screen *ebiten.Image, x, y, scale int, name, class string) {
	if img := loadEmbeddedImage("assets/characters/" + slugify(name) + ".png"); img != nil {
		drawScaledImage(screen, img, x, y, scale*11)
		return
	}
	pattern := characterPatterns["Guerrier"]
	if found, ok := characterPatterns[class]; ok {
		pattern = found
	}
	drawPixelSprite(screen, x, y, scale, pattern)
}

var characterPatterns = map[string]string{
	"Guerrier":   "....RR....\n...AAAAA..\n..AKSSKA..\n..AKWSWKA.\n..AAAAAAA.\n.ARRRRRRA.\n.ARRRRRRA.\n.ARRRRRRA.\n..ATTTTA..\n..KTT.TTK.\n.KK.....K.",
	"Chevalier":  "...AAAA...\n..AAAAAA..\n..AKSSKA..\n..AKWSWKA.\n.AAAAAAAA.\n.ABBBBBBA.\n.ABBYYBBA.\n.ABBBBBBA.\n..ANNNNA..\n..KNN.NNK.\n.KK.....K.",
	"Barbare":    ".O..OO..O.\n.OOOOOOOO.\n..HSSSSH..\n..HWSWSH..\n.HHHHHHHH.\n.CJCCCCJC.\n.CJHHHHJC.\n.CJCCCCJC.\n..JCCCCJ..\n..KJJ.JJK.\n.KK.....K.",
	"Mage":       "...PPPP...\n..PPPPPP..\n..PKSSKP..\n..PKWSWKP.\n.PPPPPPPP.\n.PUUUUUUP.\n.UUYYYYUU.\n.UUUUUUUU.\n..UUUUUU..\n..KUU.UUK.\n.KK.....K.",
	"Archer":     "....GG....\n...GGGGG..\n..GKSSKG..\n..GKWSWKG.\n.GGGGGGGG.\n.GFFFFFFG.\n.GFYY..FG.\n.GFFFFFFG.\n..FCCCCF..\n..KCC.CCK.\n.KK.....K.",
	"Assassin":   "....KK....\n...KGGGK..\n..GKSSKG..\n..GKWSWKG.\n.GGGGGGGG.\n.GKKKKKKG.\n.GK....KG.\n.GKKKKKKG.\n..K....K..\n..KTT.TTK.\n.KK.....K.",
	"Samourai":   "....KK....\n...KWWWK..\n..WKSSKW..\n..WKWSWKW.\n.WWWWWWWW.\n.WRRRRRRW.\n.WRWW..RW.\n.WRRRRRRW.\n..RKKKKR..\n..KJJ.JJK.\n.KK.....K.",
	"Clerc":      "....YY....\n...WWWWW..\n..WKSSKW..\n..WKWSWKW.\n.WWWWWWWW.\n.WYWWWWYW.\n.WWWYYWWW.\n.WYWWWWYW.\n..YWWWWY..\n..KYY.YYK.\n.KK.....K.",
	"Invocateur": "....UU....\n...UPPPU..\n..PKSSKP..\n..PKWSWKP.\n.PPPPPPPP.\n.PUUUUUUP.\n.UUY..YUU.\n.UUUUUUUU.\n..YZZZZY..\n.YZ.YY.ZY.\n.KK.....K.",
}

func drawMonsterSprite(screen *ebiten.Image, x, y, scale int, pattern string) {
	if img := loadEmbeddedImage("assets/monsters/" + slugify(pattern) + ".png"); img != nil {
		drawScaledImage(screen, img, x, y, scale*11)
		return
	}
	sprite, ok := monsterPatterns[pattern]
	if !ok {
		sprite = monsterPatterns["goblin"]
	}
	drawPixelSprite(screen, x, y, scale, sprite)
}

var monsterPatterns = map[string]string{
	"goblin":   "....KK....\n...GGGGG..\n..GKRRKG..\n.GGGGGGGG.\nGGGWWWWGG.\n.GGGGGGGG.\n..GFFFFG..\n...F..F...\n..K....K..",
	"slime":    "..........\n...GGGG...\n..GGGGGG..\n.GGWWWWGG.\nGGGGGGGGGG\nGFFFFFFFFG\n.GFFFFFFG.\n..GFFFFG..\n...F..F...",
	"ghost":    "...VVVV...\n..VVVVVV..\n.VVKWWKVV.\nVVVVVVVVVV\nVVXXXXXXVV\n.VVXXXXVV.\n..VV.VVV..\n.V.V..V.V.\nV.V....V.V",
	"golem":    "..AAAAAA..\n.ATTTTTTA.\nATTKWWKTTA\nTTTTTTTTTT\nATTTTTTTTA\n.ATTTTTTA.\n..ATTTTA..\n.A......A.\nA........A",
	"troll":    "..CCCCCC..\n.CJJJJJJC.\nCJKRRKJJC.\nCCCCCCCCCC\nJCCCCCCCCJ\n.JCCCCCCJ.\n..JCCCCJ..\n.J......J.\nJ........J",
	"duck":     "...OOOO...\n..ODDDDO..\n.ODWWWDDO.\nOOOOOOOOOO\nDOOOOOOOD.\n.DOOOOOD..\n...O..O...\n..K....K..",
	"skeleton": "..EEEEEE..\n.EKKKKKE..\nEEKWWKEE..\nEEEEEEEEEE\n.ETTTTTE..\nE.ETTTE.E.\n..E.T.E...\n.E......E.\nE........E",
	"wolf":     "..TT..TT..\n.TCCCCCCT.\nTCKRRKCT..\nTCCCCCCCT.\nCCCCCCCCCC\n.CCWWWWCC.\n..CC..CC..\n.C......C.",
	"wizard":   "....PP....\n...PPPPP..\n..PKWWKP..\nPPPPPPPPPP\n.PUUUUUUP.\nPUUYY.YYUP\n.UUUUUUUU.\n..U....U..\n.K......K.",
	"dragon":   "..OO..OO..\n.OMMMMMMO.\nOMMWWWWMMO\nOOOOOOOOOO\nMOOOOOOOOM\n.MOOOOOOM.\n.OM.OO.MO.\nOM......MO\nM........M",
	"summon":   "....BB....\n...BNNNB..\n..NKWWKN..\nNNNNNNNNNN\n.NBBBBBBN.\n..NBBBBN..\n...N..N...\n..K....K..",
}

func itemFamilyTint(lowerName string) (rune, rune) {
	switch {
	case strings.Contains(lowerName, "gobelin"):
		return 'G', 'F'
	case strings.Contains(lowerName, "slime"):
		return 'G', 'X'
	case strings.Contains(lowerName, "spectr"):
		return 'V', 'X'
	case strings.Contains(lowerName, "golem"):
		return 'A', 'T'
	case strings.Contains(lowerName, "troll"):
		return 'C', 'J'
	case strings.Contains(lowerName, "canard"):
		return 'O', 'D'
	case strings.Contains(lowerName, "squelette"):
		return 'E', 'T'
	case strings.Contains(lowerName, "loup"):
		return 'T', 'K'
	case strings.Contains(lowerName, "sorcier"):
		return 'P', 'U'
	case strings.Contains(lowerName, "dragon"):
		return 'R', 'M'
	case strings.Contains(lowerName, "aventurier"):
		return 'C', 'J'
	default:
		return 'A', 'T'
	}
}

const (
	helmetShape = "..111..\n.11111.\n1111111\n122222K\nK.....K"
	chestShape  = "1111111\n1122221\n1111111\n1122221\nK.....K\nK.....K"
	bootsShape  = "11.11\n11.11\n22.22\nK...K"
	bladeShape  = "..1..\n..1..\n..1..\n.111.\n..2.."
	hammerShape = ".111.\n11111\n12121\n..2..\n..2..\n..2.."
	staffShape  = "..1..\n..1..\n..2..\n..2..\n..2.."
	bowShape    = "1....\n.1...\n..1..\n.1...\n1...."
	clawShape   = "1.1.1\n.1.1.\n..1.."
	beakShape   = "..1.\n.11.\n1111\n..2."
)

func potionShape(base, shadow rune) string {
	return recolor("..KK..\n..KK..\n.K11K.\nK111K.\nK1122K\n.K22K.", base, shadow)
}

func spellbookShape(lowerName string) string {
	base, shadow := rune('B'), rune('N')
	switch {
	case strings.Contains(lowerName, "feu"):
		base, shadow = 'R', 'M'
	case strings.Contains(lowerName, "soin"), strings.Contains(lowerName, "régénération"), strings.Contains(lowerName, "regeneration"):
		base, shadow = 'G', 'F'
	case strings.Contains(lowerName, "poison"):
		base, shadow = 'P', 'U'
	case strings.Contains(lowerName, "brûlure"), strings.Contains(lowerName, "brulure"):
		base, shadow = 'O', 'D'
	case strings.Contains(lowerName, "éclair"), strings.Contains(lowerName, "eclair"):
		base, shadow = 'Y', 'Z'
	case strings.Contains(lowerName, "barrière"), strings.Contains(lowerName, "barriere"):
		base, shadow = 'W', 'Y'
	}
	return recolor("KKKKKKK\nK1W1W1K\nK111111\nK122222\nKKKKKKK", base, shadow)
}

func drawItemSprite(screen *ebiten.Image, x, y, scale int, itemName string) {
	if img := loadEmbeddedImage("assets/items/" + slugify(itemName) + ".png"); img != nil {
		drawScaledImage(screen, img, x, y, scale*8)
		return
	}
	lower := strings.ToLower(itemName)
	sprite := "..YY..\n.YYYY.\nYYYYYY\n.YYYY.\n..YY.."
	switch {
	case strings.Contains(lower, "potion de vie"):
		sprite = potionShape('R', 'M')
	case strings.Contains(lower, "potion de poison"):
		sprite = potionShape('G', 'F')
	case strings.Contains(lower, "potion de mana"):
		sprite = potionShape('B', 'N')
	case strings.Contains(lower, "livre"):
		sprite = spellbookShape(lower)
	case strings.Contains(lower, "amelioration"):
		sprite = "KKKKK\nKYYYK\nKYYYK\nKZZZK\nKKKKK"
	case strings.Contains(lower, "casque"), strings.Contains(lower, "heaume"), strings.Contains(lower, "capuche"), strings.Contains(lower, "carapace"), strings.Contains(lower, "chapeau"):
		base, shadow := itemFamilyTint(lower)
		sprite = recolor(helmetShape, base, shadow)
	case strings.Contains(lower, "tunique"):
		base, shadow := itemFamilyTint(lower)
		sprite = recolor(chestShape, base, shadow)
	case strings.Contains(lower, "bottes"):
		base, shadow := itemFamilyTint(lower)
		sprite = recolor(bootsShape, base, shadow)
	case strings.Contains(lower, "marteau"), strings.Contains(lower, "massue"):
		base, shadow := itemFamilyTint(lower)
		sprite = recolor(hammerShape, base, shadow)
	case strings.Contains(lower, "bâton"), strings.Contains(lower, "baton"):
		base, shadow := itemFamilyTint(lower)
		sprite = recolor(staffShape, base, shadow)
	case strings.Contains(lower, "arc"):
		base, shadow := itemFamilyTint(lower)
		sprite = recolor(bowShape, base, shadow)
	case strings.Contains(lower, "griffe"), strings.Contains(lower, "crocs"):
		base, shadow := itemFamilyTint(lower)
		sprite = recolor(clawShape, base, shadow)
	case strings.Contains(lower, "bec"):
		base, shadow := itemFamilyTint(lower)
		sprite = recolor(beakShape, base, shadow)
	case strings.Contains(lower, "épée"), strings.Contains(lower, "epee"), strings.Contains(lower, "lame"), strings.Contains(lower, "dague"):
		base, shadow := itemFamilyTint(lower)
		sprite = recolor(bladeShape, base, shadow)
	case strings.Contains(lower, "cuir"), strings.Contains(lower, "peau"), strings.Contains(lower, "fourrure"):
		sprite = ".CCCC.\nCCJJCC\nCCCCCC\nCJJJJC\n.CCCC.\n..CC.."
	case strings.Contains(lower, "plume"):
		sprite = "...B..\n..BB..\n.BNB..\nBBNB..\n.BN...\n.B...."
	}
	drawPixelSprite(screen, x, y, scale, sprite)
}

func drawResourceSprite(screen *ebiten.Image, x, y, scale int, resource string) {
	drawItemSprite(screen, x, y, scale, resource)
}
