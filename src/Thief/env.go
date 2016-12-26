package main


//file for the env struct

//null type for variable declaration and other
type null struct {

}

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

func (self *Env) declare(key string) {
	self.items[key] = null{}
}

func (self *Env) is_null(key string) bool {
	_, ok := self.items[key].(null)
	return ok
}

