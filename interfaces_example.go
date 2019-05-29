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
type human  struct {
  s  string
}

type alien  struct {
  s  string
}
func  (d dog) speaks() string {
  return d.s
}
func  (c cat) speaks() string {
  return c.s
}

func  (h human) speaks() string {
  return h.s
}


func  (a alien) speaks() string {
  return a.s
}
func display(sp speaking) {
  fmt.Println(sp.speaks())
}

func main() {
  d :=  dog{s:"bow bow"}
  c :=  cat{s:"mew mew"}
  a :=  alien{s:"sdfdsfds"}
  h :=  human{s:"hello"}
  display(d)
  display(c)
  display(a)
  display(h)
}
