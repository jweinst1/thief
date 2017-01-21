package main

import "fmt"


//file for machine functions and operations

//prints string format of some bytes
func print(env Env, data []byte) {
	fmt.Println(string(data))
}


var OPERS map[byte]func(Env, []byte)

//0 is reserved as the null terminator
func init(){
	OPERS = map[byte]func(Env, []byte){
		1:print,
	}
}


