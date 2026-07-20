package main

import (
	"errors"
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

	// task number 2

	// Point yaratamiz
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)

	circle := CircleNew(p1, 1)

	// Markazi()
	fmt.Println("Markazi:", circle.Center())

	// Radius()
	fmt.Println("Radius:", circle.Radius())

	// Yuzasi()
	fmt.Println("Yuzasi:", circle.Area())

	// Perimeter()
	fmt.Println("Perimeter:", circle.Perimeter())

	// qayerdaligini aniqlash()
	fmt.Println("P2 aylana ichidami:", circle.Contains(p2))

	// marta oshirish()
	err := circle.Scale(2)
	// scalega berib koramiz yani factor sonni agar nil qaytsa yani ifdan otib unday bolsa uni nil emasga tenglaymiz nilga nil teng emas true bolsa error qaytaramiz yani errorni
	if err != nil {
		fmt.Println(err)
		return
	}

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

//constructr nimaga nima tegishliligi

func (p Point) DistanceTo(other Point) float64 {

	// berilgan ikkta nuqtani ayirsak yani moslarini x dan x ni y dan y ni ularning uzunligi kelib chiqadi
	dx := p.x - other.x
	dy := p.y - other.y

	// Katetlar kvadratlari yigindisi gipatenuza kvadratiga tegishli - Pifagor teoremasi faqat asosiy burchagi 90 gradus bolsa ishlaydi
	return math.Sqrt(dx*dx + dy*dy)
}

//task 2

// aylananing structi markaz va radius dan tashkil topgan qiymatlari
type Circle struct {
	center Point
	radius float64
}

// constructr centerga Pointni radiusga float64 son qaytardi
func CircleNew(center Point, radius float64) Circle {
	return Circle{
		center: center,
		radius: radius,
	}
}

// bu aylana markazini qaytaradi yani Point structdagi qiymatni center deb olsak
func (c Circle) Center() Point {
	return c.center
}

// radius qaytaradi yani Circle structdagi radiusni
func (c Circle) Radius() float64 {
	return c.radius
}

// aylananing radiusini topish uchun PI yani 3.14... 17 sondan tashkil topgan ni radius ning kvadratiga kopaytiramiz
func (c Circle) Area() float64 {
	area := math.Pi * c.radius * c.radius
	return area
}

// aylaning uzunligini topish uchun 2pr yani 3.14...ni 2 ga va radius ga kopaytriamiz uzunlik kelib chiqadi
func (c Circle) Perimeter() float64 {
	perimetr := 2 * math.Pi * c.radius
	return perimetr
}

// nuqta aylanma ichidami yoki yoqmi tekshirish buyerda nima uchun DistanceTo buni hiosblaymiz Point bu hisoblay oladi yani markz va nuqatagacha masofani agar radiusdan kicih bolsa demak u ichkarida

// agar markazdan nuqtagacha bolgan masofa radiusdan kichik bolsa demak u aylana ichida malsan markazdan 3sm uzoqlikda bolsa radius 5 sm bolsa demak u  5dan kichik aylana ichida boladi
// agar markazdan nuqtagacha bolgan masofa radiusdan katta bolsa demak u aylana tashqarisida malsan markazdan 7sm uzoqlikda bolsa radius 5 sm bolsa demak u  5dan katta aylana tashqarisida boladi

func (c Circle) Contains(p Point) bool {
	return c.center.DistanceTo(p) <= c.radius
}

// buyerda *Circle qilib ishaltishizning asosiy maqsadi qiymat ozgaryapti tepadagi birorda funclarda qiymat ozgarmaydi shunchaki ihslatyapmiz buyerda esa radiusni faktorga kopaytiryapmiz

// nima uchun kopaytiryamiz topshiriqda aytilganidek marta ozgartirish oshirish ga kopaytirish bolganida oqshar edik xozir marta oshirish marta sozi kopaytirishni anglatadi
func (c *Circle) Scale(factor float64) error {

	if factor <= 0 {
		return errors.New("Faktor soni 0ga teng yoki undan kichik")
	}

	c.radius *= factor

	return nil
}
