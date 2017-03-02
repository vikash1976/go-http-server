package main

import (
  "encoding/gob"
  "fmt"
  "net"
  "os"
)
func client(msg string) {
  // connect to the server
  c, err := net.Dial("tcp", "127.0.0.1:9999")
  if err != nil {
    fmt.Println(err)
    return
  }

  // send the message
  //msg := msg //"Hello World"
  fmt.Println("Sending", msg)
  err = gob.NewEncoder(c).Encode(msg)
  if err != nil {
    fmt.Println(err)
  }

  c.Close()
}

func main() {
 
  for {
    var input string
    fmt.Scanln(&input)
    switch input {
        case "q":
        fmt.Println("Q")
        os.Exit(0)
        default: 
        go client(input)
    }
    
  }

  
}