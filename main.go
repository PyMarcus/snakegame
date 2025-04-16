package main

import (

	"github.com/PyMarcus/snakegame/snake"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	direction := snake.Point{}
	
	g := &snake.Game{Snake: snake.StartSnakePosition(), Direction: direction}
	g.SpawnFood()

	ebiten.SetWindowTitle("Snakegame")
	ebiten.SetWindowSize(snake.SCREEN_WIDTH, snake.SCREEN_HEIGHT)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
