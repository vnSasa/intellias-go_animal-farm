package main

import(
	"fmt"
	"reflect"
)

const (
	dogEating int = 10
	catEating int = 7
	cowEating int = 25

	dogMinWeight int = 4
	catMinWeight int = 2
	cowMinWeight int = 45

	isDogEdible bool = false
	isCatEdible bool = false
	isCowEdible bool = true
)

type farmController interface {
	foodNeeding() int
	animalTypeGetter() string
	minWeightGetter() int
	animalWeightGetter() int
	isAnimalEdible() bool
	edibleStatus() bool
	fmt.Stringer
}

type dog struct {
	animalType string
	weight int
	needingFood int
	isEdible bool
}

func (d dog) foodNeeding() int {
	return d.weight * d.needingFood
}

func (d dog) String() string {
	return fmt.Sprintf("The weight of the %v is %vkg. ", d.animalType, d.weight)
}

func (d dog) animalTypeGetter() string {
	return d.animalType
}

func (d dog) minWeightGetter() int {
	return dogMinWeight
}

func (d dog) animalWeightGetter() int {
	return d.weight
}

func (d dog) isAnimalEdible() bool {
	return isDogEdible
}

func (d dog) edibleStatus() bool {
	return d.isEdible
}

type cat struct {
	animalType string
	weight int
	needingFood int
	isEdible bool
}

func (c cat) foodNeeding() int {
	return c.weight * c.needingFood
}

func (c cat) String() string {
	return fmt.Sprintf("The weight of the %v is %vkg. ", c.animalType, c.weight)
}

func (c cat) animalTypeGetter() string {
	return c.animalType
}

func (c cat) minWeightGetter() int {
	return catMinWeight
}

func (c cat) animalWeightGetter() int {
	return c.weight
}

func (c cat) isAnimalEdible() bool {
	return isCatEdible
}

func (c cat) edibleStatus() bool {
	return c.isEdible
}

type cow struct {
	animalType string
	weight int
	needingFood int
	isEdible bool
}

func (c cow) foodNeeding() int {
	return c.weight * c.needingFood
}

func (c cow) String() string {
	return fmt.Sprintf("The weight of the %v is %vkg. ", c.animalType, c.weight)
}

func (c cow) animalTypeGetter() string {
	return c.animalType
}

func (c cow) minWeightGetter() int {
	return cowMinWeight
}

func (c cow) animalWeightGetter() int {
	return c.weight
}

func (c cow) isAnimalEdible() bool {
	return isCowEdible
}

func (c cow) edibleStatus() bool {
	return c.isEdible
}

func validAnimalType(animal farmController) error {
	animalType := reflect.TypeOf(animal).Name()
	if animal.animalTypeGetter() != animalType {
		return fmt.Errorf("%s can`t be %s", animalType, animal.animalTypeGetter())
	}
	return nil
}

func validAnimalWeight(animal farmController) error {
	if animal.minWeightGetter() > animal.animalWeightGetter() {
		return fmt.Errorf("Minimum weight of this type of animal is %vkg. Weight of the animal is %vkg", animal.minWeightGetter(), animal.animalWeightGetter())
	}
	return nil
}

func validIsAnimalEdible(animal farmController) error {
	if animal.edibleStatus() != animal.isAnimalEdible() {
		return fmt.Errorf("Mistake of edible status. Edible status of %v is %v", animal.animalTypeGetter(), animal.isAnimalEdible())
	}
	return nil
}

func validFarm(farm []farmController) error {
	var err error
	for _, v := range farm {
		
		if err = validAnimalType(v); err != nil {
			return fmt.Errorf("Type validation failed: %w", err)
		}
		
		if err = validAnimalWeight(v); err != nil {
			return fmt.Errorf("Weight validation failed: %w", err)
		}
		
		if err = validIsAnimalEdible(v); err != nil {
			return fmt.Errorf("Edible validation failed: %w", err)
		}
	}
	return nil
}

func main()  {

	var totalNeedingFood int

	animalFarm := []farmController {
		dog{
			animalType : "dog",
			weight : 10,
			needingFood : dogEating,
			// isEdible : false,
		},
		cat{
			animalType : "cat",
			weight : 5,
			needingFood : catEating,
			// isEdible : false,
		},
		dog{
			animalType : "dog",
			weight : 4,
			needingFood : catEating,
			// isEdible : false,
		},
		cow{
			animalType : "cow",
			weight : 98,
			needingFood : cowEating,
			isEdible : true,
		},
	}

	for _, a := range animalFarm {
		
		err := validFarm(animalFarm)
		if err != nil {
			fmt.Println(err)
			return
		}

		totalNeedingFood += a.foodNeeding()
		fmt.Printf(a.String() + "Necessary food for a month - %vkg\n", a.foodNeeding())
	}

	fmt.Printf("Total food needing: %d", totalNeedingFood)

}