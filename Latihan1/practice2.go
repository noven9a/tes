// package main

// import "fmt"

// func main() {
// 	var a int
// 	var b string
// 	var c float32
// 	var d bool

// 	a = 16
// 	b = "hello world"
// 	c = 14.1
// 	d = true

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// }


package main 

import "fmt"

func main() {
    var celcius float64
    fmt.Scan(&celcius)
    
    fahrenheit := (9.0 / 5.0) * celcius + 32.0
    
    fmt.Printf("%.2f\n", fahrenheit)
}