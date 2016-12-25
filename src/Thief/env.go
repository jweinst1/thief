package main

import "fmt"

//file for the env struct

type Env struct {
	items map[string]interface{}
}

func NewEnv() Env {
	return Env{make(map[string]interface{})}
}

//general getter
func (self *Env) get(key string) interface{} {
	val, has := self.items[key]
	if has {
		return val
	} else {
		return nil
	}
}

//general setter
func (self *Env) set(key string, value interface{}) {
	self.items[key] = value
}

func main() {
	g := NewEnv()
	g.set("foo", 4)
	g.set("hello", false)
	fmt.Println(g)

}