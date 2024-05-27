package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jdbann/lasagne"
	"github.com/jdbann/thestacks/resource"
	"github.com/jdbann/thestacks/system"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

func Run() {
	// Setup a new model
	m := model.New()

	m.TPS = 60 // Limit ticks per second of simulation
	m.FPS = 60 // Limit frames per second of UI

	// Add systems
	m.AddUISystem(&system.Camera{})
	m.AddUISystem(&system.Render{})

	// Prepare the model and systems
	m.Initialize()

	// Add resources not configured provided by systems
	ecs.AddResource(&m.World, resource.DefaultKeyBindings())
	ecs.AddResource(&m.World, resource.DefaultScene())

	// Add entities
	objectBuilder := generic.NewMap1[lasagne.Object](&m.World)

	barrelTexture := rl.LoadTexture("assets/Barrel_strip8.png")
	objectBuilder.NewWith(&lasagne.Object{
		Position: rl.NewVector3(2, 2, 1),
		Texture:  barrelTexture,
		Size:     rl.NewVector3(14, 14, 8),
	})

	chairTexture := rl.LoadTexture("assets/Chair_strip12.png")
	objectBuilder.NewWith(&lasagne.Object{
		Position: rl.NewVector3(1, 3, 1),
		Texture:  chairTexture,
		Size:     rl.NewVector3(12, 12, 12),
	})

	// Run the model
	m.Run()
}
