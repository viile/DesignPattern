package filter

import "fmt"

type FilterInterface interface {
	DoFilter([]Person)
}

type Filter struct {

}
type FilterGender struct {
	filter string
}
func NewFilterGender(filter string) *FilterGender {
	return &FilterGender{filter:filter}
}

func (this *FilterGender) DoFilter(persons []*Person) []*Person {
	var result []*Person
	for i:=0;i<len(persons);i++ {
		fmt.Println(persons[i])
		if persons[i].Gender() == this.filter {
			result = append(result,persons[i])
		}
	}
	return result
}

type PersonInterface interface {
	Gender(string)
	Sport(string)
}

type Person struct {
	gender string
	sport string
}
func NewPerson(gender,sport string) *Person {
	return &Person{gender:gender,sport:sport}
}
func (this *Person) Gender() string {
	fmt.Println(this.gender)
	return this.gender
}
func (this *Person) Sport() string {
	return this.sport
}
