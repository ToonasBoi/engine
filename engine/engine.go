package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Config(winW int32, winH int32, resW int32, resH int32, title string) {
	rl.InitWindow(winW, winH, title)
}
