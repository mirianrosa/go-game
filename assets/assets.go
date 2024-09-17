package assets

import (
	"embed"
	"image"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed *
var assets embed.FS

var PlayerSprite = mustLoadImage("player/player_ship1.png")
var KeyboardKeyLeftSprites = mustLoadImage("keyboard/keyboard_arrow_left.png")
var KeyboardKeyRightSprites = mustLoadImage("keyboard/keyboard_arrow_right.png")
var KeyboardKeySpaceSprites = mustLoadImage("keyboard/keyboard_space.png")
var KeyboardKeyEnterSprites = mustLoadImage("keyboard/keyboard_enter.png")

var MeteorSprites = mustLoadImages("meteors/*.png")
var LaserSprite = mustLoadImage("player/laser.png")
var GopherPlayer = mustLoadImage("player/gopher_player.png")
var StarsSprites = mustLoadImages("stars/*.png")
var PlanetsSprites = mustLoadImages("planets/*.png")

var ScoreFontBig = mustLoadFont("fonts/font.ttf", 27)
var FontUiBig = mustLoadFont("fonts/fontui.ttf", 27)
var ScoreFontSmall = mustLoadFont("fonts/font.ttf", 11)
var FontUiSmall = mustLoadFont("fonts/fontui.ttf", 11)

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func mustLoadImages(path string) []*ebiten.Image {
	matches, err := fs.Glob(assets, path)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = mustLoadImage(match)
	}

	return images
}

func mustLoadFont(name string, size float64) font.Face {
	f, err := assets.ReadFile(name)
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(f)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     100,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}

	return face
}
