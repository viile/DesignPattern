package main

import (
	"fmt"
	"singleton"
	"factory"
	"factoryAbstract"
)

func main() {
	fmt.Println("start")
	mSingleton, nSingleton := singleton.NewSingleton("hello"), singleton.NewSingleton("hi")
	mSingleton.DoSomething()
	nSingleton.DoSomething()

	var fac factory.CarFactory
	car := fac.CreateCar("benz")
	car.Create("benz 3")
	car.Drive()

	c := new(factoryAbstract.ProgrammerCreator)
	p := c.Create()
	taskProject := "project"
	p.Produce(&taskProject)
}
