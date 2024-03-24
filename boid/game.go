//nolint:typecheck
package boid

import (
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth, screenHeight = 720, 480
	boidCount                 = 500
	viewRadius                = 13
	adjustmentRate            = 0.015
)

var (
	green   = color.RGBA{R: 60, G: 179, B: 133, A: 255}
	boids   = [boidCount]*Boid{}
	boidMap = [screenWidth + 1][screenHeight + 1]int{}
	lock    = sync.RWMutex{}
)

type Game struct{}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, b := range boids {
		screen.Set(int(b.position.X+1), int(b.position.Y), green)
		screen.Set(int(b.position.X-1), int(b.position.Y), green)
		screen.Set(int(b.position.X), int(b.position.Y-1), green)
		screen.Set(int(b.position.X), int(b.position.Y+1), green)
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	return nil
}

func Run() {
	for i, row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}

	for i := 0; i < boidCount; i++ {
		b := createBoid(i)
		boids[i] = b
		boidMap[int(b.position.X)][int(b.position.Y)] = b.id

		go boids[i].start()
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Boids in a box")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
