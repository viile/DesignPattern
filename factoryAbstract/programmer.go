package factoryAbstract

import "fmt"

type Programmer struct {

}

func (p *Programmer) Produce(task *string) {
		fmt.Println("code",*task)
}

type ProgrammerCreator struct {

}

func (c * ProgrammerCreator) Create() Producer {
	s := new(Programmer)
	return s
}
