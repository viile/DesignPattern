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
	"proxy"
	"filter"
	"template"
	"strategy"
	"state"
	"observer"
	"chanOfResponsibility"
	"visotor"
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
	//proxy
	prasd := proxy.NewCarFactory("benz")
	prp := prasd.Product()
	prp.Drive("shanghai")
	//filter
	filas1 := filter.NewPerson("male","basketball")
	filas2 := filter.NewPerson("male","football")
	filas3 := filter.NewPerson("male","swim")
	filas4 := filter.NewPerson("female","swim")
	filas5 := filter.NewPerson("female","football")
	var persons []*filter.Person
	persons = append(persons,filas1)
	persons = append(persons,filas2)
	persons = append(persons,filas3)
	persons = append(persons,filas4)
	persons = append(persons,filas5)
	fildo := filter.NewFilterGender("female")
	fmt.Println(fildo.DoFilter(persons))
	//template
	temone := new(template.SmsSO)
	temone.Sms.Parent = temone
	temone.Initialize("one")
	temone.Send("18723123")
	temtwo := new(template.SmsST)
	temtwo.Sms.Parent = temtwo
	temtwo.Initialize("two")
	temtwo.Send("123123123")
	//strategy
	straa := new(strategy.StrategyOne)
	straa.DoSth("one")
	straa.SomeOp()
	strab := new(strategy.StrategyTwo)
	strab.DoSth("two")
	strab.SomeOp()
	//state
	staaa := state.NewFarmer()
	staaa.Harvest()
	staaa.Grow()
	staaa.Harvest()
	staaa.Grow()
	staaa.Harvest()
	staaa.Grow()
	staaa.Harvest()
	//observer
	obaaaa := observer.NewObs()
	obc1 := observer.NewObc("1")
	obc2 := observer.NewObc("2")
	obc3 := observer.NewObc("3")
	obc4 := observer.NewObc("4")
	obaaaa.Attach(obc1)
	obaaaa.Attach(obc2)
	obaaaa.Attach(obc3)
	obaaaa.Attach(obc4)
	obaaaa.Notify()
	//chanOfResponsibility
	chr := chanOfResponsibility.NewRequest("test")
	chha1 := chanOfResponsibility.NewHandler("one")
	chha2 := chanOfResponsibility.NewHandler("two")
	chha3 := chanOfResponsibility.NewHandler("three")
	chha1.SetNext(chha2).SetNext(chha3)
	chha1.Check(chr)
	//visotor
	vp := visotor.NewPerson("zhangsan")
	vp.Eat(new(visotor.VisotorAmerica))
	vp.Eat(new(visotor.VisotorAsia))
}
