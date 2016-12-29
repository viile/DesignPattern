package singleton

import "fmt"

type Singleton interface {
	DoSomething()
}

type singleton struct {
	text string
}

var _singleton Singleton

func NewSingleton(text string) Singleton {
	if _singleton == nil {
		_singleton = &singleton{text:text}
	}

	return _singleton
}

func (this *singleton) DoSomething() {
	fmt.Println(this.text)
}
