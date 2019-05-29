package main
import "fmt"
import "time"
var p = fmt.Println
func main() {
 i := 2
 p("hello there i m ",i)
 switch i {
  case 0:
    p("sunday")
  case 1:
    p("monday")
  case 2:
    p("tuesday")
 }

 switch  time.Now().Weekday() {
  case  time.Saturday, time.Sunday:
   p("weekend party time...")
  default:
   p("its boring weeek day...")
 }
 t := time.Now()
  p("date time now is : ",t)
  // here no expression or literal for switch it assumes to be a boolean and there can be multiple case statement
 switch {
   case t.Hour() < 12:
    p("its before noon")
   case t.Hour() > 12:
    p("its after 12")
 }

 whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a boolean")
        case int:
            fmt.Println("I'm an integer")
        default:
            fmt.Printf("Don't know type GL.. %T\n", t)
        }
    }
    whatAmI(false)
    whatAmI(1)
    whatAmI("hey")

}
