package factoryAbstract

import "fmt"

type Farmer struct {

}

func (f *Farmer) Produce(task *string) {
	fmt.Println("corn")
}

type FarmerCreator struct {

}

func (c *FarmerCreator) Create() Producer {
	s := new(Farmer)
	return s
} 
