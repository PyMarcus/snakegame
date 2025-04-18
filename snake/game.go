package snake

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Snake      []Point
	Direction  Point
	food       Point
	lastUpdate time.Time
	GameOver   bool
}

func (g *Game) Update() error {
	if g.GameOver {
		return nil
	}

	g.keyboardListenerBind()

	if time.Since(g.lastUpdate) < time.Duration(GAME_SPEED) {
		return nil
	}

	g.lastUpdate = time.Now()

	g.updateSnake()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawSnakePoints(screen)
}

func (g *Game) Layout(outSideWidth, outSideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g Game) drawSnakePoints(screen *ebiten.Image) {
	for _, point := range g.Snake {
		vector.DrawFilledRect(
			screen,
			float32(point.x*RECT_SIZE),
			float32(point.y*RECT_SIZE),
			float32(RECT_SIZE),
			float32(RECT_SIZE),
			color.RGBA{0, 255, 0, 55},
			true,
		)
	}

	vector.DrawFilledRect(
		screen,
		float32(g.food.x*RECT_SIZE),
		float32(g.food.y*RECT_SIZE),
		float32(RECT_SIZE),
		float32(RECT_SIZE),
		color.RGBA{255, 0, 0, 55},
		true,
	)

	if g.GameOver {
		face := &text.GoTextFace{
			Source: MplusFaceSource,
			Size: 48,
		}
		
		text.Draw(screen, "Game Over", face, nil)		
	}
}

// startSnakePosition: sets snake on center of the screen.
func StartSnakePosition() []Point {
	return []Point{
		Point{
			x: SCREEN_WIDTH / RECT_SIZE / 2,
			y: SCREEN_HEIGHT / RECT_SIZE / 2,
		},
	}
}

func (g *Game) SpawnFood() {
	g.food = Point{x: rand.Intn(SCREEN_WIDTH / RECT_SIZE), y: rand.Intn(SCREEN_HEIGHT / RECT_SIZE)}
}

func (g *Game) updateSnake() {
	head := &g.Snake[0]

	newHead := Point{head.x + g.Direction.x, head.y + g.Direction.y}

	if g.isCollision(newHead){
		g.GameOver = true 
		return
	}

	if newHead == g.food {
		g.SpawnFood()
		g.Snake = append([]Point{newHead}, g.Snake[:len(g.Snake)]...)
	} else {
		g.Snake = append([]Point{newHead}, g.Snake[:len(g.Snake)-1]...)
	}

}

var (
	dirUp           = Point{x: 0, y: -1}
	dirDown         = Point{x: 0, y: 1}
	dirLeft         = Point{x: -1, y: 0}
	dirRight        = Point{x: 1, y: 0}
	MplusFaceSource *text.GoTextFaceSource
)

func (g *Game) keyboardListenerBind() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Direction = dirUp
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Direction = dirDown
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Direction = dirLeft
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.Direction = dirRight
	}
}

func (g *Game) isCollision(head Point) bool{
	if head.x < 0 || head.y < 0 || head.x >= SCREEN_WIDTH/RECT_SIZE || head.y >= SCREEN_HEIGHT/RECT_SIZE {
		return true
	}
	if g.collide(head){
		return true
	}
	
	return false
}

func (g *Game) collide(search Point) bool{
	for i, point := range g.Snake{
		if i == 0{
			continue
		}
		if point == search{
			return true
		}
	}
	return false
}