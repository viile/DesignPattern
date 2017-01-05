package main

import (
	"fmt"
	"singleton"
	"factory"
	"factoryAbstract"
	"prototype"
	"builder"
	"bridge"
)

func main() {
	fmt.Println("start")
	//singleton
	mSingleton, nSingleton := singleton.NewSingleton("hello"), singleton.NewSingleton("hi")
	mSingleton.DoSomething()
	nSingleton.DoSomething()
	//factory
	var fac factory.CarFactory
	car := fac.CreateCar("benz")
	car.Create("benz 3")
	car.Drive()
	//factoryAbstract
	c := new(factoryAbstract.ProgrammerCreator)
	p := c.Create()
	taskProject := "project"
	p.Produce(&taskProject)
	//prototype
	pa := new(prototype.Coder)
	pa.SetName("a")
	pa.SetLanguage("c++")
	pb := pa.Clone()
	pb.SetName("b")
	pb.SetLanguage("c#")
	pa.Show()
	pb.Show()
	//builder
	ba := builder.NewPhone("yummy","i7","16g","256g","1200w","6")
	ba.Display()
	//bridge
	bba := new(bridge.ToolByFork)
	bba.Use("fork")
	fmt.Println(bba)
	//bba.Show()
	//bbb := bridge.NewPerson("viile",bba)
	//bbb.Eat("apple")
}
