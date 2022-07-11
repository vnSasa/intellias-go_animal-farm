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

type Dog struct {
	animalType string
	weight int
	needingFood int
	isEdible bool
}

func (dog Dog) foodNeeding() int {
	return dog.weight * dog.needingFood
}

func (dog Dog) String() string {
	return fmt.Sprintf("The weight of the %v is %vkg. ", dog.animalType, dog.weight)
}

func (dog Dog) animalTypeGetter() string {
	return dog.animalType
}

func (dog Dog) minWeightGetter() int {
	return dogMinWeight
}

func (dog Dog) animalWeightGetter() int {
	return dog.weight
}

func (dog Dog) isAnimalEdible() bool {
	return isDogEdible
}

func (dog Dog) edibleStatus() bool {
	return dog.isEdible
}

type Cat struct {
	animalType string
	weight int
	needingFood int
	isEdible bool
}

func (cat Cat) foodNeeding() int {
	return cat.weight * cat.needingFood
}

func (cat Cat) String() string {
	return fmt.Sprintf("The weight of the %v is %vkg. ", cat.animalType, cat.weight)
}

func (cat Cat) animalTypeGetter() string {
	return cat.animalType
}

func (cat Cat) minWeightGetter() int {
	return catMinWeight
}

func (cat Cat) animalWeightGetter() int {
	return cat.weight
}

func (cat Cat) isAnimalEdible() bool {
	return isCatEdible
}

func (cat Cat) edibleStatus() bool {
	return cat.isEdible
}

type Cow struct {
	animalType string
	weight int
	needingFood int
	isEdible bool
}

func (cow Cow) foodNeeding() int {
	return cow.weight * cow.needingFood
}

func (cow Cow) String() string {
	return fmt.Sprintf("The weight of the %v is %vkg. ", cow.animalType, cow.weight)
}

func (cow Cow) animalTypeGetter() string {
	return cow.animalType
}

func (cow Cow) minWeightGetter() int {
	return cowMinWeight
}

func (cow Cow) animalWeightGetter() int {
	return cow.weight
}

func (cow Cow) isAnimalEdible() bool {
	return isCowEdible
}

func (cow Cow) edibleStatus() bool {
	return cow.isEdible
}

func validAnimalType(animal farmController) error {
	var err error
	animalType := reflect.TypeOf(animal).Name()
	if animal.animalTypeGetter() != animalType {
		err = fmt.Errorf("%s can`t be %s", animalType, animal.animalTypeGetter())
	}
	return err
}

func validAnimalWeight(animal farmController) error {
	var err error
	if animal.minWeightGetter() > animal.animalWeightGetter() {
		err = fmt.Errorf("Minimum weight of this type of animal is %vkg. Weight of the animal is %vkg", animal.minWeightGetter(), animal.animalWeightGetter())
	}
	return err
}

func validIsAnimalEdible(animal farmController) error {
	var err error
	if animal.edibleStatus() != animal.isAnimalEdible() {
		err = fmt.Errorf("Mistake of edible status. Edible status of %v is %v", animal.animalTypeGetter(), animal.isAnimalEdible())
	}
	return err
}

func validFarm(farm []farmController) error {
	var err error
	for _, v := range farm {
		err = validAnimalType(v)
		if err != nil {
			err = fmt.Errorf("Type validation failed: %w", err)
			return err
		}
		err = validAnimalWeight(v)
		if err != nil {
			err = fmt.Errorf("Weight validation failed: %w", err)
			return err
		}
		err = validIsAnimalEdible(v)
		if err != nil {
			err = fmt.Errorf("Edible validation failed: %w", err)
			return err
		}
	}
	return nil
}

func main()  {

	var totalNeedingFood int

	animalFarm := []farmController {
		Dog{
			animalType : "Dog",
			weight : 10,
			needingFood : dogEating,
			// isEdible : false,
		},
		Cat{
			animalType : "Cat",
			weight : 5,
			needingFood : catEating,
			// isEdible : false,
		},
		Dog{
			animalType : "Dog",
			weight : 4,
			needingFood : catEating,
			// isEdible : false,
		},
		Cow{
			animalType : "Cow",
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