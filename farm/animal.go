package farm

type Animal interface {
	Name() string
}

type FarmAnimal interface {
	Animal
	ProduceFood() string
}

type LargeFarmAnimal interface {
	FarmAnimal
	Greet(animal *SmallFarmAnimal) string
}

type SmallFarmAnimal interface {
	FarmAnimal
	Greet(animal LargeFarmAnimal) string
}

type SmallPredator interface {
	Animal
	Eat(animal *SmallFarmAnimal)
}

type LargePredator interface {
	Animal
	Eat(animal *LargeFarmAnimal)
	EatCompetition(animal *SmallPredator)
}
