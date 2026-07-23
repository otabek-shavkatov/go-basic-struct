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

// part 2 No 5
type Triangle struct {
	a Point
	b Point
	c Point
}

// interface ning vazifasi u oz ichiga methiodlarni olib turadi maslaan boshqa go fileda shaope ichiga olgan func larni shapoe orqali chaqirib ishlata olamiz
// yani shapega ichidagi methodlar shape.Area() qilib ishaltsa boaldi shunday
type Shape interface {
	Area() float64
	Perimeter() float64
	// String() string chunki memnda string ishlatilmagan
}

type Drawable interface {
	Draw() string
}

// Containable
type Containable interface {
	Contains(p Point) bool
}

// aylananing structi markaz va radius dan tashkil topgan qiymatlari
type Circle struct {
	center Point
	radius float64
}
type Rectangle struct {
	topLeft Point
	height  float64
	width   float64
}

type Color struct {
	r uint8
	b uint8
	g uint8
}

// aylana bor edi uning rangli varaintini qyatirsh uchun
type ColoredCircle struct {
	Circle
	color Color
}

type ColoredRectangle struct {
	Rectangle
	color Color
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

	// nuqtalarni beramiz corduinatlarini va ular orqali nuqtalar topiladi va orasidagi masofalar topilib qoshilib perimetr hisoblanadi
	triangle := Triangle{
		a: NewPoint(0, 0),
		b: NewPoint(3, 0),
		c: NewPoint(0, 4),
	}

	fmt.Println("Uchburchak yuzasi", triangle.Area())
	fmt.Println("Uchburchak turlari:", triangle.Type())
	fmt.Println("Perimeter:", triangle.Perimeter())
	fmt.Println("Is valid:", triangle.IsValid())

}

// Pointni qaytaradi structe yaratilganidan keyin unga mos obyekt yaratilishi kerak qiymatlarni qnadya qaytarish kerakligini belgilshga ishaltish uchun struct ni

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) X() float64 {
	return p.x
}

// bu Drawable interface ga tegishli chunki unga draw methiod belguilangan chunki drw method bor
func (c Circle) Draw() string {
	return "   ***\n *   *\n   ***"
}

// Drawable interface
func (r Rectangle) Draw() string {
	return "*****\n*   *\n*****"
}

func (p Point) Y() float64 {
	return p.y
}

func (p Point) String() string {
	return fmt.Sprintf("(%.1f, %.1f)", p.x, p.y)
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

// stringda qaytaradi qiymatni
func (c Circle) String() string {
	return fmt.Sprintf("Circle radius=%.2f", c.radius)
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

//task 3

func ConstructrRectangle(topLeft Point, width, height float64) (Rectangle, error) {
	if width <= 0 || height <= 0 {
		return Rectangle{}, errors.New("width and height ")
	}

	return Rectangle{
		topLeft: topLeft,
		width:   width,
		height:  height,
	}, nil
}

// tortburchakning yuzasini hisoblash
func (r Rectangle) Area() float64 {
	return r.height * r.width
}

// tortburchakning perimetrini hisoblash
func (r Rectangle) Perimeter() float64 {
	return r.height*2 + r.width*2
}

func (r Rectangle) Contains(p Point) bool {
	return p.x >= r.topLeft.x &&
		p.x <= r.topLeft.x+r.width &&
		p.y >= r.topLeft.y &&
		p.y <= r.topLeft.y+r.height
}

// string qaytarish uchun
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle %.1fx%.1f", r.width, r.height)
}
func (r Rectangle) Diagonal() float64 {
	return math.Sqrt(r.width*r.width + r.height*r.height)
}

func (r *Rectangle) Scale(factor float64) error {

	if factor <= 0 {
		return errors.New("ffaktor 0 danm katta bolishgi kerak oke")
	}

	r.width *= factor
	r.height *= factor

	return nil
}

// part 2 started No 4

// rgb qaytaradi

func (c Color) ColorHex() string {
	return fmt.Sprintf("#%02X%02X%02X", c.r, c.g, c.b)
}

// Point orqali nuqtalarni topib olamiz

// a dan b gacha b dan c gacha va a dan c gacha masofalar topilib perimetrlari hisoblandi
func (t Triangle) Perimeter() float64 {
	ab := t.a.DistanceTo(t.b)
	bc := t.b.DistanceTo(t.c)
	ac := t.a.DistanceTo(t.c)

	return ab + bc + ac
}

// uchburchakning yuzasini hiosblash uchun method har bir nuqtaning orasdiagi masofalarni topamiz DistanceTo orqali
func (t Triangle) Area() float64 {

	a := t.a.DistanceTo(t.b)
	b := t.b.DistanceTo(t.c)
	c := t.a.DistanceTo(t.c)

	s := (a + b + c) / 2

	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

// bir burchakga yopshgan ikk tomon yigindisi burchak qarshisidagi tomondan katta bolsa shart qanoatlantiriladi
func (t Triangle) IsValid() bool {
	ab := t.a.DistanceTo(t.b)
	bc := t.b.DistanceTo(t.c)
	ac := t.a.DistanceTo(t.c)

	return ab+bc > ac &&
		ab+ac > bc &&
		ac+bc > ab
}

// uchburchakning turini aniqlash uchun method
func (t Triangle) Type() string {

	a := t.a.DistanceTo(t.b)
	b := t.b.DistanceTo(t.c)
	c := t.a.DistanceTo(t.c)

	if a == b && b == c {
		return "teng yonli uchburchak, yani 90 gradus boladi bir ichki burchagi"
	}

	if a == b || b == c || a == c {
		return "teng tomonli yani ichki burchaklari 60 gradusdan boladi"
	}

	return "har hil tomonli uchburchak"
}

// buyterda Move bolayotgani uchun mos kordinatalari qoshiladi va asl qiymatni ozgarititishimzi mumkin bolishini taminlash uchun * ishlatamiz
func (c *Circle) Move(dx, dy float64) {
	c.center.x += dx
	c.center.y += dy
}

// istalgan qiymatni shunchaki bu qyataradi yani qiymat ozgarmaydi bunday xolatsda if orqali tekshirsak boladi ...
func (c Circle) Describe() string {
	return fmt.Sprintf("Radius: %.2f", c.radius)
}

// buyerda bunday biolyapti shapes aytyaopti men faqat draw ni chaqira olaman deyapti Drawable ichida Draw bor
// Draw esa Circle va Rectangle larda bor
func Render(shapes []Drawable) {
	for _, shape := range shapes {
		fmt.Println(shape.Draw())
	}
}

// Polymorphism - bu bitta qilib aytganda yuza qisimda qulay bolgan xolda bitta narsa bilan ishalsh

// masalan men shape.Draw() deb ishaltyapaman buning tagida Draw ishlatilgan methodlar ... chiqadi, yani if bu bormi bu bormi deb soramaymiz borlarini go ozi bilib qaytaradi

// part 4 No 10 started -----

// outer Containable bu qasyi biri bolsa ham olaver degani Circle yoki Rectancle
func FitInside(outer Containable, points []Point) []Point {
	var result []Point

	for _, p := range points {
		if outer.Contains(p) {
			result = append(result, p)
		}
	}
	return result
}
