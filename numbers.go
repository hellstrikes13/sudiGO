package  main

import (
     "fmt"
)

func main() {
  fmt.Println("1")
  defer  deferme()
  fmt.Println("2")
}

func deferme(){
  fmt.Println("..")
}
