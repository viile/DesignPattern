package main

import (
	"fmt"
	"singleton"
	"factory"
	"factoryAbstract"
	"prototype"
	"builder"
	"bridge"
	"flyweight"
	"facade"
	"adapter"
	"decorator"
	"composite"
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
	var bba bridge.Tool
	bba.Use("fork")
	fmt.Println(bba)
	bba.Show()
	bbb := bridge.NewPerson("viile",bba)
	bbb.Eat("apple")
	//flyweight
	var flyc flyweight.CarFactory
	flyc.Cars = make(map[string]flyweight.Car)
	fcar1 := flyc.CreateCar("benz")
	fcar2 := flyc.CreateCar("benz")
	fcar3 := flyc.CreateCar("audi")
	fcar1.Create("benz 3")
	fcar1.Drive()
	fmt.Println(fcar2,fcar3)
	//facade
	fabbb := facade.NewAnimalWhoop()
	fabbb.WhoopDog()
	fabbb.WhoopFrog()
	//adapter
	adapl := new(adapter.AudioPlayer)
	adapl.Play("mp3","one")
	adapl.Play("wma","two")
	//decorator
	deapl := decorator.NewAnimal("frog")
	deaa := decorator.NewAnimalWhoop(deapl)
	deaa.Whoop()
	//composite
	comroot := composite.NewFolder("root")
	comdata := composite.NewFolder("data")
	comtmp := composite.NewFolder("tmp")
	comf1 := composite.NewFile("f1.go")
	comroot.Add(comdata)
	comdata.Add(comtmp)
	comtmp.Add(comf1)
	comroot.Show()
}
