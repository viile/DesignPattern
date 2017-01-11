package composite

import "fmt"

type FileInterface interface {
	Show()
}

type File struct {
	name string
}
func NewFile(name string) *File {
	return &File{name:name}
}
func (this *File) Show() {
	fmt.Println(this.name)
}

type Folder struct {
	name string
	files []FileInterface
}
func NewFolder(name string) *Folder {
	folder := &Folder{name:name}
	return folder
}
func (this *Folder) Add(file FileInterface) {
	this.files = append(this.files,file)
}
func (this *Folder) Show() {
	fmt.Println(this.name)
	for i := 0; i < len(this.files);i++ {
		this.files[i].Show()
	}
}
