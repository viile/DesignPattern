package proxy

import "fmt"

type Car interface {
	Drive(string)
}

type BaseCar struct {
	name string
}

func (this *BaseCar) Drive(name string) {
	this.name = name
	fmt.Println(name)
}

type BenzCar struct {
	BaseCar
}

type BmwCar struct {
	BaseCar
}

type AudiCar struct {
	BaseCar
}

type CarFactory struct {
	name string
}

func NewCarFactory (name string) *CarFactory {
	return &CarFactory{name:name}
}

func (this *CarFactory) Product() (car Car) {
	switch this.name {
		case "benz":
			car = new(BenzCar)
		case "bmw":
			car = new(BmwCar)
		case "audi":
			car = new(AudiCar)
		default:
			panic("unkonw car")
	}
	return
}


