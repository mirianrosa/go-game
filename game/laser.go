package game

import (
	"go-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Laser struct {
	image    *ebiten.Image
	position Vector
}

func NewLaser(position Vector) *Laser {

	image := assets.LaserSprite
	laserBounds := image.Bounds()
	metadeLaserBoundsLargura := float64(laserBounds.Dx() / 2) // Metade da largura da imagem do laser
	metadeLaserBoundsAltura := float64(laserBounds.Dy() / 2)  // Metade da altura da imagem do laser

	position.X -= metadeLaserBoundsLargura
	position.Y -= metadeLaserBoundsAltura

	return &Laser{
		image:    image,
		position: position,
	}
}

func (l *Laser) Update() {
	speed := 7.0

	l.position.Y -= speed
}

func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//Posição XY que a desenhara a imagem na tela
	op.GeoM.Translate(l.position.X, l.position.Y)

	screen.DrawImage(l.image, op)
}

func (l *Laser) LaserArea() DivRectangle {
	laserBounds := l.image.Bounds()

	return NewDivRectangle(l.position.X, l.position.Y, float64(laserBounds.Dx()), float64(laserBounds.Dy()))

}
