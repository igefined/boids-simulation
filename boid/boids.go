package boid

import (
	"math"
	"math/rand"
	"time"

	"github.com/igefined/boids-simulation/vector"
)

type Boid struct {
	id       int
	position vector.Vector2D
	velocity vector.Vector2D
}

func (b *Boid) calcAcceleration() vector.Vector2D {
	upper, lower := b.position.AddV(viewRadius), b.position.AddV(-viewRadius)

	var (
		averageVelocity = vector.Vector2D{}
		averagePosition = vector.Vector2D{}
		separation      = vector.Vector2D{}
		count           float64
	)

	lock.RLock()
	for i := math.Max(lower.X, 0); i <= math.Min(upper.X, screenWidth); i++ {
		for j := math.Max(lower.Y, 0); j <= math.Min(upper.Y, screenHeight); j++ {
			if otherBoidID := boidMap[int(i)][int(j)]; otherBoidID != -1 && otherBoidID != b.id {
				if distance := boids[otherBoidID].position.Distance(b.position); distance < viewRadius {
					count++
					averageVelocity = averageVelocity.Add(boids[otherBoidID].velocity)
					averagePosition = averagePosition.Add(boids[otherBoidID].position)
					separation = separation.Add(b.position.Subtract(boids[otherBoidID].position).DivisionV(distance))
				}
			}
		}
	}
	lock.RUnlock()

	acceleration := vector.Vector2D{
		X: b.borderBounce(b.position.X, screenWidth),
		Y: b.borderBounce(b.position.Y, screenHeight),
	}

	if count > 0 {
		averagePosition, averageVelocity = averagePosition.DivisionV(count), averageVelocity.DivisionV(count)
		accelerationAligment := averageVelocity.Subtract(b.velocity).MultiplyV(adjustmentRate)
		accelerationCohesion := averagePosition.Subtract(b.position).MultiplyV(adjustmentRate)
		accelerationSeparation := separation.MultiplyV(adjustmentRate)
		acceleration = acceleration.Add(accelerationAligment).Add(accelerationCohesion).Add(accelerationSeparation)
	}

	return acceleration
}

func (b *Boid) borderBounce(position, maxBorder float64) float64 {
	if position < viewRadius {
		return 1 / position
	} else if position > maxBorder-viewRadius {
		return 1 / (position - maxBorder)
	}

	return 0
}

func (b *Boid) moveOne() {
	acceleration := b.calcAcceleration()

	lock.Lock()
	defer lock.Unlock()

	b.velocity = b.velocity.Add(acceleration).Limit(-1, 1)

	boidMap[int(b.position.X)][int(b.position.Y)] = -1
	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.X)][int(b.position.Y)] = b.id
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
		position: vector.Vector2D{X: rand.Float64() * screenWidth, Y: rand.Float64() * screenHeight},
		velocity: vector.Vector2D{X: (rand.Float64() * 2) - 1.0, Y: (rand.Float64() * 2) - 1.0},
	}
}
