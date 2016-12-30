package factory

import "fmt"

type Car interface {
	Create(string)
	Drive()
}

type BaseCar struct {
	name string
}

func (this *BaseCar) Create(name string) {
	this.name = name
}
func (this *BaseCar) Drive() {
	fmt.Println(this.name)
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

}

func (this *CarFactory) CreateCar(name string) (car Car) {
	switch name {
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

