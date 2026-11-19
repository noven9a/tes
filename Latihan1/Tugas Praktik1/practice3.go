package main

import "fmt"

func main() {
	var panjang int
	var lebar int
	var luas int

	fmt.Println("Panjang:")
	fmt.Scanln(&panjang)
	fmt.Println("Lebar:")
	fmt.Scanln(&lebar)

	luas = panjang * lebar 
	fmt.Println ("Luas Persegi Panjang")

	fmt.Println(luas)
}

