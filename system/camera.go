package system

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jdbann/thestacks/lasagne"
	"github.com/jdbann/thestacks/resource"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type Camera struct {
	cameraRes generic.Resource[lasagne.Camera]
	keysRes   generic.Resource[resource.KeyBindings]
	sceneRes  generic.Resource[lasagne.Scene]
}

func (c *Camera) Finalize(w *ecs.World) {}

func (c *Camera) Initialize(w *ecs.World) {
	c.cameraRes = generic.NewResource[lasagne.Camera](w)
	c.keysRes = generic.NewResource[resource.KeyBindings](w)
	c.sceneRes = generic.NewResource[lasagne.Scene](w)

	camera := lasagne.NewCamera()
	c.reset(camera)
	c.cameraRes.Add(camera)
}

func (c *Camera) Update(w *ecs.World) {
	if !c.cameraRes.Has() || !c.keysRes.Has() || !c.sceneRes.Has() {
		return
	}

	camera := c.cameraRes.Get()
	keys := c.keysRes.Get()
	mouseDelta := rl.GetMouseDelta()

	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		c.rotate(camera, mouseDelta)
	} else if keys.PanCamera.IsDown() {
		c.pan(camera, mouseDelta)
	} else if keys.ResetCamera.IsPressed() {
		c.reset(camera)
	} else {
		c.zoom(camera, rl.GetMouseWheelMove())
	}
}

func (c *Camera) pan(camera *lasagne.Camera, v rl.Vector2) {
	c.sceneRes.Get().MoveCamera(camera, rl.Vector3{X: v.X, Y: v.Y})
}

func (c *Camera) reset(camera *lasagne.Camera) {
	*camera = lasagne.Camera{
		Rotation: rl.Vector2{X: rl.Pi / 4, Y: rl.Pi / 3},
		Target:   rl.Vector3{X: 2, Y: 2, Z: 1},
		Zoom:     2,
	}
}

func (c *Camera) rotate(camera *lasagne.Camera, v rl.Vector2) {
	camera.Rotation.X = rl.Wrap(
		camera.Rotation.X-(v.X/float32(rl.GetScreenWidth()))*rl.Pi*2,
		0, rl.Pi*2,
	)

	camera.Rotation.Y = rl.Clamp(
		camera.Rotation.Y-(v.Y/float32(rl.GetScreenHeight()))*rl.Pi/2,
		0, rl.Pi/2-0.005,
	)
}

func (c *Camera) zoom(camera *lasagne.Camera, v float32) {
	camera.Zoom = rl.Clamp(camera.Zoom+v, 0.5, 8)
}

var _ model.System = (*Camera)(nil)
