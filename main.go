package main

import (
	"go3d/lib"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	//* Init SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	sdl.SetRelativeMouseMode(true)

	//* Init Window
	window, err := sdl.CreateWindow("3D", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		1280, 720, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()

	//* Init Renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatal(err)
	}

	renderer.Clear()
	defer renderer.Destroy()

	//* Run
	world := lib.NewDefaultWorld()

	//? Check for quit
	running := true
	for running {
		renderer.SetDrawColor(255, 255, 255, 255)

		lib.Render(renderer, world)

		state := sdl.GetKeyboardState()

		lib.KeyboardHandle(world, state)

		//? On Exit
		if state[sdl.SCANCODE_ESCAPE] == 1 {
			println("Quitting...")
			running = false
			break
		}

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				println("Quitting...")
				running = false
			case *sdl.MouseMotionEvent:
				lib.MouseHandle(world, e)
			}
		}
		sdl.Delay(5)
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
	}
}
