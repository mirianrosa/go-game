package game

import (
	"go-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/exp/rand"
)

type Star struct {
	image    *ebiten.Image
	speed    float64
	position Vector
}

func FirstStar() *Star {
	image := assets.StarsSprites[rand.Intn(len(assets.StarsSprites))] // Seleciona aleatoriamente uma das imagens de estrelas dos assets

	speed := (rand.Float64() * 1) // Escolhe uma velocidade diferente para cada estrela que spawna

	position := Vector{
		X: rand.Float64() * screenWidth,  // Começa aleatoriamente dentro dos limites da largura da tela
		Y: rand.Float64() * screenHeight, // Começa aleatoriamente dentro dos limites da altura da tela
	}

	return &Star{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func NewStar() *Star {
	image := assets.StarsSprites[rand.Intn(len(assets.StarsSprites))] // Seleciona aleatoriamente uma das imagens de estrelas dos assets

	speed := (rand.Float64() * 1) // Escolhe uma velocidade diferente para cada estrela que spawna

	position := Vector{
		X: rand.Float64() * screenWidth, // Começa aleatoriamente dentro dos limites da largura da tela
		Y: -100,                         // Começa antes dos limites da tela
	}

	return &Star{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func (s *Star) Update() {
	s.position.Y += s.speed
}

func (s *Star) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//Posição XY que a desenhará a imagem na tela
	op.GeoM.Translate(s.position.X, s.position.Y)

	screen.DrawImage(s.image, op)
}
