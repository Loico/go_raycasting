package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

func drawCircle(renderer *sdl.Renderer, centreX int32, centreY int32, radius int32) {
	diameter := radius * 2
	x := radius - 1
	var y int32 = 0
	var tx int32 = 1
	var ty int32 = 1
	err := tx - diameter

	for x >= y {
		renderer.DrawPoint(centreX+x, centreY-y)
		renderer.DrawPoint(centreX+x, centreY+y)
		renderer.DrawPoint(centreX-x, centreY-y)
		renderer.DrawPoint(centreX-x, centreY+y)
		renderer.DrawPoint(centreX+y, centreY-x)
		renderer.DrawPoint(centreX+y, centreY+x)
		renderer.DrawPoint(centreX-y, centreY-x)
		renderer.DrawPoint(centreX-y, centreY+x)
		if err <= 0 {
			y++
			err += ty
			ty += 2
		}
		if err > 0 {
			x--
			tx += 2
			err += tx - diameter
		}
	}
}

type playerRenderer struct {
	container     *element
	width, height int
	rays          []ray
}

func newPlayerRenderer(container *element, renderer *sdl.Renderer) *playerRenderer {
	r := &playerRenderer{}

	r.width, r.height = 16, 16
	r.container = container
	var a float64
	for a = 0; a < 360; a += 0.5 {
		r.rays = append(r.rays, ray{r.container.position, a * math.Pi / 180})
	}

	return r
}

func (r *playerRenderer) onUpdate() error {
	for i := range r.rays {
		r.rays[i].pos.x = r.container.position.x
		r.rays[i].pos.y = r.container.position.y
	}
	return nil
}

func (r *playerRenderer) onDraw(renderer *sdl.Renderer) error {
	renderer.SetDrawColor(255, 255, 255, 255)
	drawCircle(renderer, int32(r.container.position.x), int32(r.container.position.y), 8)

	for _, ray := range r.rays {
		var d float64 = 0
		var minD float64 = screenHeight * screenWidth
		var wallIsFound bool = false
		var record coordinates
		for _, b := range walls {
			ret, pt := cast(b, ray)
			if ret {
				wallIsFound = true
				d = (pt.x-ray.pos.x)*math.Cos(ray.dir) + (pt.y-ray.pos.y)*math.Sin(ray.dir)
				if d < minD {
					minD = d
					record = pt
				}
			}
		}
		if wallIsFound {
			renderer.DrawLine(int32(ray.pos.x), int32(ray.pos.y), int32(record.x), int32(record.y))
		}
	}
	return nil
}

func (r *playerRenderer) onCollision(other *element) error {
	return nil
}

func newPlayer(renderer *sdl.Renderer) *element {
	player := &element{}

	player.position = coordinates{
		x: 200,
		y: 400}

	r := newPlayerRenderer(player, renderer)
	player.addComponent(r)

	mover := newKeyboardMover(player, 5)
	player.addComponent(mover)

	player.active = true

	return player
}
