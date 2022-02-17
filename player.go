package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func newPlayer(renderer *sdl.Renderer) *element {
	player := &element{}

	player.position = coordinates{
		x: 200,
		y: 400}
	player.rotation = 0

	r := newPlayerRenderer(player, renderer)
	player.addComponent(r)

	mover := newKeyboardMover(player, 4, 0.05)
	player.addComponent(mover)

	player.active = true

	return player
}
