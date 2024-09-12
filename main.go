package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

func (game *Game) Update() error {
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {

}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {

	game := &Game{}

	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}

}
