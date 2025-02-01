package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Snake struct {
	body []Point
	dir  Point
	size int
}

func (s *Snake) Update() {
	var nextLoc = Point{
		x: s.body[0].x + s.dir.x,
		y: s.body[0].y + s.dir.y,
	}
	s.body = append([]Point{nextLoc}, (s.body[:s.size-1])...)
}

func (s *Snake) Draw(screen *ebiten.Image) {
	for _, b := range s.body {
		vector.DrawFilledRect(
			screen,
			float32(b.x*gridSize),
			float32(b.y*gridSize),
			gridSize, gridSize,
			color.White,
			true,
		)
	}
}
