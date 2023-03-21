package main

import (
    "fmt"
    "math"
)

type Point struct {
    x float64
    y float64
}

func NewPoint(x float64, y float64) *Point {
    return &Point{x: x, y: y}
}

func (p *Point) GetX() float64 {
    return p.x
}

func (p *Point) GetY() float64 {
    return p.y
}

func (p *Point) DistanceTo(other *Point) float64 {
    x := p.GetX() - other.GetX()
    y := p.GetY() - other.GetY()

    return math.Sqrt(x * x + y * y)
}

func main() {
    p1 := NewPoint(1, 2)
    p2 := NewPoint(2, 1)
    fmt.Printf("distance: %v\n", p1.DistanceTo(p2))
}
