package main

import "math"

type circle struct {
	center vector
	radius float64
}

func collides(c1, c2 circle) bool {
	dx := c2.center.x - c1.center.x
	dy := c2.center.y - c1.center.y
	dist := math.Sqrt(dx*dx + dy*dy)

	return dist <= (c1.radius + c2.radius)
}

func checkCollisions() error {
	for i := 0; i < len(elements)-1; i++ {
		for j := i + 1; j < len(elements); j++ {
			for _, c1 := range elements[i].collisions {
				for _, c2 := range elements[j].collisions {
					if elements[i].active && elements[j].active && collides(c1, c2) {
						err := elements[i].collision(elements[j])
						if err != nil {
							return err
						}
						err = elements[j].collision(elements[i])
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}
