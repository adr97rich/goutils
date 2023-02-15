package goutils

import (
    "fmt"
    "math/rand"
)

func Abs(x float64) (float64) {

    if x < 0 {
        return -1*x
    } else {
        return x
    }

}

func Round_Float64(f float64, dec int) (float64) {

    str := fmt.Sprintf("%." + To_String(dec) + "f", f)
    return To_Float64(str)

}

func Random_Float64(inf float64, sup float64) (float64) {

    return inf + rand.Float64() * (sup - inf)

}
