package goutils

import (
    "fmt"
    "time"
    "strconv"
    "encoding/json"

    "github.com/mitchellh/mapstructure"
)

const (
    date_layoutUS = "January 2, 2006"
)

// Converters //

func To_String(T interface{}) (string) {
    
    switch T.(type) {
    case int:
        return strconv.Itoa(T.(int))
    case bool:
        return strconv.FormatBool(T.(bool))
    case float64:
        return fmt.Sprintf("%g", T.(float64))
    default:
        return T.(string)
    }

}

func To_Int(T interface{}) (int) {

    switch T.(type) {
    case string:
        intVar, _ := strconv.Atoi(T.(string))
        return intVar
    case float64:
        return int(T.(float64))
    default:
        return T.(int)
    }

}

func To_Float64(T interface{}) (float64) {

    switch T.(type) {
    case string:
        float64Var, _ := strconv.ParseFloat(T.(string), 64)
        return float64Var
    case int:
        return float64(T.(int))
    default:
        return T.(float64)
    }

}

func String_To_Date(date_str string) (time.Time) {

    date, _ := time.Parse(date_layoutUS, date_str)
    return date

}

func To_Struct(inter interface{}, struc interface{}) {

    mapstructure.Decode(inter, struc)

}

func Bytes_To_Json(byteValue []byte, json_data interface{}) {

    json.Unmarshal(byteValue, json_data)

}