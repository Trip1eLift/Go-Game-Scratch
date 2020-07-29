package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *element
	speed     float64
	sr        *spriteRenderer
}

func newKeyboardMover(container *element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()
	posX := mover.container.position.x

	if keys[sdl.SCANCODE_LEFT] == 1 && posX-(mover.sr.width/2.0) > 0 {
		// Move player left
		mover.container.position.x -= mover.speed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 && posX+(mover.sr.width/2.0) < screenWidth {
		// Move player right
		mover.container.position.x += mover.speed
	}

	return nil
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *keyboardMover) onCollision(other *element) error {
	return nil
}

// --------------------------Keyboard Shooter-------------------

type keyboardShooter struct {
	container *element
	cooldown  time.Duration
	lastShot  time.Time
}

func newKeyboardShooter(container *element, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		cooldown:  cooldown,
	}
}

func (mover *keyboardShooter) onUpdate() error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(mover.lastShot) >= mover.cooldown {
			mover.shoot(25, -20)
			mover.shoot(-25, -20)
			mover.lastShot = time.Now()
		}
	}
	return nil
}

func (mover *keyboardShooter) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *keyboardShooter) shoot(x_off, y_off float64) {
	if bul, ok := bulletFromPool(); ok {
		bul.active = true
		bul.position.x = mover.container.position.x + x_off
		bul.position.y = mover.container.position.y + y_off
		bul.rotation = 270 * (math.Pi / 180)
	}
}

func (mover *keyboardShooter) onCollision(other *element) error {
	return nil
}
