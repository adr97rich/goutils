package goutils

import (
	"fmt"
	"math"
	"math/rand"
)

// Generics

func abs(x float64) float64 {

	if x < 0 {
		return -1 * x
	} else {
		return x
	}

}

func Random_Float64(inf, sup float64) float64 {

	return inf + (rand.Float64() * (sup - inf))

}

func Random_Int(inf, sup int) int {

	return inf + (rand.Int() * (sup - inf))

}

// For goutils types

// For Absolute Value
func (f Float64) Abs() Float64 {
	return Float64(abs(f.Value()))
}
func (i Int) Abs() Int {
	return Int(abs(i.To_Float64().Value()))
}
func (i Int64) Abs() Int64 {
	return Int64(abs(i.To_Float64().Value()))
}

// For rounding a number
func (f Float64) Round(dec int) Float64 {
	dec_str := ("%." + Int(dec).To_String() + "f")
	str := String(fmt.Sprintf(dec_str.Value(), float64(f)))
	return str.To_Float64()
}

// Power
func (f Float64) Power(n float64) Float64 {
	return Float64(math.Pow(f.Value(), n))
}
func (i Int) Power(n float64) Float64 {
	return Float64(math.Pow(i.To_Float64().Value(), n))
}
