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

	//! Run
	world := lib.NewDefaultWorld()

	//? Check for quit
	running := true
	for running {
		renderer.SetDrawColor(255, 255, 255, 255)

		lib.Render(renderer, world)
		state := sdl.GetKeyboardState()
		if state[sdl.SCANCODE_W] == 1 {
			world.ActiveCamera.Translate(0, 0, 0.1)
		}
		if state[sdl.SCANCODE_S] == 1 {
			world.ActiveCamera.Translate(0, 0, -0.1)
		}
		if state[sdl.SCANCODE_D] == 1 {
			world.ActiveCamera.Translate(0.1, 0, 0)
		}
		if state[sdl.SCANCODE_A] == 1 {
			world.ActiveCamera.Translate(-0.1, 0, 0)
		}
		if state[sdl.SCANCODE_SPACE] == 1 {
			world.ActiveCamera.Translate(0, -0.1, 0)
		}
		if state[sdl.SCANCODE_LSHIFT] == 1 {
			world.ActiveCamera.Translate(0, 0.1, 0)
		}

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
		sdl.Delay(10)
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
	}
}
