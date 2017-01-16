package strategy

import "fmt"

type StrategyInterface interface {
	SomeOp()
}

type Strategy struct {
	name string
}
func (this *Strategy) DoSth(name string) {
	this.name = name
}
func (this *Strategy) SomeOp() {
	fmt.Println(this.name)
}

type StrategyOne struct {
	Strategy
}
type StrategyTwo struct {
	Strategy
}
