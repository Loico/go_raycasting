package main

import (
	"fmt"
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 800
	screenHeight = 800

	targetTicksPerSecond = 60
)

type coordinates struct {
	x float64
	y float64
}

var delta float64

type ray struct {
	pos coordinates
	dir float64
}

func cast(b wall, r ray) (ret bool, pt coordinates) {
	x1 := b.a.x
	y1 := b.a.y
	x2 := b.b.x
	y2 := b.b.y
	x3 := r.pos.x
	y3 := r.pos.y
	x4 := r.pos.x + math.Cos(r.dir)
	y4 := r.pos.y + math.Sin(r.dir)

	den := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if den == 0 {
		return false, pt
	}
	var t, u float64
	t = ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) / den
	u = (-((x1-x2)*(y1-y3) - (y1-y2)*(x1-x3))) / den
	if t > 0 && t < 1 && u > 0 {
		pt.x = x1 + t*(x2-x1)
		pt.y = y1 + t*(y2-y1)
		return true, pt
	} else {
		return false, pt
	}
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Initialisation SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Go Raycasting",
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Initialisation window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initialisation renderer:", err)
		return
	}
	defer renderer.Destroy()

	elements = append(elements, newPlayer(renderer))

	readMap()

	for {
		frameStartTime := time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		for _, elem := range elements {
			if elem.active {
				err = elem.update()
				if err != nil {
					fmt.Println("updating element:", err)
					return
				}
				err = elem.draw(renderer)
				if err != nil {
					fmt.Println("drawing element:", err)
				}
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		for _, b := range walls {
			renderer.DrawLine(int32(b.a.x), int32(b.a.y), int32(b.b.x), int32(b.b.y))
		}

		renderer.Present()

		delta = time.Since(frameStartTime).Seconds() * targetTicksPerSecond
	}
}
