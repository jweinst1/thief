package main


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

func (self *Env) contains(key string) bool {
	_, has := self.items[key]
	return has
}

func (self *Env) declare(key string) {
	self.items[key] = Null{}
}

func (self *Env) is_null(key string) bool {
	_, ok := self.items[key].(Null)
	return ok
}

//deletes an item from storage
func (self *Env) del(key string) {
	delete(self.items, key)
}

