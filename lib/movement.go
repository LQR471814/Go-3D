package lib

import (
	"github.com/veandco/go-sdl2/sdl"
)

const movementFactor = 0.05
const rotationFactor = 0.002

//CameraTranslate handles translation of the camera
func CameraTranslate(world *World, state []uint8) {
	//? Camera translation
	if state[sdl.SCANCODE_W] == 1 {
		world.ActiveCamera.Translate(0, 0, -movementFactor)
	}
	if state[sdl.SCANCODE_S] == 1 {
		world.ActiveCamera.Translate(0, 0, movementFactor)
	}
	if state[sdl.SCANCODE_D] == 1 {
		world.ActiveCamera.Translate(-movementFactor, 0, 0)
	}
	if state[sdl.SCANCODE_A] == 1 {
		world.ActiveCamera.Translate(movementFactor, 0, 0)
	}
	if state[sdl.SCANCODE_SPACE] == 1 {
		world.ActiveCamera.Translate(0, movementFactor, 0)
	}
	if state[sdl.SCANCODE_LSHIFT] == 1 {
		world.ActiveCamera.Translate(0, -movementFactor, 0)
	}
}

//MouseHandle handles mouse rotation for a given mouse motion event
func MouseHandle(world *World, e *sdl.MouseMotionEvent) {
	world.ActiveCamera.Rotate(0, rotationFactor*float64(e.XRel), 0)
	world.ActiveCamera.Rotate(rotationFactor*float64(e.YRel), 0, 0)
}

//ObjectRotate handles rotation of a given object
func ObjectRotate(world *World, state []uint8, objID string) {
	//? Cube rotation
	if state[sdl.SCANCODE_PERIOD] == 1 {
		world.Objects[objID].Rotate(0, 2*-rotationFactor, 0)
	}
	if state[sdl.SCANCODE_COMMA] == 1 {
		world.Objects[objID].Rotate(0, 2*rotationFactor, 0)
	}
}

//KeyboardHandle Handles all movement
func KeyboardHandle(world *World, state []uint8) {
	CameraTranslate(world, state)

	ObjectRotate(world, state, "cube")
}
