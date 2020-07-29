package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	container     *element
	tex           *sdl.Texture
	width, height float64
}

func newSpriteRenderer(container *element, renderer *sdl.Renderer, filename string) *spriteRenderer {
	tex := textureFromBMP(renderer, filename)
	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}

	return &spriteRenderer{
		container: container,
		tex:       textureFromBMP(renderer, filename),
		width:     float64(width),
		height:    float64(height),
	}
}

func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	x := sr.container.position.x - sr.width/2.0
	y := sr.container.position.y - sr.height/2.0
	width := int32(sr.width)
	height := int32(sr.height)
	renderer.CopyEx(
		sr.tex,
		&sdl.Rect{X: 0, Y: 0, W: width, H: height},
		&sdl.Rect{X: int32(x), Y: int32(y), W: width, H: height},
		sr.container.rotation,
		&sdl.Point{X: width / 2, Y: height / 2},
		sdl.FLIP_NONE)

	return nil
}

func (sr *spriteRenderer) onUpdate() error {
	return nil
}

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer img.Free()
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from %v: %v", filename, err))
	}

	return tex
}
