package lib

import "github.com/veandco/go-sdl2/sdl"

//CameraTranslate handles translation of the camera
func CameraTranslate(world *World, state []uint8) {
	//? Camera translation
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
}

//CameraRotate handles rotation of the camera
func CameraRotate(world *World, state []uint8) {
	//? Camera rotation
	if state[sdl.SCANCODE_KP_6] == 1 {
		world.ActiveCamera.Rotate(0, 0, 0.02)
	}
	if state[sdl.SCANCODE_KP_4] == 1 {
		world.ActiveCamera.Rotate(0, 0, -0.02)
	}
	if state[sdl.SCANCODE_KP_8] == 1 {
		world.ActiveCamera.Rotate(0.02, 0, 0)
	}
	if state[sdl.SCANCODE_KP_2] == 1 {
		world.ActiveCamera.Rotate(-0.02, 0, 0)
	}
}

//ObjectRotate handles rotation of a given object
func ObjectRotate(world *World, state []uint8, objID string) {
	//? Cube rotation
	if state[sdl.SCANCODE_PERIOD] == 1 {
		world.Objects[objID].Rotate(0, 0, 0.02)
	}
	if state[sdl.SCANCODE_COMMA] == 1 {
		world.Objects[objID].Rotate(0, 0, -0.02)
	}
}

//MovementHandle Handles all movement
func MovementHandle(world *World, state []uint8) {
	CameraTranslate(world, state)
	CameraRotate(world, state)

	ObjectRotate(world, state, "cube")
}
