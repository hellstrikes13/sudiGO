package main
import "fmt"

type speaking  interface {
  speaks() string
}

type dog  struct {
  s  string
}

type cat  struct {
  s  string
}

func  (d dog) speaks() string {
  return "baw baw"
}
func  (c cat) speaks() string {
  return "mew mew"
}

func display(sp speaking) {
  fmt.Println(sp.speaks())
}

func main() {
  d :=  dog{}
  c :=  cat{}
  display(d)
  display(c)
}
