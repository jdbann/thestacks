package component

import rl "github.com/gen2brain/raylib-go/raylib"

type Goals struct {
	MaxSpeed float32

	Wander               float32
	WanderCircleDistance float32
	WanderCircleRadius   float32
}

type Position rl.Vector2

type Velocity rl.Vector2
