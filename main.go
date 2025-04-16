package main

import (
	"bytes"

	"github.com/PyMarcus/snakegame/snake"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

)

func main() {
	s, _ := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	snake.MplusFaceSource = s

	direction := snake.Point{}

	g := &snake.Game{Snake: snake.StartSnakePosition(), Direction: direction}
	g.SpawnFood()

	ebiten.SetWindowTitle("Snakegame")
	ebiten.SetWindowSize(snake.SCREEN_WIDTH, snake.SCREEN_HEIGHT)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
