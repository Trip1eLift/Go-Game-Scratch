package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	enemyBulletSpeed = 10
)

func newEnemyBullet(renderer *sdl.Renderer) *element {
	bullet := &element{}

	sr := newSpriteRenderer(bullet, renderer, "sprites/enemy_bullet.bmp")
	bullet.addComponent(sr)

	mover := newBulletMover(bullet, enemyBulletSpeed)
	bullet.addComponent(mover)

	bullet.active = false
	bullet.tag = "enemyBullet"

	col := circle{
		center: bullet.position,
		radius: 8,
	}
	bullet.collisions = append(bullet.collisions, col)

	return bullet
}

var enemyBulletPool []*element

func initEnemyBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bul := newBullet(renderer)
		elements = append(elements, bul)
		enemyBulletPool = append(enemyBulletPool, bul)
	}
}

func enemyBulletFromPool() (*element, bool) {
	for _, bul := range enemyBulletPool {
		if !bul.active {
			return bul, true
		}
	}

	return nil, false
}
