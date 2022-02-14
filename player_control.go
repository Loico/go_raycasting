package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *element
	speed     float64

	r *playerRenderer
}

func newKeyboardMover(container *element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		r:         container.getComponent(&playerRenderer{}).(*playerRenderer),
	}
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()

	newPos := mover.container.position

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if keys[sdl.SCANCODE_UP] == 1 {
			newPos.x -= mover.speed * delta * 0.707
			newPos.y -= mover.speed * delta * 0.707
		} else if keys[sdl.SCANCODE_DOWN] == 1 {
			newPos.x -= mover.speed * delta * 0.707
			newPos.y += mover.speed * delta * 0.707
		} else {
			newPos.x -= mover.speed * delta
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if keys[sdl.SCANCODE_UP] == 1 {
			newPos.x += mover.speed * delta * 0.707
			newPos.y -= mover.speed * delta * 0.707
		} else if keys[sdl.SCANCODE_DOWN] == 1 {
			newPos.x += mover.speed * delta * 0.707
			newPos.y += mover.speed * delta * 0.707
		} else {
			newPos.x += mover.speed * delta
		}
	} else if keys[sdl.SCANCODE_UP] == 1 {
		newPos.y -= mover.speed * delta
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		newPos.y += mover.speed * delta
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
