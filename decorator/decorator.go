package decorator

import "fmt"

type AnimalInterface interface {
	Whoop()
}

type Animal struct {
	name string
}
func NewAnimal(name string) *Animal {
	animal := &Animal{name:name}
	return animal
}
func (this *Animal) Whoop() {
	fmt.Println(this.name)
}

type AnimalWhoop struct {
	animal *Animal
}
func NewAnimalWhoop(animal *Animal) *AnimalWhoop {
	a := new(AnimalWhoop)
	a.animal = animal
	return a
}
func (this *AnimalWhoop) Whoop() {
	this.animal.Whoop()
}





