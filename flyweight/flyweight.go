package flyweight

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
	Cars map[string]Car
}

func (this *CarFactory) CreateCar(name string) (car Car) {
	defer fmt.Println(this.Cars)
	switch name {
		case "benz":
			c,ok := this.Cars["benz"]
			if ok {
				return c
			} else {
				car = new(BenzCar)
				this.Cars["benz"] = car
			}
		case "bmw":
			c,ok := this.Cars["bmw"]
			if ok {
				return c
			} else {
				car = new(BmwCar)
				this.Cars["bmw"] = car
			}
		case "audi":
			c,ok := this.Cars["audi"]
			if ok {
				return c
			} else {
				car = new(AudiCar)
				this.Cars["audi"] = car
			}
		default:
			panic("unkonw car")
	}
	return
}

