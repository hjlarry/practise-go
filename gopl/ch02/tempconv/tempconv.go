package tempconv

import "fmt"

// Celsius 摄氏温度
type Celsius float64

// Fahrenheit 华氏温度
type Fahrenheit float64

const (
	// AbsoluteZeroC 绝对零度
	AbsoluteZeroC Celsius = -273.15
	// FreezingC 冰点温度
	FreezingC Celsius = 0
	// BoilingC 沸点温度
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
