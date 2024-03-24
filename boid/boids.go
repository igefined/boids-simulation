package boid

import (
	"math"
	"math/rand"
	"time"

	"github.com/igefined/boids-simulation/vector"
)

type Boid struct {
	id       int
	position *vector.Vector2D
	velocity *vector.Vector2D
}

func (b *Boid) calcAcceleration() *vector.Vector2D {
	upper, lower := b.position.AddV(viewRadius), b.position.AddV(-viewRadius)
	averagevelocity := &vector.Vector2D{X: 0, Y: 0}
	var count float64

	for i := math.Max(lower.X, 0); i <= math.Min(upper.X, screenWidth); i++ {
		for j := math.Max(lower.Y, 0); j < math.Min(upper.Y, screenHeight); j++ {
			if otherBoid := boidMap[int(i)][int(j)]; otherBoid != -1 && otherBoid != b.id {
				if distance := boids[otherBoid].position.Distance(b.position); distance < viewRadius {
					count++
					averagevelocity = averagevelocity.Add(boids[otherBoid].velocity)
				}
			}
		}
	}

	acceleration := &vector.Vector2D{}
	if count > 0 {
		averagevelocity = averagevelocity.DivisionV(count)
		acceleration = averagevelocity.Subtract(b.velocity).MultiplyV(adjustmentRate)
	}

	return acceleration
}

func (b *Boid) moveOne() {
	b.velocity = b.velocity.Add(b.calcAcceleration()).Limit(-1, 1)

	boidMap[int(b.position.X)][int(b.position.Y)] = -1
	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.X)][int(b.position.Y)] = b.id

	next := b.position.Add(b.velocity)

	if next.X > screenWidth || next.X < 0 {
		b.velocity = &vector.Vector2D{X: -b.velocity.X, Y: b.velocity.Y}
	}

	if next.Y >= screenHeight || next.Y < 0 {
		b.velocity = &vector.Vector2D{X: b.velocity.X, Y: -b.velocity.Y}
	}
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(time.Millisecond * 5)
	}
}

func createBoid(id int) *Boid {
	return &Boid{
		id:       id,
		position: &vector.Vector2D{X: rand.Float64() * screenWidth, Y: rand.Float64() * screenHeight},
		velocity: &vector.Vector2D{X: (rand.Float64() * 2) - 1.0, Y: (rand.Float64() * 2) - 1.0},
	}
}
