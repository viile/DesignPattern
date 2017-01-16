package state

import "fmt"

type FarmInterface interface {
	Harvest()
}

type Farm struct {
	name string
}
func NewFarm(name string) *Farm {
	return &Farm{name:name}
}
func (this *Farm) Harvest() {
	fmt.Println(this.name)
}


type Farmer struct {
	name string
	farm *Farm
}
func NewFarmer() *Farmer{
	nf := new(Farmer)
	nf.name = "spring"
	nf.farm = NewFarm(nf.name)
	return nf
}
func (this *Farmer) Grow(){
	switch this.name {
		case "spring":
			this.name = "summer"
		case "summer":
			this.name = "autumn"
		case "autumn":
			this.name = "winter"
		case "winter":
			this.name = "spring"
		default:
			panic("unkonw")
	}
	this.farm = NewFarm(this.name)
}
func (this *Farmer) Harvest() {
	this.farm.Harvest()
}
