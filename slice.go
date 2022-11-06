package goutils

import (
    "reflect"
)

func Reverse_Slice(slice interface{}) {

    n := reflect.ValueOf(slice).Len()
    swap := reflect.Swapper(slice)
    for i, j := 0, n-1; i < j; i, j = i+1, j-1 { swap(i, j) }

}

func Swap_Slice(index1 int, index2 int, slice interface{}) {

    swap := reflect.Swapper(slice)
    swap(index1, index2)

}

func Slice_Pop(slice interface{}, index int) (bool) {

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

func Slice_Sum(a interface{}) (float64) {
	
    var sum float64 = 0	
    n := reflect.ValueOf(a).Len()
    s := make([]float64, n, n)
	
    if n == 0 { return 0, 0 }
	
    switch a.(type) {
    case []int:
        for i, e := range a.([]int) { s[i] = float64(e) }
    case []float64:
        s = a.([]float64)
    }
	
    for _, e := range s {
	sum += e
    }
    
    return sum
	
}

func InSlice(e interface{}, s interface{}) (bool) {
    
    if Slice_Index(e, s) != -1 {
        return true
    }
    return false

}

func Slice_Index(e interface{}, s interface{}) (int) {

    sr := reflect.ValueOf(s)
    er := reflect.ValueOf(e)
    n := sr.Len()
    for i := 0; i < n; i++ {
        if reflect.DeepEqual(sr.Index(i).Interface(), er.Interface()) { return i }
    }
    return -1

}

func Min(a interface{}) (int, float64) {

    if reflect.ValueOf(a).Len() == 0 { return 0, 0 }

    n := reflect.ValueOf(a).Len()
    s := make([]float64, n, n)

    switch a.(type) {
    case []int:
        for i, e := range a.([]int) { s[i] = float64(e) }
    case []float64:
        s = a.([]float64)
    }

    min := s[0]
    index := 0
    for i, value := range s {
        if value < min {
		        min = value
            index = i
		    }
    }
    return index, min

}

func Max(a interface{}) (int, float64) {

    if reflect.ValueOf(a).Len() == 0 { return 0, 0 }

    n := reflect.ValueOf(a).Len()
    s := make([]float64, n, n)

    switch a.(type) {
    case []int:
        for i, e := range a.([]int) { s[i] = float64(e) }
    case []float64:
        s = a.([]float64)
    }

    max := s[0]
    index := 0
    for i, value := range s {
        if value > max {
			      max = value
            index = i
		    }
    }
    return index, max

}
