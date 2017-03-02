package main

import (
  "fmt"
  "net"
  "net/rpc"
)

type Server struct {}
func (this *Server) Negate(i int64, reply *int64) error {
  *reply = -i
  return nil
}

func server() {
  rpc.Register(new(Server))
  ln, err := net.Listen("tcp", ":9990")
  if err != nil {
    fmt.Println(err)
    return
  }
  for {
    c, err := ln.Accept()
    if err != nil {
      continue
    }
    go rpc.ServeConn(c)
  }
}

func main() {
  go server()
 

  var input string
  fmt.Scanln(&input)
}