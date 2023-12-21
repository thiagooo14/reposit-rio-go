package main

import "fmt"

const ebulicaoK = 373

func main() {
	tempK := ebulicaoK
	tempC := tempK - 273

	fmt.Println("Ebolução em °F", tempK)
	fmt.Println("Ebolução em °C", tempC)
}
