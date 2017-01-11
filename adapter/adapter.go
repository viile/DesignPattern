package adapter

import "fmt"

type AudioInterface interface {
	Play(string)
}

type Audio struct {
	name string
}
func NewAudio(name string) *Audio {
	audio := &Audio{name:name}
	return audio
}
func (this *Audio) Play(file string) {
	fmt.Println("playing " ,this.name," ",file)
}

type AudioPlayer struct {

}
func (this *AudioPlayer) Play (name string,file string) {
	switch name{
		case "mp3" :
			player := NewAudio("mp3")
			player.Play(file)
		case "wma" :
			player := NewAudio("wma")
			player.Play(file)
		default :
			panic("unkonw name")
	}
}
