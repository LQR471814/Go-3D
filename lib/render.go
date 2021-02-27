package lib

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

//Render renders the world using given SDL renderer
func Render(r *sdl.Renderer, w *World) {
	width := float64(r.GetViewport().W)
	height := float64(r.GetViewport().H)

	renderedVertices := [][2]int{}

	for _, obj := range w.Objects {
		for _, vertex := range obj.Geometry.Vertices {

			renderCoordinates := AsRelativeToSystem(w.ActiveCamera.ObjSys, ToSystem(obj.ObjSys, vertex.Pos))

			ratio := w.ActiveCamera.FocalLength / renderCoordinates.Z
			renderX := ratio * renderCoordinates.X
			renderY := ratio * renderCoordinates.Y

			normX, normY := NormalizeScreen(
				width,
				height,
				renderX,
				renderY,
			)

			rasterX := int(math.Floor(normX*width) + width/2)
			rasterY := int(math.Floor(normY*height) + height/2)

			renderedVertices = append(renderedVertices, [2]int{rasterX, rasterY})

			DrawCircle(r, rasterX, rasterY, 5)
		}
		for _, edge := range obj.Geometry.Edges {
			r.DrawLine(
				int32(renderedVertices[edge.From][0]),
				int32(renderedVertices[edge.From][1]),
				int32(renderedVertices[edge.To][0]),
				int32(renderedVertices[edge.To][1]),
			)
		}
	}

	r.Present()
}
