package system

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jdbann/thestacks/lasagne"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type Render struct {
	cameraRes generic.Resource[lasagne.Camera]
	sceneRes  generic.Resource[lasagne.Scene]
}

func (r *Render) FinalizeUI(w *ecs.World) {
	rl.CloseWindow()
}

func (r *Render) InitializeUI(w *ecs.World) {
	rl.InitWindow(1280, 720, "the stacks")
}

func (r *Render) PostUpdateUI(w *ecs.World) {
	if rl.WindowShouldClose() {
		terminate := generic.NewResource[resource.Termination](w)
		terminate.Get().Terminate = true
	}
}

func (r *Render) UpdateUI(w *ecs.World) {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.NewColor(10, 10, 24, 255))

	camera := generic.NewResource[lasagne.Camera](w)
	scene := generic.NewResource[lasagne.Scene](w)
	if camera.Has() && scene.Has() {
		scene.Get().Draw(*camera.Get())
	}
}

var _ model.UISystem = (*Render)(nil)
