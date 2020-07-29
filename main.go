package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800

	targetTicksPerSecond = 60
	frameRate            = 60.0
)

var delta float64

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SEL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Gaming in Go Episode 2",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)

	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	elements = append(elements, newPlayer(renderer))

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*screenWidth + (basicEnemySize / 2.0)
			y := float64(j)*basicEnemySize + (basicEnemySize / 2.0)

			enemy := newBasicEnemy(renderer, vector{x, y})
			elements = append(elements, enemy)
		}
	}

	initBulletPool(renderer)

	for {
		frameStartTime := time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		for _, ele := range elements {
			if ele.active {
				err = ele.update()
				if err != nil {
					fmt.Println("updating element:", err)
					return
				}
				err = ele.draw(renderer)
				if err != nil {
					fmt.Println("drawing element:", err)
					return
				}
			}
		}

		if err := checkCollisions(); err != nil {
			fmt.Println("checking collisions:", err)
			return
		}

		renderer.Present()

		frameTime := 1.0 / frameRate
		for time.Since(frameStartTime).Seconds() < frameTime {
			time.Sleep(time.Duration(frameTime/4) * time.Second)
		}
		delta = time.Since(frameStartTime).Seconds() * targetTicksPerSecond
		//fmt.Println(1 / time.Since(frameStartTime).Seconds())
	}
}
