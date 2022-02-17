package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

func intersect(a coordinates, b coordinates, c coordinates, d coordinates) (ret bool, pt coordinates) {
	x1 := a.x
	y1 := a.y
	x2 := b.x
	y2 := b.y
	x3 := c.x
	y3 := c.y
	x4 := d.x
	y4 := d.y

	den := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if den == 0 {
		return false, pt
	}
	var t, u float64
	t = ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) / den
	u = (-((x1-x2)*(y1-y3) - (y1-y2)*(x1-x3))) / den
	if t > 0 && t < 1 && u > 0 && u < 1 {
		pt.x = x1 + t*(x2-x1)
		pt.y = y1 + t*(y2-y1)
		return true, pt
	} else {
		return false, pt
	}
}

func angleFromABC(a coordinates, b coordinates, c coordinates) float64 {
	ab := math.Sqrt((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y))
	bc := math.Sqrt((b.x-c.x)*(b.x-c.x) + (b.y-c.y)*(b.y-c.y))
	ac := math.Sqrt((a.x-c.x)*(a.x-c.x) + (a.y-c.y)*(a.y-c.y))
	return math.Acos((ab*ab + bc*bc - ac*ac) / (2 * ab * bc))
}

type keyboardMover struct {
	container *element
	speed     float64
	rotSpeed  float64

	r *playerRenderer
}

func newKeyboardMover(container *element, speed float64, rotSpeed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		rotSpeed:  rotSpeed,
		r:         container.getComponent(&playerRenderer{}).(*playerRenderer),
	}
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()

	newPos := mover.container.position

	if keys[sdl.SCANCODE_Q] == 1 || keys[sdl.SCANCODE_LEFT] == 1 {
		mover.container.rotation -= mover.rotSpeed * delta
	}

	if keys[sdl.SCANCODE_E] == 1 || keys[sdl.SCANCODE_RIGHT] == 1 {
		mover.container.rotation += mover.rotSpeed * delta
	}

	moveAngle := mover.container.rotation
	if keys[sdl.SCANCODE_W] == 1 || keys[sdl.SCANCODE_S] == 1 || keys[sdl.SCANCODE_A] == 1 || keys[sdl.SCANCODE_D] == 1 {

		if keys[sdl.SCANCODE_A] == 1 {
			if keys[sdl.SCANCODE_W] == 1 {
				moveAngle -= math.Pi / 4
			} else if keys[sdl.SCANCODE_S] == 1 {
				moveAngle -= 3 * math.Pi / 4
			} else {
				moveAngle -= math.Pi / 2
			}
		} else if keys[sdl.SCANCODE_D] == 1 {
			if keys[sdl.SCANCODE_W] == 1 {
				moveAngle += math.Pi / 4
			} else if keys[sdl.SCANCODE_S] == 1 {
				moveAngle += 3 * math.Pi / 4
			} else {
				moveAngle += math.Pi / 2
			}
		} else if keys[sdl.SCANCODE_S] == 1 {
			moveAngle += math.Pi
		}

		newPos.x += mover.speed * delta * math.Cos(moveAngle)
		newPos.y += mover.speed * delta * math.Sin(moveAngle)
	}

	// Check collisions
	for _, w := range walls {
		ret, pt := intersect(w.a, w.b, mover.container.position, newPos)
		if ret {
			// TODO: Add sliding collisions
			// Calculate angle
			//angle := angleFromABC(mover.container.position, pt, w.a)
			_ = pt
			newPos = mover.container.position
		}

	}
	if int(newPos.y)-(mover.r.height/2.0) < 0 {
		newPos.y = float64(mover.r.height) / 2.0
	}
	if int(newPos.y)+(mover.r.height/2.0) > screenHeight {
		newPos.y = screenHeight - (float64(mover.r.height) / 2.0)
	}
	if int(newPos.x)-(mover.r.width/2.0) < 0 {
		newPos.x = float64(mover.r.width) / 2.0
	}
	if int(newPos.x)+(mover.r.width/2.0) > screenWidth {
		newPos.x = screenWidth - (float64(mover.r.width) / 2.0)
	}

	mover.container.position = newPos

	return nil
}
