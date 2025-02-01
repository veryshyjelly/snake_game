package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	gridSize     = 40
	xBlocks      = screenWidth / gridSize
	yBlocks      = screenHeight / gridSize
	refreshRate  = time.Second / 7
)

var (
	right            = Point{x: 1, y: 0}
	left             = Point{x: -1, y: 0}
	up               = Point{x: 0, y: -1}
	down             = Point{x: 0, y: 1}
	mplusFaceSource  *text.GoTextFaceSource
	gameOverFontSize = 40
)
