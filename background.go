package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func newBackground(renderer *sdl.Renderer) *element {
	background := &element{}
	background.position = vector{
		x: screenWidth / 2.0,
		y: screenHeight / 2.0,
	}

	background.active = true

	sr := newSpriteRenderer(background, renderer, "sprites/background.bmp")
	background.addComponent(sr)

	return background
}
