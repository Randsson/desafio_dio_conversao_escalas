package main

import "fmt"

func main() {

	var kelvin float32

	fmt.Print("Informa a temperatura em Kelvin a ser convertida: ")
	fmt.Scan(&kelvin)

	celsius := kelvin - 273.15

	fmt.Printf("A temperaura %g em Celsius é: º%.2f", kelvin, celsius)
}
