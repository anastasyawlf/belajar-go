package main

import (
	"fmt"
)

type Patient struct {
	Name string
	Age  int
	Celcius
}

type Celcius float64

func (c Celcius) toFahrenheit() float64 {
	return float64(c*9/5 + 32)
}

func (c *Celcius) add(value float64) {
	*c += Celcius(value)
}

func main() {
	var temperature Celcius = 20.0
	fmt.Println("temperature:", temperature)
	fmt.Println("Suhu di ruangan ini", temperature.toFahrenheit(), "derajat Fahrenheit")
	temperature.add(3)
	fmt.Println("Suhu di ruangan ini menjadi", temperature, "derajat Celcius")

	fmt.Println("====================================")

	newPatient := Patient{Name: "Kenji", Age: 25, Celcius: 39.0}
	fmt.Printf("newPatient: %+v \n", newPatient)
}
