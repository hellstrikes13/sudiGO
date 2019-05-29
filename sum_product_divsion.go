package  main

import (
     "fmt"
     "errors"
)
func spm(x , y int) (int , int , int, int, error ) {
  if y == 0  {
   return 0,0,0,0 , errors.New("Divide by zero y cannot be 0")
  }
  return x + y , x - y ,  x* y , x / y , nil
}

func main() {
  s ,m,p,d,err :=  spm(5,0)
  if err != nil {
  fmt.Println("there was an error: ",err)
  return
  }
  fmt.Println(s,m,p,d)
}

