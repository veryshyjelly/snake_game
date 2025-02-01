package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	snake      Snake
	food       Food
	gameOver   bool
	points     int
	lastUpdate time.Time
}

func NewGame() Game {
	snake := Snake{
		body: []Point{{screenWidth / gridSize / 2, screenHeight / gridSize / 2}},
		dir:  right,
		size: 1,
	}
	return Game{
		snake: snake,
		food:  GenerateFood(snake),
	}
}

func (g *Game) Update() error {
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			ga := NewGame()
			g.snake = ga.snake
			g.food = ga.food
			g.gameOver = false
		}
		return nil
	}

	g.snake.Controls()

	if time.Since(g.lastUpdate) < refreshRate {
		return nil
	}

	if err := g.snake.Update(g.food); err != nil {
		g.gameOver = true
		return nil
	}

	if g.snake.body[0] == g.food {
		g.points += 1
		g.food = GenerateFood(g.snake)
	}
	g.lastUpdate = time.Now()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < screenWidth/gridSize; i++ {
		for j := 0; j < screenHeight/gridSize; j++ {
			if (i+j)%2 == 0 {
				vector.DrawFilledRect(
					screen,
					float32(i*gridSize),
					float32(j*gridSize),
					gridSize,
					gridSize,
					color.RGBA{G: 155, R: 50},
					false,
				)
			} else {
				vector.DrawFilledRect(
					screen,
					float32(i*gridSize),
					float32(j*gridSize),
					gridSize,
					gridSize,
					color.RGBA{G: 255, R: 50},
					false,
				)
			}
		}
	}

	g.snake.Draw(screen)
	g.food.Draw(screen)

	op := &text.DrawOptions{}
	face := &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   24,
	}

	points := fmt.Sprintf("Score: %v", g.points)
	w, _ := text.Measure(points, face, 0)
	op.GeoM.Translate(screenWidth-w-gridSize, 0)
	op.ColorScale.ScaleWithColor(color.Black)
	text.Draw(screen, points, face, op)

	if g.gameOver {
		op := &text.DrawOptions{}
		op.ColorScale.ScaleWithColor(color.Black)
		face.Size = float64(gameOverFontSize)
		w, h := text.Measure("Game Over!", face, 0)
		op.GeoM.Translate(screenWidth/2-w/2, screenHeight/2-h)
		text.Draw(screen, "Game Over!", face, op)
		return
	}

}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return screenWidth, screenHeight
}
