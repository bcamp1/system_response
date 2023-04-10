package main

import (
	"bcamp/rk4/rk"
	"bcamp/rk4/vec"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func spring(t float64, y vec.Vec) vec.Vec {
	m := 10.0
	b := 10.0
	k := 5000.0
	F := 0.0

	dx0 := y[1]
	dx1 := (F - b*y[1] - k*y[0]) / m

	return vec.Vec{dx0, dx1}
}

type Spring struct {
	start   Point
	stretch float64
	vel     float64
	length  float64
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func (s *Spring) draw(screen *ebiten.Image) {
	end := Point{s.start.x, s.start.y + s.length + s.stretch}
	vector.StrokeLine(screen, float32(s.start.x), float32(s.start.y), float32(end.x), float32(end.y), float32(min(max(4.0-(s.stretch/10.0), 2.0), 4.0)), color.White, true)
	vector.DrawFilledCircle(screen, float32(s.start.x), float32(end.y), 30, color.White, false)
	vector.StrokeLine(screen, float32(s.start.x-20), float32(s.start.y), float32(s.start.x+20), float32(s.start.y), 2.0, color.White, true)
}

type Game struct {
	spring Spring
	tps    int
}

type Point struct{ x, y float64 }

type Body struct {
	pos   Point
	color color.Color
	mass  float64
}

var mousePos Point
var mouseClick, pMouseClick bool

var c map[string]color.Color = map[string]color.Color{
	"red":   color.RGBA{255, 0, 0, 255},
	"black": color.RGBA{0, 0, 0, 255},
	"green": color.RGBA{0, 255, 0, 255},
	"blue":  color.RGBA{0, 0, 255, 255},
	"white": color.RGBA{255, 255, 255, 255},
}

func (g *Game) Update() error {
	pressed := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0)
	released := inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0)
	stretchPrev := g.spring.stretch

	if pressed && !mouseClick {
		mouseClick = true
	}

	if released && mouseClick {
		mouseClick = false
	}

	//fmt.Println(mouseClick)

	x, y := ebiten.CursorPosition()
	mousePos = Point{float64(x), float64(y)}

	deltaT := 1 / float64(g.tps)

	if mouseClick {
		g.spring.stretch = mousePos.y - g.spring.start.y - g.spring.length
		g.spring.vel = (g.spring.stretch - stretchPrev) / deltaT
	} else {
		next_state := rk.Step(spring, vec.Vec{g.spring.stretch, g.spring.vel}, 0, deltaT)
		g.spring.stretch = next_state[0]
		g.spring.vel = next_state[1]
	}

	return nil
}

// ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f, TPS: %.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
func (g *Game) Draw(screen *ebiten.Image) {
	g.spring.draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func main() {
	g := Game{
		spring: Spring{
			start:   Point{500, 100},
			stretch: 0,
			vel:     0,
			length:  300,
		},
		tps: 100,
	}

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetTPS(g.tps)
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
