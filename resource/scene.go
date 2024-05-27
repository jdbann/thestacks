package resource

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jdbann/lasagne"
)

func DefaultScene() *lasagne.Scene {
	tileSet := lasagne.NewTileSet(16)
	bridgeTile := tileSet.AddTile(rl.LoadTexture("assets/Bridge_strip16.png"))
	floorTile := tileSet.AddTile(rl.LoadTexture("assets/FloorCentrePlain_strip16.png"))
	floorPatternTile := tileSet.AddTile(rl.LoadTexture("assets/FloorCentrePattern_strip16.png"))
	wallTile := tileSet.AddTile(rl.LoadTexture("assets/WallCentreA_strip16.png"))

	tileMap := lasagne.NewTileMap([][][]int{
		{
			{floorTile, floorTile, floorTile, -1},
			{floorTile, floorTile, floorTile, bridgeTile},
			{floorPatternTile, floorTile, floorTile, -1},
			{floorPatternTile, floorPatternTile, floorTile, -1},
		},
		{
			{-1, -1, -1, -1},
			{-1, wallTile, -1, -1},
			{-1, -1, -1, -1},
			{-1, -1, -1, -1},
		},
	})

	return lasagne.NewScene(lasagne.SceneParams{
		Objects: []lasagne.Object{},
		TileMap: tileMap,
		TileSet: tileSet,
	})
}
