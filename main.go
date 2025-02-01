package main

import "github.com/hajimehoshi/ebiten/v2"

type Point struct {
	x, y int
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake Game")
	game := NewGame()
	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
