package main

import (
	"fmt"
)

const ebuK = 373.0

func main() {
	tempK := ebuK
	tempC := (tempK - 273.0)
	fmt.Println("A temperatura de ebulição em Celsius é: ", tempC, "°C")

}
