package main



//file for the stack component of the virtual machine

type Stack struct {
	items map[int]interface{}
	count int
}

func NewStack() Stack {
	return Stack{make(map[int]interface{}), 0}
}

func (self *Stack) push(item interface{}) {
	self.items[self.count] = item
	self.count++
}

func (self *Stack) pop() interface{} {
	if self.count == 0 {
		//error
		return nil
	} else {
		item := self.items[self.count-1]
		delete(self.items, self.count-1)
		self.count--
		return item
	}
}

func (self *Stack) peek() interface{} {
	if self.count == 0 {
		//error
		return nil
	} else {
		return self.items[self.count-1]
	}
}


