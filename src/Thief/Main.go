package main

import (
   "fmt"
   "io/ioutil"
   "os"
)

func main() {
  //gets command line args
	userArgs := os.Args[1:]

    b, err := ioutil.ReadFile(userArgs[0]) // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

    fmt.Println(b) // print the content as 'bytes'

    str := string(b) // convert content to a 'string'

    fmt.Println(str) // print the content as a 'string'
    //fmt.Println(Byte_to_i([]byte{0, 0, 0, 14}))
}