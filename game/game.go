package game

import (
	"fmt"
	"go-game/assets"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	player           *Player
	stars            []*Star
	starSpawnTimer   *Timer
	lasers           []*Laser
	meteors          []*Meteor
	meteorSpawnTimer *Timer
	score            int
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(24),
		starSpawnTimer:   NewTimer(70),
	}
	player := NewPlayer(g)
	g.player = player

	for i := 1; i <= 10; i++ {
		s := FirstStar()
		g.stars = append(g.stars, s)
	}

	return g
}

// A lib ebiten garante que o jogo rode em 60 FPS
// As funções Update e Draw atualizam 60 vezes por segundo.
func (g *Game) Update() error {
	g.player.Update()

	g.starSpawnTimer.Update()
	if g.starSpawnTimer.IsReady() {
		g.starSpawnTimer.Reset()

		s := NewStar()
		g.stars = append(g.stars, s)
	}

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, s := range g.stars {
		s.Update()
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
				g.score += 1
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, s := range g.stars {
		s.Draw(screen)
	}

	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	for _, l := range g.lasers {
		l.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("Pontos: %d", g.score), assets.ScoreFont, 20, 100, color.White)
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
	g.starSpawnTimer.Reset()
	g.score = 0
	g.stars = nil
	for i := 1; i <= 10; i++ {
		s := FirstStar()
		g.stars = append(g.stars, s)
	}
}
