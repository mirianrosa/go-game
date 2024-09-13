package game

import (
	"go-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/exp/rand"
)

type Meteor struct {
	image    *ebiten.Image
	speed    float64
	position Vector
}

func NewMeteor() *Meteor {
	image := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))] // Seleciona aleatoriamente uma das imagens de meteoro dos assets

	speed := (rand.Float64() * 13) // Escolhe uma velocidade diferente para cada meteoro que spawna

	position := Vector{
		X: rand.Float64() * screenWidth, // Começa aleatoriamente dentro dos limites da largura da tela
		Y: -100,                         // Começa sempre acima do limite de altura da tela
	}

	return &Meteor{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func (m *Meteor) Update() {
	m.position.Y += m.speed
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//Posição XY que a desenhara a imagem na tela
	op.GeoM.Translate(m.position.X, m.position.Y)

	screen.DrawImage(m.image, op)
}

func (m *Meteor) MeteorArea() DivRectangle {
	meteorBounds := m.image.Bounds()

	return NewDivRectangle(m.position.X, m.position.Y, float64(meteorBounds.Dx()), float64(meteorBounds.Dy()))

}
