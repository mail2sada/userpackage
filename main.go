package main

import "fmt"
import mt "github.com/mail2sada/userpackage/math"
import adv "github.com/mail2sada/userpackage/math/advanced"


func main() {

	fmt.Println("Demo package....", mt.Add(10, 20), mt.Subtract(20, 10))

	fmt.Println("result is ", adv.Square(100))

}