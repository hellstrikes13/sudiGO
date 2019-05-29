package main
import "fmt"

func main() {
  cc := map[string] string {
    "India" : "Delhi",
    "Australia" : "Canbera",
    "US": "Washington DC",
  }
  for k ,v  := range cc {
    fmt.Printf("%s --> %s \n",k,v)
  }

  cs := map[string]map[string]string {
    "India": map[string]string{ "Maharashtra": "mumbai","Karnatak":"bengaluru"},
    "US": map[string][string]{"California" : "sanjose"},
   }
  fmt.Println(cs)
}
