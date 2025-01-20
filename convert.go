package goutils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"
)

const (
	date_layoutUS = "January 2, 2006"
)

// Generics

func String_To_Date(date_str string) time.Time {

	date, _ := time.Parse(date_layoutUS, date_str)
	return date

}

func To_Struct(inter interface{}, struc interface{}) {

	mapstructure.Decode(inter, struc)

}

func Bytes_To_Json(byteValue []byte, json_data interface{}) {

	json.Unmarshal(byteValue, json_data)

}

// Int Converters
func (i Int) Value() int {
	return int(i)
}
func (i_slice Int_Slice) Value() []int {
	tmp_slice := make([]int, len(i_slice))
	for i := range i_slice {
		tmp_slice[i] = i_slice[i].Value()
	}
	return tmp_slice
}
func (i Int) To_Float64() Float64 {
	return Float64(i)
}
func (i Int) To_Int64() Int64 {
	return Int64(i)
}
func (i Int) To_String() String {
	return String(strconv.Itoa(int(i)))
}
func (i Int) To_Bool() Bool {
	if i == 1 {
		return Bool(true)
	}
	return Bool(false)
}

// Float64 Converters
func (f Float64) Value() float64 {
	return float64(f)
}
func (f_slice Float64_Slice) Value() []float64 {
	tmp_slice := make([]float64, len(f_slice))
	for i := range f_slice {
		tmp_slice[i] = f_slice[i].Value()
	}
	return tmp_slice
}
func (f Float64) To_Int() Int {
	return Int(f)
}
func (f Float64) To_Int64() Int64 {
	return Int64(f)
}
func (f Float64) To_String() String {
	return String(fmt.Sprintf("%g", float64(f)))
}
func (f Float64) To_Bool() Bool {
	if f == 1.0 {
		return Bool(true)
	}
	return Bool(false)
}

// String converters
func (s String) Value() string {
	return string(s)
}
func (s_slice String_Slice) Value() []string {
	tmp_slice := make([]string, len(s_slice))
	for i := range s_slice {
		tmp_slice[i] = s_slice[i].Value()
	}
	return tmp_slice
}
func (s String) To_Int() Int {
	i, _ := strconv.Atoi(string(s))
	return Int(i)
}
func (s String) To_Float64() Float64 {
	f, _ := strconv.ParseFloat(string(s), 64)
	return Float64(f)
}
func (s String) To_Bool() Bool {
	switch s {
	case "true", "True", "TRUE", "1":
		return Bool(true)
	default:
		return Bool(false)
	}
}
func (s String) To_Int64() Int64 {
	i, _ := strconv.Atoi(string(s))
	return Int64(i)
}

// Bool converters
func (b Bool) Value() bool {
	return bool(b)
}
func (b Bool) To_Int() Int {
	if b {
		return 1
	}
	return 0
}
func (b Bool) To_Int64() Int64 {
	if b {
		return Int64(1)
	}
	return Int64(0)
}
func (b Bool) To_Float64() Float64 {
	return Float64(b.To_Int())
}
func (b Bool) To_String() String {
	if b {
		return String("true")
	}
	return String("false")
}

// Int64 Converters
func (i Int64) Value() int64 {
	return int64(i)
}
func (i Int64) To_Int() Int {
	return Int(i)
}
func (i Int64) To_Float64() Float64 {
	return Float64(i)
}
func (i Int64) To_String() String {
	return String(strconv.FormatInt(int64(i), 10))
}
func (i Int64) To_Bool() Bool {
	if int(i) == 1 {
		return Bool(true)
	}
	return Bool(false)
}
