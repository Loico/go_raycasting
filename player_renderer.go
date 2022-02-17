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

type ray struct {
	pos coordinates
	end coordinates
	dir float64
}

func cast(b wall, r ray, rotation float64) (ret bool, pt coordinates) {
	x1 := b.a.x
	y1 := b.a.y
	x2 := b.b.x
	y2 := b.b.y
	x3 := r.pos.x
	y3 := r.pos.y
	x4 := r.pos.x + math.Cos(r.dir+rotation)
	y4 := r.pos.y + math.Sin(r.dir+rotation)

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
	for a = -(fov / 2); a < fov/2; a += 0.5 {
		r.rays = append(r.rays, ray{r.container.position, r.container.position, a * math.Pi / 180})
	}

	return r
}

func (r *playerRenderer) onUpdate() error {
	for i := range r.rays {
		r.rays[i].pos.x = r.container.position.x
		r.rays[i].pos.y = r.container.position.y

		var d float64 = 0
		var minD float64 = screenHeight * screenWidth
		var record coordinates
		for _, b := range walls {
			ret, pt := cast(b, r.rays[i], r.container.rotation)
			if ret {
				d = math.Abs((pt.x-r.rays[i].pos.x)*math.Cos(r.rays[i].dir) + (pt.y-r.rays[i].pos.y)*math.Sin(r.rays[i].dir))
				if d < minD {
					minD = d
					record = pt
				}
			}
		}
		r.rays[i].end.x = record.x
		r.rays[i].end.y = record.y
	}
	return nil
}

func (r *playerRenderer) onDraw(renderer *sdl.Renderer) error {
	renderer.SetDrawColor(255, 255, 255, 255)
	drawCircle(renderer, int32(r.container.position.x), int32(r.container.position.y), 8)

	for _, ray := range r.rays {
		renderer.DrawLine(int32(ray.pos.x), int32(ray.pos.y), int32(ray.end.x), int32(ray.end.y))
	}
	return nil
}

func (r *playerRenderer) onCollision(other *element) error {
	return nil
}
