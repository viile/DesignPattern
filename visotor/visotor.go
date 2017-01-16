package visotor

import "fmt"

type Person struct {
	name string
}
func NewPerson(name string) *Person {
	return &Person{name:name}
}
func (this *Person) Eat(v VisotorInterface) {
	v.Eat(this.name)
}

type VisotorInterface interface {
	Eat(string)
}
type VisotorAmerica struct {

}
func (this *VisotorAmerica) Eat (name string) {
	fmt.Println(name," kfc")
}
type VisotorAsia struct {

}
func (this *VisotorAsia) Eat (name string) {
	fmt.Println(name," shaxian")
}
