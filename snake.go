package main

import (
	"errors"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Snake struct {
	body []Point
	dir  Point
	size int
}

func (s *Snake) Update(food Food) error {
	var nextLoc = Point{
		x: (s.body[0].x + s.dir.x + xBlocks) % xBlocks,
		y: (s.body[0].y + s.dir.y + yBlocks) % yBlocks,
	}
	if nextLoc == food {
		s.size += 1
	}
	for _, b := range s.body {
		if b == nextLoc {
			return errors.New("Snake bit itself")
		}
	}
	s.body = append([]Point{nextLoc}, (s.body[:s.size-1])...)
	return nil
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

	// Draw Eyes
	var head = s.body[0]
	var cx1, cy1, cx2, cy2 float32
	if s.dir == left {
		cx1 = float32(head.x*gridSize + gridSize/4)
		cy1 = float32(head.y*gridSize + gridSize/4)
		cx2 = float32(head.x*gridSize + gridSize/4)
		cy2 = float32(head.y*gridSize + 3*gridSize/4)
	} else if s.dir == right {
		cx1 = float32(head.x*gridSize + 3*gridSize/4)
		cy1 = float32(head.y*gridSize + gridSize/4)
		cx2 = float32(head.x*gridSize + 3*gridSize/4)
		cy2 = float32(head.y*gridSize + 3*gridSize/4)
	} else if s.dir == up {
		cx1 = float32(head.x*gridSize + 3*gridSize/4)
		cy1 = float32(head.y*gridSize + gridSize/4)
		cx2 = float32(head.x*gridSize + gridSize/4)
		cy2 = float32(head.y*gridSize + gridSize/4)
	} else {
		cx1 = float32(head.x*gridSize + 3*gridSize/4)
		cy1 = float32(head.y*gridSize + 3*gridSize/4)
		cx2 = float32(head.x*gridSize + gridSize/4)
		cy2 = float32(head.y*gridSize + 3*gridSize/4)
	}

	vector.DrawFilledCircle(
		screen, cx1, cy1,
		gridSize/9,
		color.Black,
		false,
	)
	vector.DrawFilledCircle(
		screen, cx2, cy2,
		gridSize/9,
		color.Black,
		false,
	)
}

func (s *Snake) Controls() {
	// Keyboard controls
	if (s.size == 1 || s.dir != down) &&
		(ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp)) {
		s.dir = up
	} else if (s.size == 1 || s.dir != up) &&
		(ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown)) {
		s.dir = down
	} else if (s.size == 1 || s.dir != right) &&
		(ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft)) {
		s.dir = left
	} else if (s.size == 1 || s.dir != left) &&
		(ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight)) {
		s.dir = right
	}
}
