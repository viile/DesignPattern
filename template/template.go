package template

import (
	"fmt"
	"math/rand"
	"strconv"
	"reflect"
)

type SmsInterface interface {
	Initialize(string)
	MakeText()
	Send(string)
}
type Sms struct {
	text string
	config string
	mobile string
	Parent interface{}
}
func (this *Sms) MakeText() {
	this.text = this.config + strconv.Itoa(rand.Intn(100))
}
func (this *Sms) Send(mobile string) {
	this.MakeText()
	c := reflect.ValueOf(this.Parent)
	method := c.MethodByName("SendSms")
	params := make([]reflect.Value,1)
	params[0] = reflect.ValueOf(mobile)
	method.Call(params)
}
func (this *Sms) Initialize(config string) {
	this.config = config
}
type SmsSO struct {
	Sms
}
func (this *SmsSO) SendSms(mobile string) {
	this.Sms.mobile = mobile
	fmt.Println(mobile,this.Sms.text)
	fmt.Println(mobile,this.Sms.text)
}
type SmsST struct {
	Sms
}
func (this *SmsST) SendSms(mobile string) {
	this.Sms.mobile = mobile
	fmt.Println(mobile,this.Sms.text)
}
