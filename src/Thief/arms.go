package main

import "fmt"

//subregisters with specialized behavior
//these are called arms



type Printarm struct {
	val interface{}
}

func NewPrintArm(arg interface{}) Printarm {
	return Printarm{arg}
}

func (self *Printarm) put(value interface{}){
	self.val = value
}

func (self *Printarm) print(){
	fmt.Println(self.val)
}



func main(){
	
}