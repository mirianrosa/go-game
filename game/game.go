package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player           *Player
	lasers           []*Laser
	meteors          []*Meteor
	meteorSpawnTimer *Timer
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(24),
	}
	player := NewPlayer(g)
	g.player = player
	return g
}

// A lib ebiten garante que o jogo rode em 60 FPS
// As funções Update e Draw atualizam 60 vezes por segundo.
func (g *Game) Update() error {
	g.player.Update()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, l := range g.lasers {
		l.Update()
	}

	for _, m := range g.meteors {
		if m.MeteorArea().Intersects(g.player.PlayerArea()) {
			g.Reset()
		}
	}

	for meteorIndex, meteorUnity := range g.meteors {
		for laserIndex, laserUnity := range g.lasers {
			if meteorUnity.MeteorArea().Intersects(laserUnity.LaserArea()) {
				g.meteors = append(g.meteors[:meteorIndex], g.meteors[meteorIndex+1:]...)
				g.lasers = append(g.lasers[:laserIndex], g.lasers[laserIndex+1:]...)
			}
		}

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	for _, l := range g.lasers {
		l.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}

func (g *Game) CollisionDetected() {
	//To do
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.lasers = nil
	g.meteorSpawnTimer.Reset()
}
