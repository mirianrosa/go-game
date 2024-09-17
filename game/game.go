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
	bestScore        int
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(24),
		starSpawnTimer:   NewTimer(10),
	}
	player := NewPlayer(g)
	g.player = player

	for i := 1; i <= 15; i++ {
		s := NewStar(true)
		g.stars = append(g.stars, s)
	}

	g.bestScore = 0

	return g
}

// A lib ebiten garante que o jogo rode em 60 FPS
// As funções Update e Draw atualizam 60 vezes por segundo.
func (g *Game) Update() error {
	g.player.Update()

	g.starSpawnTimer.Update()
	if g.starSpawnTimer.IsReady() {
		g.starSpawnTimer.Reset()

		s := NewStar(false)
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

	for meteorIndex := len(g.meteors) - 1; meteorIndex >= 0; meteorIndex-- {
		meteorUnity := g.meteors[meteorIndex]

		if meteorUnity.position.Y > screenHeight { // Se o meteoro sair da área visível da tela pela parte inferior, ele pode ser removido
			g.meteors = append(g.meteors[:meteorIndex], g.meteors[meteorIndex+1:]...)
			continue
		}

		for laserIndex := len(g.lasers) - 1; laserIndex >= 0; laserIndex-- {
			laserUnity := g.lasers[laserIndex]
			if laserUnity.position.Y < -50 { // Se o laser sair da área visível da tela pela parte superior, ele pode ser removido
				g.lasers = append(g.lasers[:laserIndex], g.lasers[laserIndex+1:]...)
				continue
			}
			if meteorUnity.MeteorArea().Intersects(laserUnity.LaserArea()) {

				//fmt.Println("Colisão detectada entre meteoro #", meteorIndex, "e laser #", laserIndex)
				//fmt.Println("Range de meteoros: ", g.meteors)
				g.meteors = append(g.meteors[:meteorIndex], g.meteors[meteorIndex+1:]...)
				g.lasers = append(g.lasers[:laserIndex], g.lasers[laserIndex+1:]...)
				g.score += 1
				break
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

	text.Draw(screen, fmt.Sprintf("Seu Melhor: %d", g.bestScore), assets.ScoreFontBig, 15, 50, color.White)
	text.Draw(screen, fmt.Sprintf("Pontos: %d", g.score), assets.ScoreFontBig, 15, 100, color.RGBA{R: 255, G: 255, B: 0, A: 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}

func (g *Game) PlayerCollisionDetected() {
	// To do
	// Player loses. Pause every movement and creation of stars, meteors, player, lasers.
	// Ask if want to try again? Press any key to continue...
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.lasers = nil
	g.meteorSpawnTimer.Reset()
	g.starSpawnTimer.Reset()
	g.stars = nil
	for i := 1; i <= 15; i++ {
		s := NewStar(true)
		g.stars = append(g.stars, s)
	}

	if g.score > g.bestScore {
		g.bestScore = g.score
	}
	g.score = 0
}
