package observer

import "fmt"

type ObsInterface interface {
	Attach(ObcInterface)
	//Detach(ObcInterface)
	Notify()
}

type Obs struct {
	obcs []ObcInterface
}
func NewObs() *Obs {
	ob := new(Obs)
	return ob
}
func (this *Obs) Attach (obc ObcInterface) {
	this.obcs = append(this.obcs,obc)
}
func (this *Obs) Notify() {
	for i := 0;i<len(this.obcs);i++ {
		this.obcs[i].DoSth()
	}
}

type ObcInterface interface {
	DoSth()
}

type Obc struct {
	name string
}
func NewObc(name string) *Obc{
	return &Obc{name:name}
}
func (this *Obc) DoSth() {
	fmt.Println(this.name)
}
