package system

import (
	"golang.org/x/exp/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jdbann/thestacks/component"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type Steer struct {
	randRes generic.Resource[resource.Rand]

	filter *generic.Filter3[component.Goals, component.Position, component.Velocity]

	rng *rand.Rand
}

func (s *Steer) Finalize(w *ecs.World) {}

func (s *Steer) Initialize(w *ecs.World) {
	s.randRes = generic.NewResource[resource.Rand](w)

	s.filter = generic.NewFilter3[component.Goals, component.Position, component.Velocity]()
}

func (s *Steer) Update(w *ecs.World) {
	s.rng = rand.New(s.randRes.Get())

	query := s.filter.Query(w)
	for query.Next() {
		goals, pos, vel := query.Get()
		steer := rl.Vector2Scale(s.wander(goals, rl.Vector2(*pos), rl.Vector2(*vel)), goals.Wander)
		vel.X += steer.X
		vel.Y += steer.Y
	}
}

func (s *Steer) wander(goals *component.Goals, pos, vel rl.Vector2) rl.Vector2 {
	theta := s.rng.Float32() * rl.Pi * 2

	target := rl.Vector2Add(
		rl.Vector2Add(pos, rl.Vector2Scale(rl.Vector2Normalize(vel), goals.WanderCircleDistance)), // Centre of wander circle
		rl.Vector2Rotate(rl.NewVector2(0, goals.WanderCircleRadius), theta),                       // Random point on wander circle
	)

	// Head towards target at max speed
	seek := rl.Vector2Scale(rl.Vector2Normalize(rl.Vector2Subtract(target, pos)), goals.MaxSpeed)

	return rl.Vector2Subtract(seek, vel)
}

var _ model.System = (*Steer)(nil)
