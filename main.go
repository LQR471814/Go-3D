package main

import (
	"3DEngine/lib"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	//* Init SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	//* Init Window
	window, err := sdl.CreateWindow("3D", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		1280, 720, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	//* Init Renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	renderer.Clear()
	defer renderer.Destroy()

	renderer.SetDrawColor(255, 255, 255, 255)

	//! Run
	lib.Render(renderer, lib.NewDefaultWorld())

	//? Check for quit
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}
