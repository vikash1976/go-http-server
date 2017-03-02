package main

import ("fmt";"flag";"math/rand")

func main() {
  // Define flags
  maxp := flag.Int("max", 6, "-max=<val> or -max <val>")
  // Parse
  flag.Parse()
  // Generate a number between 0 and max
  fmt.Println(rand.Intn(*maxp))
}