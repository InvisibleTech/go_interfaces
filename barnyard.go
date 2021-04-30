package main

import (
	"fmt"

	"github.com/InvisibleTech/go_interfaces/farm"
)

func doAnimalStuff(a farm.Animal, b farm.Animal) {
	fmt.Printf("Okay we have %s, and %s\n", a.Name(), b.Name())

	switch aType := a.(type) {
	case farm.LargeFarmAnimal:
		switch bType := b.(type) {
		case farm.SmallFarmAnimal:
			//aType.Greets(bType)
			fmt.Println(aType.Greet(&bType))
		default:
			fmt.Println("Do nothing Large and Small")
		}
	case farm.SmallFarmAnimal:
		switch bType := b.(type) {
		case farm.LargeFarmAnimal:
			fmt.Println(aType.Greet(bType))
		}
	default:
		fmt.Println("Do nothing a and b")
	}
}

func main() {
	var h farm.SmallFarmAnimal = Hen{
		Animal: Animal{
			name: "Clucky",
		},
		henWay: 16,
	}

	fmt.Printf("What is the type of h? %T\n", h)
	fmt.Printf("What is the value of h? %+v\n", h)

	var c farm.LargeFarmAnimal = Cow{
		Animal: Animal{
			name: "Bessy",
		},
		milkFlavor: "Anchovy",
	}

	fmt.Printf("What is the type of c? %T\n", c)
	fmt.Printf("What is the value of c? %+v\n", c)

	doAnimalStuff(h, c)
	doAnimalStuff(c, h)
}
