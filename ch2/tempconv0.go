package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g摄氏度", c)
}

func (f Fahrenheit) String() string  {
	return fmt.Sprintf("%g华氏度", f)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	fmt.Println("沸点：")
	fmt.Println(BoilingC.String())
	fmt.Println(CToF(BoilingC).String())
	fmt.Printf("绝对零度：%v\n", AbsoluteZeroC)
}


