package main

import "fmt"

func main() {
    var edades [5]int
    var nombres [5]string
    for i := 0 ; i < len(edades); f++ {
        fmt.Print("Ingrese nombre:")
        fmt.Scan(&nombres[i])
        fmt.Print("Ingrese edad:")
        fmt.Scan(&edades[i])
    }
    fmt.Println("Personas mayores de edad.")
    for i := 0; i < len(edades); i++ {
        if edades[i] >= 18 {
            fmt.Print(nombres[i]," ")
        }
    }
}