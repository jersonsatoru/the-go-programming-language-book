package temperature

import "fmt"

type Celsius float64
type Fahrenheit float64

func ConvertToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*(9/5) + 32)
}

func ConvertToCelsius(f Fahrenheit) Celsius {
	return Celsius((f - 32) / (9 / 5))
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g celsius", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g fahrenheit", f)
}
