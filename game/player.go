package game

import (
	"go-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image             *ebiten.Image
	position          Vector
	game              *Game
	laserLoadingTimer *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite

	naveBounds := image.Bounds()
	metadeNaveBounds := float64(naveBounds.Dx()) / 2 // Metade da largura da imagem da nave

	position := Vector{
		X: (screenWidth / 2) - metadeNaveBounds,
		Y: 500,
	}
	return &Player{
		image:             image,
		game:              game,
		position:          position,
		laserLoadingTimer: NewTimer(12),
	}
}

func (p *Player) Update() {

	speed := 8.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) { // Ir para a esquerda
		p.position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) { // Ir para a direita
		p.position.X += speed
	}

	p.laserLoadingTimer.Update()
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserLoadingTimer.IsReady() { // Atirar laser

		p.laserLoadingTimer.Reset()

		naveBounds := p.image.Bounds()
		metadeNaveBoundsPosicaoAtualX := float64(naveBounds.Dx()) / 2 // Metade da largura da imagem da nave na posição atual dela

		metadeNaveBoundsPosicaoAtualY := float64(naveBounds.Dy()) / 2 // Metade da largura da imagem da nave na posição atual dela

		spawnPosition := Vector{
			p.position.X + metadeNaveBoundsPosicaoAtualX,
			p.position.Y - metadeNaveBoundsPosicaoAtualY/2,
		}

		laser := NewLaser(spawnPosition)
		p.game.AddLasers(laser)

	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//Posição XY que a desenhara a imagem na tela
	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.image, op)
}
