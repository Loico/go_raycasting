package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

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

func (mover *keyboardMover) onCollision(other *element) error {
	return nil
}
