package main

import (
	"fmt"

	"github.com/jersonsatoru/golang-book/src/part2/temperature"
)

func main() {
	var c temperature.Celsius = 10.0
	f := temperature.ConvertToFahrenheit(c)
	fmt.Println(f.String())
	fmt.Println(temperature.ConvertToCelsius(f).String())
}
