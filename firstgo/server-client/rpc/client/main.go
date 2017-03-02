package main

import (
  "fmt"
  
  "net/rpc"
)


func client() {
  c, err := rpc.Dial("tcp", "127.0.0.1:9990")
  
  retryCount := 3
   defer func(){
       if(retryCount != 0){
           fmt.Println("Here")
      fmt.Println(recover())
       }
      //c, err = rpc.Dial("tcp", "127.0.0.1:9990")
   }()
   
   for ;retryCount >0; retryCount-- {
    if err != nil {
        
            fmt.Println("Retrying: ",retryCount)
            c, err = rpc.Dial("tcp", "127.0.0.1:9990")
            panic(err)
        }
      
    //fmt.Println(err)
    //return
  }
  var result int64
  err = c.Call("Server.Negate", int64(999), &result)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Server.Negate(999) =", result)
  }
}
func main() {
  
  go client()

  var input string
  fmt.Scanln(&input)
}