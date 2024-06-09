package v1

import (
	"fmt"
	"math"

	"v1/point"
)

func computeDistanceBetweenPoints(a, b point.Point) float64 {
	dx := b.GetX() - a.GetY()
	dy := b.GetY() - a.GetY()
	b.GetX() // want `RETC1: return value for function call b.GetX() is ignored`
	return math.Sqrt(float64(computeSquaredDistance(dx, dy)))
}

func computeSquaredDistance(dx, dy int) int {
	return dx*dx + dy*dy
}

func returnsNoValues() {
}

func SpecialSumOfDistances(target point.Point, points []point.Point) float64 {
	var out float64
	target.Unpack()              // want `RETC1: return values for function call target.Unpack() are ignored`
	target.Copy().Copy().SetY(5) // want `RETC1: return value for function call target.Copy().Copy().SetY(...) is ignored`
	target.Copy()                // want `RETC1: return value for function call target.Copy() is ignored`
	switch target.GetX() {
	case target.GetY():
		return 0
	}
	for _, point := range points {
		computeDistanceBetweenPoints(target.Copy(), point) // want `RETC1: return value for function call computeDistanceBetweenPoints(\.\.\.) is ignored`
		out += computeDistanceBetweenPoints(target.Copy(), point)
		_ = computeDistanceBetweenPoints(target.Copy(), point) // this is okay... for now
	}
	return out
}

type Handler struct {
}

type HandlerRequest struct {
	Target point.Point
	Others []point.Point
}

type HandlerResponse struct {
	DistanceSum float64
}

func (me *Handler) emit(msg string) (int, error) {
	return fmt.Println(msg)
}

func (me *Handler) HandleRequest(req HandlerRequest) HandlerResponse {
	go me.emit("In middling of handling request") // don't complain if return value is ignored when spawning Goroutine
	me.emit("Handling request")                   // want `RETC1: return values for function call me.emit(...) are ignored`
	defer me.emit("Handled request")              // don't complain if return value is ignored in `defer` call
	output := SpecialSumOfDistances(req.Target, req.Others)
	output = SpecialSumOfDistances(req.Target, req.Others)
	returnsNoValues()
	_ = output
	return HandlerResponse{
		DistanceSum: output,
	}
}
