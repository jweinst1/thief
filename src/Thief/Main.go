package main

import (
   "fmt"
   "io/ioutil"
   "os"
)

func main() {
  //gets command line args
	userArgs := os.Args[1:]

    bytes, err := ioutil.ReadFile(userArgs[0]) // just pass the file name
    if err != nil {
        fmt.Print(err)
    }
    machine := NewEnv()
    ReadBytes(machine, bytes)

}