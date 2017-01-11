package facade

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
	dog *Animal
	frog *Animal
}
func NewAnimalWhoop() *AnimalWhoop {
	animal := new(AnimalWhoop)
	animal.dog = NewAnimal("dog")
	animal.frog = NewAnimal("frog")
	return animal
}
func (this *AnimalWhoop) WhoopDog() {
	this.dog.Whoop()
}
func (this *AnimalWhoop) WhoopFrog() {
	this.frog.Whoop()
}





