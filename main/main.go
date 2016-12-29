package main

import (
	"fmt"
	"singleton"
)

func main() {
	mSingleton, nSingleton := singleton.NewSingleton("hello"), singleton.NewSingleton("hi")
	mSingleton.DoSomething()
	nSingleton.DoSomething()
}
