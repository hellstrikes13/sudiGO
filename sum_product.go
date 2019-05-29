package  main

import (
     "fmt"
)
func spm(x int , y int) (int , int , int) {
  return x + y , x - y ,  x* y
}

func main() {
  s , _, m :=  spm(5,2)
  a , b , c := spm(100,50)
  fmt.Println(s,m)
  fmt.Println(a,b,c)
}

