package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	snake      Snake
	gameOver   bool
	lastUpdate time.Time
}

func NewGame() Game {
	return Game{
		snake: Snake{
			body: []Point{{screenWidth / gridSize / 2, screenHeight / gridSize / 2}},
			dir:  right,
			size: 1,
		},
	}
}

func (g *Game) Update() error {
	if time.Since(g.lastUpdate) < time.Second/6 {
		return nil
	}
	g.snake.Update()
	g.lastUpdate = time.Now()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.snake.Draw(screen)
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return screenWidth, screenHeight
}
