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
	pos    coordinates
	end    coordinates
	dir    float64
	length float64
}

// cast checks if a ray intersects a wall
// If the ray intersects the wall, it returns true and the intersection point coordinates
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
	}
	return false, pt
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
	var step float64 = float64(fov) / float64(nbRay)
	for a = -(fov / 2); a < fov/2; a += step {
		r.rays = append(r.rays, ray{r.container.position, r.container.position, a * math.Pi / 180, 0})
	}

	return r
}

// onUpdate updates rays struct according to player position
func (r *playerRenderer) onUpdate() error {
	for i := range r.rays {
		r.rays[i].pos.x = r.container.position.x
		r.rays[i].pos.y = r.container.position.y

		var d float64 = 0
		var minD float64 = mapHeight * mapWidth
		var record coordinates
		for _, w := range walls {
			ret, pt := cast(w, r.rays[i], r.container.rotation)
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
		l := math.Sqrt((r.rays[i].pos.x-r.rays[i].end.x)*(r.rays[i].pos.x-r.rays[i].end.x) + (r.rays[i].pos.y-r.rays[i].end.y)*(r.rays[i].pos.y-r.rays[i].end.y))
		r.rays[i].length = math.Cos(r.rays[i].dir) * l
	}
	return nil
}

// onDraw draws player's rays on 2D and 3D view
func (r *playerRenderer) onDraw(renderer *sdl.Renderer) error {
	renderer.SetDrawColor(255, 255, 255, 255)
	drawCircle(renderer, int32(r.container.position.x), int32(r.container.position.y), 8)

	for i, ray := range r.rays {
        // Draw 2D top down view
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.DrawLine(int32(ray.pos.x), int32(ray.pos.y), int32(ray.end.x), int32(ray.end.y))

        // Draw 3D first person view
		var rect sdl.Rect
		rect.X = int32((screenWidth / 2) + i*(screenWidth/2)/nbRay)
		rect.Y = int32((screenHeight / 2) - (wallHeightConst/ray.length)/2)
		rect.W = int32((screenWidth / 2) / nbRay)
		rect.H = int32(wallHeightConst / ray.length)
		brightness := brightnessConst / ray.length
		if brightness < 0 {
			brightness = 0
		} else if brightness > 255 {
			brightness = 255
		}
		renderer.SetDrawColor(uint8(brightness), uint8(brightness), uint8(brightness), 255)
		renderer.FillRect(&rect)
	}
	return nil
}
