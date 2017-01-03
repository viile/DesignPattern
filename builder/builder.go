package builder

import "fmt"

type HardwareInterface interface {
	Produce(name string)
	Show()
}

type Hardware struct {
	name string
}
func (this *Hardware) Produce(name string) {
	this.name = name
}
func (this *Hardware) Show() string {
	//fmt.Println(this.name)
	return this.name
}

type Cpu struct {
	Hardware
}
type Ram struct {
	Hardware
}
type Storage struct {
	Hardware
}
type Camera struct {
	Hardware
}
type Screen struct {
	Hardware
}

type Phone struct {
	name string
	cpu Cpu
	ram Ram
	storage Storage
	camera Camera
	screen Screen
}
func NewPhone(name,cpu,ram,storage,camera,screen string) Phone {
	_phone := new(Phone)
	_phone.name = name
	_phone.cpu.Produce(cpu)
	_phone.ram.Produce(ram)
	_phone.storage.Produce(storage)
	_phone.camera.Produce(camera)
	_phone.screen.Produce(screen)
	return *_phone
}
func (this *Phone) Display() {
	fmt.Println("name : ",this.name,this.cpu.Show(),this.ram.Show())
}
