package main
import "fmt"
var p = fmt.Println
func main() {
 var  a[4] int
 p("empty array with 5 len having  all 0",a)

 b := [6]string {"s","u","d","e","e","p"}
 p("array b of type string",b,"len of arr is",len(b))

 var matrix  [2][3]int
   for i:= 0 ;i < 2 ; i++ {
     for j := 0 ; j < 3 ; j++ {
       matrix[i][j] = i + j
     }
   }
  p("2 dimensional array or matrix",matrix)

//
s := make( []string ,6)
p("emty slice",s)
s[0] =  "s"
s[1] = "u"
s[2] = "d"
s[3] = "e"
s[4] = "e"
s[5] = "p"
p(s,"len: ",len(s),"capacity: ",cap(s))
s = append(s,"X")
p(s,"len: ",len(s),"capacity: ",cap(s))

s = append(s,"Y","Z")
p(s,"len: ",len(s),"capacity: ",cap(s))

s = append(s,"2","1","3","4","5","6")
p(s,"len: ",len(s),"capacity: ",cap(s))


s1 := make([]string ,len(s))
copy(s1,s)
p("s1",s1)
p(s1[2:])
//negative slice is not allowed	

 t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)








}
