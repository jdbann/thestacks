package system

import (
	"github.com/jdbann/lasagne"
	"github.com/jdbann/thestacks/component"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type Movement struct {
	systemsRes generic.Resource[model.Systems]

	filter *generic.Filter3[component.Position, component.Velocity, lasagne.Object]
}

func (m *Movement) Finalize(w *ecs.World) {}

func (m *Movement) Initialize(w *ecs.World) {
	m.systemsRes = generic.NewResource[model.Systems](w)

	m.filter = generic.NewFilter3[component.Position, component.Velocity, lasagne.Object]().
		Optional(generic.T[lasagne.Object]())
}

func (m *Movement) Update(w *ecs.World) {
	tps := m.systemsRes.Get().TPS
	query := m.filter.Query(w)
	for query.Next() {
		pos, vel, object := query.Get()

		pos.X += vel.X / float32(tps)
		pos.Y += vel.Y / float32(tps)

		if object != nil {
			object.Position.X = pos.X
			object.Position.Y = pos.Y
		}
	}
}

var _ model.System = (*Movement)(nil)
