package main

import (
	"fmt"

	"github.com/InvisibleTech/go_interfaces/farm"
)

type Animal struct {
	name string
}

// Hen
type Hen struct {
	Animal
	henWay int
}

func (h Hen) WhatsAHenway() int {
	return h.henWay
}

func (h Hen) Name() string {
	return h.name
}

func (h Hen) ProduceFood() string {
	return fmt.Sprintf("%s produces eggs", h.name)
}

func (h Hen) Greet(large farm.LargeFarmAnimal) string {
	return fmt.Sprintf("Hen %s says 'hello' to a %T", h.name, large)
}

// Cow
type Cow struct {
	Animal
	milkFlavor string
}

func (c Cow) WhatsOnTap() string {
	return c.milkFlavor
}

func (c Cow) Name() string {
	return c.name
}

func (c Cow) ProduceFood() string {
	return fmt.Sprintf("%s produces %s milk", c.name, c.milkFlavor)
}

func (c Cow) Greet(small *farm.SmallFarmAnimal) string {
	return fmt.Sprintf("Cow %s says 'hello' to a %T", c.name, *small)
}
