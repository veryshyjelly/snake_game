package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Food = Point

func GenerateFood(s Snake) Food {
again:
	for {
		var food = Food{
			x: rand.Intn(screenWidth / gridSize),
			y: rand.Intn(screenHeight / gridSize),
		}

		for _, b := range s.body {
			if b == food {
				continue again
			}
		}
		return food
	}
}

func (f *Food) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen,
		float32(f.x*gridSize),
		float32(f.y*gridSize),
		gridSize, gridSize,
		color.RGBA{R: 255, A: 255},
		false,
	)
}
