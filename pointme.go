package main
import "fmt"
func main() {
  p := 20
  fmt.Println("address of var  p : ",&p,"contents of var p:",p)
  pt := &p
  fmt.Println("address of pointer  pt : ", &pt ,"contents of var p:", pt ,"points to var p", *pt)

}

/*
func main() {

	// Declare an integer variable with the value of 20.
	value := 20

	// Display the address of and value of the variable.
	fmt.Println("Address Of:", &value, "Value Of:", value)

	// Declare a pointer variable of type int. Assign the
	// address of the integer variable above.
	p := &value

	// Display the address of, value of and the value the pointer
	// points to.
	fmt.Println("Address Of:", &p, "Value Of:", p, "Points To:", *p)
}
*/
