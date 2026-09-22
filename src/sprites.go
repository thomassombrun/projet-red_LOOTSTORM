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
)

//go:embed assets/*.png
var spriteAssets embed.FS

// characterImagePaths maps a character's name to a hand-drawn sprite asset.
// Characters without an entry fall back to the procedural pixel sprite below.
var characterImagePaths = map[string]string{
	"Himiko Toga": "assets/himiko_toga.png",
}

var characterImageCache = map[string]*ebiten.Image{}

func loadCharacterImage(name string) *ebiten.Image {
	path, hasImage := characterImagePaths[name]
	if !hasImage {
		return nil
	}
	if img, cached := characterImageCache[name]; cached {
		return img
	}
	data, err := spriteAssets.ReadFile(path)
	if err != nil {
		return nil
	}
	decoded, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil
	}
	img := ebiten.NewImageFromImage(decoded)
	characterImageCache[name] = img
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
	'R': {R: 205, G: 67, B: 74, A: 255},
	'B': {R: 68, G: 137, B: 211, A: 255},
	'G': {R: 68, G: 182, B: 111, A: 255},
	'P': {R: 152, G: 92, B: 211, A: 255},
	'W': {R: 224, G: 230, B: 238, A: 255},
	'O': {R: 226, G: 127, B: 54, A: 255},
	'C': {R: 117, G: 79, B: 52, A: 255},
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

func drawCharacterSprite(screen *ebiten.Image, x, y, scale int, name, class string) {
	if img := loadCharacterImage(name); img != nil {
		drawScaledImage(screen, img, x, y, scale*11)
		return
	}
	pattern := "....KK....\n...KSSK...\n...KWWK...\n..KKKKKK..\n...BBBB...\n..BBBBBB..\n..B....B..\n.K......K.\n.K......K."
	switch class {
	case "Mage":
		pattern = "....PP....\n...PPPPP..\n...KWWK...\n..PPPPPP..\n...BBBB...\n..BBBBBB..\n..B....B..\n.K......K.\n.K......K."
	case "Guerrier", "Chevalier", "Barbare":
		pattern = "...KKKK...\n..KSSSSK..\n..KWWWWK..\n.KKKKKKKK.\n...RRRR...\n..RRRRRR..\n..R....R..\n.K......K.\n.K......K."
	case "Archer", "Assassin":
		pattern = "....KK....\n...KSSK...\n...KWWK...\n..KKKKKK..\n...GGGG...\n..GGGGGG..\n..G....G..\n.K......K.\n.K......K."
	}
	drawPixelSprite(screen, x, y, scale, pattern)
}

func drawMonsterSprite(screen *ebiten.Image, x, y, scale int, pattern string) {
	sprite := "....KK....\n...KRRK...\n..RRRRRR..\n.RRWWWWRR.\nRRRRRRRRRR\n.RRRRRRRR.\n..RRRRRR..\n...R..R...\n..K....K.."
	switch pattern {
	case "slime":
		sprite = "...GGGG...\n..GGGGGG..\n.GGWWWWGG.\nGGGGGGGGGG\nGGGGGGGGGG\n.GGGGGGGG.\n..GGGGGG..\n...G..G..."
	case "dragon":
		sprite = "..OO..OO..\n.OOOOOOOO.\nOOOWWWWOOO\nOOOOOOOOOO\n.OOOOOOOO.\n..OOOOOO..\n.OO.OO.OO.\nOO......OO\nK........K"
	case "duck":
		sprite = "...OOOO...\n..OOOOOO..\n.OOWWWWOO.\nOOOOOOOOOO\n.OOOOOOOO.\n..OOOOOO..\n...O..O...\n..K....K.."
	case "golem":
		sprite = "..KKKKKK..\n.KCCCCCCK.\nKCCWWWWCCK\nCCCCCCCCCC\nKCCCCCCCCK\n.KCCCCCCK.\n..KCCCCK..\n.K......K.\nK........K"
	}
	drawPixelSprite(screen, x, y, scale, sprite)
}

func drawItemSprite(screen *ebiten.Image, x, y, scale int, itemName string) {
	sprite := "..SSSS..\n.SSSSSS.\nSSSSSSSS\nSSSSSSSS\n.SSSSSS.\n..SSSS.."
	if strings.Contains(itemName, "Potion") {
		sprite = "..KK..\n..KK..\n.KKKK.\nKRRRRK\nKRRRRK\n.KKKK."
	} else if strings.Contains(itemName, "Livre") {
		sprite = "KWWWWK\nKBBBBK\nKWWWWK\nKBBBBK\nKWWWWK"
	} else if strings.Contains(itemName, "Plume") {
		sprite = "...B..\n..BB..\n.BBB..\nBBBB..\n.BB...\n.B...."
	} else if strings.Contains(itemName, "Cuir") || strings.Contains(itemName, "Peau") || strings.Contains(itemName, "Fourrure") {
		sprite = ".CCCC.\nCCCCCC\nCCCCCC\n.CCCC.\n..CC.."
	} else if strings.Contains(itemName, "Épée") || strings.Contains(itemName, "Lame") || strings.Contains(itemName, "Dague") || strings.Contains(itemName, "Griffe") {
		sprite = "...S..\n...S..\n..SS..\n.SS...\nSS....\n.KK..."
	}
	drawPixelSprite(screen, x, y, scale, sprite)
}

func drawResourceSprite(screen *ebiten.Image, x, y, scale int, resource string) {
	drawItemSprite(screen, x, y, scale, resource)
}
