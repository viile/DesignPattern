package prototype

import "fmt"

type Coder struct {
	name string
	language string
}

func (this *Coder) SetName(name string) {
	this.name = name
}

func (this *Coder) SetLanguage(language string) {
	this.language = language
}


func (this *Coder) Show() {
	fmt.Println("name : ",this.name," language : ",this.language)
}

func (this *Coder) Clone() *Coder{
	newObj := (*this)
	return &newObj
}


