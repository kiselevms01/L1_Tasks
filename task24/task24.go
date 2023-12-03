package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

// конструктор для структуры Point
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// Функция определение расстояния между двумя точками
func (point1 *Point) DistanceBetween(point2 *Point) float64 {
	dx := point1.x - point2.x
	dy := point1.y - point2.y

	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	var pointOne = NewPoint(0.0, 10.0)
	var pointTwo = NewPoint(10.0, 0.0)

	fmt.Printf("Расстояние между точками = %v", pointOne.DistanceBetween(pointTwo))
}
