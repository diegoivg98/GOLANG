package main

import "fmt"

func main() {

	//declarar variables
	var a float64
	var b float64
	var c float64
    var promedio float64

	//Pedir ingreso de datos
	fmt.Print("Ingrese una nota del 1.0 al 7.0: ")
    fmt.Scan(&a)
    fmt.Print("Ingrese una nota del 1.0 al 7.0: ")
    fmt.Scan(&b)
    fmt.Print("Ingrese una nota del 1.0 al 7.0: ")
    fmt.Scan(&c)
    
    //OBTENER PROMEDIO DE 3 NOTAS
    promedio = (a+b+c)/3

    if(promedio>=4.0){
        fmt.Println("")
    	fmt.Println("APROBADO CON UN:",promedio)
    }else{
        fmt.Println("")
        fmt.Println("REPROBADO CON UN:",promedio)
    }
}