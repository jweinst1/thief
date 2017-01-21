package main

import (
   "fmt"
   //"io/ioutil"
   //"os"
)

func main() {
  //gets command line args
	/*userArgs := os.Args[1:]

    bytes, err := ioutil.ReadFile(userArgs[0]) // just pass the file name
    if err != nil {
        fmt.Print(err)
    }
    fmt.Println(bytes)*/
    iter := NewByteStream([]byte{44, 55, 55, 66, 47})
    b, ok := iter.next()
    for ;ok; {
      
    }

}