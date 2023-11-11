package goutils

import (
    "fmt"
    
    "github.com/eiannone/keyboard"
)

// field_size : Max size of text (must be greater than 1).
// prompt : Text displayed before input.
// border : Delimiters of textbox (if not needed can be left like an empty string).
func Read_String(field_size int, prompt, border string) (string) {

    var str string = ""
    
    if field_size < 2 {
        panic("goutils.Read_String(): Max size of text must be greater than 1.")
    }

    keys, _ := keyboard.GetKeys(1)
    defer func() {
        _ = keyboard.Close()
    }()

    // Print prompt and textbox
    fmt.Printf(prompt + border)
    for i := 0; i < field_size + 2; i++ { fmt.Print(" ") }
    fmt.Printf(border)
    for j := 0; j < field_size + len(border) + 1; j++ { fmt.Printf("\x08") }

    for {
        e := <- keys
        s := ""
        // Characters to type
        if e.Rune != 0 {
            s = string(e.Rune)
        } else if e.Key == 32 {    // Space
            s = " "
        }
        if s != "" {
            if len(str) < field_size {
                fmt.Printf(s)
                str += s
            } else {
                for i := 0; i < field_size; i++ {
                    fmt.Printf("\x08" + " " + "\x08")
                }
                str += s
                fmt.Printf("_" + str[len(str)-field_size+1:])
            }
        }
        // Backspace and Enter
        if (e.Key == 127 || e.Key == 8) && 0 < len(str) {
            str = str[0:len(str)-1]
            if len(str) < field_size {
                for i := 0; i < len(str)+1; i++ {
                    fmt.Printf("\x08" + " " + "\x08")
                }
                fmt.Printf(str)
            } else if len(str) == field_size {
                for i := 0; i < field_size; i++ {
                    fmt.Printf("\x08" + " " + "\x08")
                }
                fmt.Printf(str)
            } else if field_size < len(str) {
                for i := 0; i < field_size; i++ {
                    fmt.Printf("\x08" + " " + "\x08")
                }
                fmt.Printf("_" + str[len(str)-field_size+1:])
            }
        } else if e.Key == 13 {
            break
        }
    }

    fmt.Println("")

    return str    // Final string inputed by user

}
