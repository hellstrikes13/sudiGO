package main
import "fmt"
import "math"
var  p = fmt.Println
func main() {
	a := 100
	c := 100.903490
        var d ,e int = 999,899
        var x  int
	var s string
	const roi  = 8.9
	const tan_theta = 30
	p("1+1 : ",1+1)
	p(false && false)
	p(false && true)
	p(true && false)
	p(true && true)
	p(false || false)
	p(false || true)
	p(true || false)
	p(true || true)
	p(!false)
        p("hi i m ",a)
	var z  = float64(a) // type casting
	p("sin theta after type casting",math.Sin(z))
	p("A numeric constant has no type until it’s given one, such as by an explicit cast. there tan_theta is const with type not specified")
	p("tan theta ",math.Tan(tan_theta))
        fmt.Printf(" i  m  of type %T %T \n",a,c)
        fmt.Printf(" i  m  of type and values %v %v %v\n",a,c,roi)
        fmt.Printf(" i  m  of type and values %#v %#v \n",a,c) // field name and value 
        fmt.Printf(" i  m  of type and values %+v %+v %+v %+v\n",d,e,x,s)

}
