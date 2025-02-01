package main

const (
	screenWidth  = 640
	screenHeight = 480
	gridSize     = 20
)

var (
	right = Point{x: 1, y: 0}
	left  = Point{x: -1, y: 0}
	up    = Point{x: 0, y: -1}
	down  = Point{x: 0, y: 1}
)
