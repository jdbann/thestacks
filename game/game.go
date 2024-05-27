package game

import (
	"github.com/jdbann/thestacks/resource"
	"github.com/jdbann/thestacks/system"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

func Run() {
	// Setup a new model
	m := model.New()

	m.TPS = 60 // Limit ticks per second of simulation
	m.FPS = 60 // Limit frames per second of UI

	// Add systems
	m.AddSystem(&system.Camera{})
	m.AddUISystem(&system.Render{})

	// Prepare the model and systems
	m.Initialize()

	// Add resources not configured provided by systems
	ecs.AddResource(&m.World, resource.DefaultKeyBindings())
	ecs.AddResource(&m.World, resource.DefaultScene())

	// Run the model
	m.Run()
}
