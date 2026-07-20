package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func main() {

	firstPoint := NewPoint(0, 0)
	secondPoint := NewPoint(3, 4)

	// Distanceto bu method Pointga tegishli shuning uchun firstPoin. qilib ishalta olamiz
	distance := firstPoint.DistanceTo(secondPoint)
	fmt.Println("Distance:", distance)
}

// Pointni qaytaradi structe yaratilganidan keyin unga mos obyekt yaratilishi kerak qiymatlarni qnadya qaytarish kerakligini belgilshga ishaltish uchun struct ni

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) X() float64 {
	return p.x
}

func (p Point) Y() float64 {
	return p.y
}

func (p Point) DistanceTo(other Point) float64 {

	// berilgan ikkta nuqtani ayirsak yani moslarini x dan x ni y dan y ni ularning uzunligi kelib chiqadi
	dx := p.x - other.x
	dy := p.y - other.y

	// Katetlar kvadratlari yigindisi gipatenuza kvadratiga tegishli - Pifagor teoremasi faqat asosiy burchagi 90 gradus bolsa ishlaydi
	return math.Sqrt(dx*dx + dy*dy)
}
