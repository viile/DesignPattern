package bridge

import "fmt"

type PersonInterface interface {
	Eat(string)
}

type Person struct {
	name string
	tool Tool
}

func NewPerson(name string,tool Tool) *Person{
	var _person = new(Person)
	_person.name = name
	_person.tool = tool
	return _person
}
func (this *Person) Eat(food string) {
	fmt.Println(this.name," use ",this.tool.Show()," eat ",food)
}

type ToolInterface interface {
	Use(string)
	Show()
}

type Tool struct {
	name string
}
func (this *Tool ) Use(name string){
	this.name = name
}
func (this *Tool) Show() string {
	fmt.Println(this.name)
	return this.name
}
