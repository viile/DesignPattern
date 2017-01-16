package chanOfResponsibility

import "fmt"

type Request struct {
	name string
}
func NewRequest(name string) *Request {
	return &Request{name:name}
}

type HandlerInterface interface {
	SetNext(HandlerInterface)
	Check(Request)
}

type Handler struct {
	name string
	next *Handler
}
func NewHandler(name string) *Handler{
	return &Handler{name:name}
}
func (this *Handler) SetNext(next *Handler) *Handler {
	this.next = next
	return next
}
func (this *Handler) Check(req *Request) {
	fmt.Println(req.name,this.name)
	if this.next != nil {
		this.next.Check(req)
	}
}

type HandlerOne struct {
	Handler
}
type HandlerTwo struct {
	Handler
}
type HandlerThree struct {
	Handler
}
