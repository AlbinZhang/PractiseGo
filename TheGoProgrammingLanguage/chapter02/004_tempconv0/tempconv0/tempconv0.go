package tempconv0

import (
	"fmt"
)

type Celsius float64    //摄氏温度
type Fahrenheit float64 //华氏温度
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15       //绝对零度
	FreezingC     Celsius = 0             //结冰点温度
	BoilingC      Celsius = 100           //沸水温度
	AbsoluteZeroK         = AbsoluteZeroC //绝对零度

)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c + 273) }
func KToC(k Kelvin) Celsius { return Celsius(k - 273) }

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func test() {
	c := FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
}
