package goutils

import (
	"reflect"
	"strings"
	"time"
)

// To sort slices in ascending order
// 's' must be a pointer to a slice of int, float64, time or string
// Algorithm used: quick sort
func Sort_Slice(s interface{}) {

	var slice reflect.Value = reflect.ValueOf(s)
	var quickSort func(reflect.Value, int, int)
	var partition func(reflect.Value, int, int) int = func(slice reflect.Value, l, h int) int {
		i := l
		slice_indir := reflect.Indirect(slice)
		pivot := slice_indir.Index(h)
		swap := reflect.Swapper(slice_indir.Interface())
		for j := l; j < h; j++ {
			switch pivot.Kind() {
			case reflect.Int:
				if slice_indir.Index(j).Int() < pivot.Int() {
					swap(i, j)
					i++
				}
			case reflect.Float64:
				if slice_indir.Index(j).Float() < pivot.Float() {
					swap(i, j)
					i++
				}
			case reflect.ValueOf(time.Now()).Kind(): // Sort times
				isBefore := func(first reflect.Value, sec reflect.Value) bool {
					return first.MethodByName("Before").Call([]reflect.Value{sec})[0].Bool()
				}
				if isBefore(slice_indir.Index(j), pivot) {
					swap(i, j)
					i++
				}
			case reflect.String: // Sort strings by alphabetic order
				j_first_letter := strings.ToUpper(slice_indir.Index(j).String())[0]
				pivot_first_letter := strings.ToUpper(pivot.String())[0]
				if int(j_first_letter) < int(pivot_first_letter) {
					swap(i, j)
					i++
				}
			}
		}
		swap(i, h)
		return i
	}
	quickSort = func(slice reflect.Value, l, h int) {
		if l < h {
			p := partition(slice, l, h)
			quickSort(slice, l, p-1)
			quickSort(slice, p+1, h)
		}
	}

	if slice.Kind() != reflect.Pointer {
		panic("Not a pointer to a slice.")
	} else if slice.Kind() == reflect.Pointer && reflect.Indirect(slice).Kind() != reflect.Slice {
		panic("Pointer does not point to a slice.")
	} else if reflect.Indirect(slice).Index(0).Type().String() == "interface {}" {
		panic("Invalid slice type.")
	}

	quickSort(slice, 0, reflect.Indirect(slice).Len()-1)

}

func Reverse_Slice(slice interface{}) {

	n := reflect.ValueOf(slice).Len()
	swap := reflect.Swapper(slice)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}

}

func Swap_Slice(index1 int, index2 int, slice interface{}) {

	swap := reflect.Swapper(slice)
	swap(index1, index2)

}

func Slice_Pop(slice interface{}, index int) bool {

	sr := reflect.ValueOf(slice)
	sr_indir := reflect.Indirect(sr)
	n := sr_indir.Len()
	ns := reflect.MakeSlice(sr.Type().Elem(), 0, 0)
	poped := false
	for i := 0; i < n; i++ {
		if i != index {
			ns = reflect.Append(ns, sr_indir.Index(i))
		} else {
			poped = true
		}
	}
	sr.Elem().Set(ns)
	return poped

}

func Slice_Remove_Duplicate(slice interface{}) {

	sr := reflect.ValueOf(slice)
	sr_indir := reflect.Indirect(sr)
	n := sr_indir.Len()
	tmp_ns := reflect.MakeSlice(sr.Type().Elem(), 0, 0)
	ns := reflect.MakeSlice(sr.Type().Elem(), 0, 0)
	for i := n - 1; i >= 0; i-- {
		duplicated := false
		for j := i - 1; j >= 0; j-- {
			if reflect.DeepEqual(sr.Elem().Index(i).Interface(), sr.Elem().Index(j).Interface()) {
				duplicated = true
				break
			}
		}
		if !duplicated {
			tmp_ns = reflect.Append(tmp_ns, sr_indir.Index(i))
		}
	}

	for i := tmp_ns.Len() - 1; i >= 0; i-- {
		ns = reflect.Append(ns, tmp_ns.Index(i))
	}

	sr.Elem().Set(ns)

}

func InSlice(e interface{}, s interface{}) bool {

	return Slice_Index(e, s) != -1

}

func Slice_Index(e interface{}, s interface{}) int {

	sr := reflect.ValueOf(s)
	er := reflect.ValueOf(e)
	n := sr.Len()
	for i := 0; i < n; i++ {
		if reflect.DeepEqual(sr.Index(i).Interface(), er.Interface()) {
			return i
		}
	}
	return -1

}

func Length(s interface{}) Int {
	return Int(reflect.ValueOf(s).Len())
}

// For Goutils types

// Max Value
func (f_slice Float64_Slice) Max() (int, Float64) {
	max := f_slice[0]
	index := 0
	for i, value := range f_slice {
		if value > max {
			max = value
			index = i
		}
	}
	return index, Float64(max)
}
func (i_slice Int_Slice) Max() (int, Int) {
	max := i_slice[0]
	index := 0
	for i, value := range i_slice {
		if value > max {
			max = value
			index = i
		}
	}
	return index, Int(max)
}

// Min Value
func (f_slice Float64_Slice) Min() (int, Float64) {
	max := f_slice[0]
	index := 0
	for i, value := range f_slice {
		if value < max {
			max = value
			index = i
		}
	}
	return index, Float64(max)
}
func (i_slice Int_Slice) Min() (int, Int) {
	max := i_slice[0]
	index := 0
	for i, value := range i_slice {
		if value < max {
			max = value
			index = i
		}
	}
	return index, Int(max)
}

// Summary
func (f_slice Float64_Slice) Sum() Float64 {

	var sum Float64 = 0

	for _, e := range f_slice {
		sum += Float64(e)
	}

	return sum

}
func (i_slice Int_Slice) Sum() Int {

	var sum Int = 0

	for _, e := range i_slice {
		sum += e
	}

	return sum

}

// Adding New String
func (s String) Prepend(n_s string) String {
	return String(n_s + s.Value())
}

func (s String) Append(n_s string) String {
	return String(s.Value() + n_s)
}

// Concatenate String Slice
func (s_slice String_Slice) Concatenate(sep string) String {
	str := String("")
	l := Length(s_slice)
	for i := Int(0); i < l; i++ {
		switch i {
		case (l - 1):
			str = str.Append(s_slice[i].Value())
		default:
			str = str.Append((s_slice[i].Value() + sep))
		}
	}
	return str
}
