package main

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type vector struct {
	x, y float64
}

type component interface {
	onUpdate() error
	onDraw(render *sdl.Renderer) error
	onCollision(other *element) error
}

type element struct {
	position   vector
	rotation   float64
	active     bool
	tag        string
	collisions []circle
	components []component
}

func (ele *element) draw(renderer *sdl.Renderer) error {
	for _, comp := range ele.components {
		err := comp.onDraw(renderer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ele *element) update() error {
	for _, comp := range ele.components {
		err := comp.onUpdate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (ele *element) collision(other *element) error {
	for _, comp := range ele.components {
		err := comp.onCollision(other)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ele *element) addComponent(new component) {
	for _, existing := range ele.components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf("attempt to add new component with existing type %v",
				reflect.TypeOf(new)))
		}
	}
	ele.components = append(ele.components, new)
}

func (ele *element) getComponent(withType component) component {
	typ := reflect.TypeOf(withType)
	for _, comp := range ele.components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}

	panic(fmt.Sprintf("no component with %v", reflect.TypeOf(withType)))
}

var elements []*element
