package main
import "fmt"

func main() {
  for i :=0 ; i < 11; i++  {
	fmt.Println(i)
  }
  for {
  fmt.Println("blah blah blah infinite loop ")
  break
  }
  for j := 1 ; j < 101; j++ {
   if  j % 7 == 0 && j % 11 == 0 {
   fmt.Println("numbers divisible by 7 and 11 are",j)
   }
  }
  fmt.Println(" i shall print all odd numbers")
  for i :=0 ; i < 11; i++  {
   if i  % 2 == 0 {
	continue
   }
	fmt.Println(i)
  }
  if  num := 768 ; num < 500 {
   fmt.Println("num: ",num,"is not less than 500")
  }else if num > 500 {
   fmt.Println("num: ",num,"is greter than 500")
  }else {
   fmt.Println("GL")
  }
}
