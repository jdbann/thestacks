package system

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jdbann/lasagne"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/ecs/event"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/arche/listener"
)

type Render struct {
	cameraRes generic.Resource[lasagne.Camera]
	sceneRes  generic.Resource[lasagne.Scene]
}

func (r *Render) FinalizeUI(w *ecs.World) {
	rl.CloseWindow()
}

func (r *Render) InitializeUI(w *ecs.World) {
	r.cameraRes = generic.NewResource[lasagne.Camera](w)
	r.sceneRes = generic.NewResource[lasagne.Scene](w)

	r.setupObjectListener(w)

	rl.InitWindow(1280, 720, "the stacks")
}

func (r *Render) PostUpdateUI(w *ecs.World) {
	if rl.WindowShouldClose() {
		terminate := generic.NewResource[resource.Termination](w)
		terminate.Get().Terminate = true
	}
}

func (r *Render) UpdateUI(w *ecs.World) {
	if !r.cameraRes.Has() || !r.sceneRes.Has() {
		return
	}

	camera := r.cameraRes.Get()
	scene := r.sceneRes.Get()

	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.NewColor(10, 10, 24, 255))
	scene.Draw(camera)
}

func (r *Render) setupObjectListener(w *ecs.World) {
	objectID := ecs.ComponentID[lasagne.Object](w)
	objectMap := generic.NewMap1[lasagne.Object](w)

	callbackFn := func(w *ecs.World, ee ecs.EntityEvent) {
		scene := r.sceneRes.Get()

		if ee.Added.Get(objectID) {
			object := objectMap.Get(ee.Entity)
			scene.AddObject(object)
		}

		if ee.Removed.Get(objectID) {
			object := objectMap.Get(ee.Entity)
			scene.RemoveObject(object)
		}
	}

	objectListener := listener.NewCallback(callbackFn, event.Components, objectID)
	w.SetListener(&objectListener)
}

var _ model.UISystem = (*Render)(nil)
