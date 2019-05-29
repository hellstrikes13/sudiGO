package main
import "fmt"
var p = fmt.Println
func main() {
 m  := map[string] int {
   "bond" : 007,
   "khiladi" : 768, //note (,) comma here is compulsary or else go will cry
 }
p(m,"len: ",len(m))
//To create an empty map, use the builtin make: make(map[key-type]val-type).
c := make(map[string]int)
p(c,"len: ",len(c))

v := m["bond"]
p(v)
var v1,fine = m["X"]
p(v1,fine)

//OR u can write this way _ indicates blank identifier we dont need value so we ignored it !

names := map[string]string{"sudeep": "melekar","mahesh": "kolhe","chetan":"gavankar"}
_, blah := names["amit"]
p(names,blah)

var k , ok = m["khiladi"]
p(k,ok)

delete(m,"khiladi")
p("after deletion",m)


country_capital := map[string]string  {
 "india" : "delhi",
 "us" : "washignton DC",
 "newzeland" : "wellington",
}
for k , v  := range(country_capital) {
 p ( k ,"\t->\t" ,v)
}
p("\nNESTED MAPS....\n")
country_state_capital := map[string]map[string]string {
 "india":  {
   "maharastra " : "mumbai",
   "punjab" : "chandigarh",
   "rajasthan" : "jaipur",
   },
 "US" : {
  "maryland" : "annapolis",
  "NeyYork"  : "albany",
  },
}
 for c,s := range country_state_capital["india"] {
    p(c,"\t->\t ",s)
 }
 print("displaying ASCII value\n")
 for ks,vs  := range "ABC" {
   p(ks,"\t",vs)
 }
}
