package main

import "fmt"

type featuresStruct struct {
	Color   [3]string
	weight  int
	height  int
	opacity float64
	status  bool
}

func main() {

	// u := auth.User{
	// 	Name:  "Otabek",
	// 	Phone: 123456789,
	// }

	// products := map[string]float64{
	// 	"Apple":  1.99,
	// 	"Banana": 0.99,
	// 	"Orange": 1.49,
	// }

	// category := map[string]map[string]float64{
	// 	"Fruits": products,
	// }

	featuresMap := map[string]string{
		"Color":   "Red",
		"weight":  "100",
		"height":  "200",
		"opacity": "0.99",
	}

	featuresStruct := featuresStruct{
		Color:   [3]string{"Red", "Green", "Blue"},
		weight:  100,
		height:  200,
		opacity: 0.99,
		status:  true,
	}

	fmt.Println(featuresMap)
	fmt.Println(featuresStruct)

	// fmt.Println("Category:")
	// fmt.Println(category)
	// fmt.Println("Products:")
	// fmt.Println(products)
	// fmt.Println(u.GetData())
}
