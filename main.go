package main

import (
	"fmt"
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
