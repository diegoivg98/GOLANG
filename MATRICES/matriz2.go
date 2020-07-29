package main

import "fmt"

func main() {
    var mat [4][4]int //matriz de 4x4
    for f :=0 ; f < 4; f++ {
        for c := 0; c < 4; c++ {
            fmt.Print("Ingrese componente:")
            fmt.Scan(&mat[f][c])
        }
    }
    fmt.Println();

    //IMPRIMIR MATRIZ
    for f := 0; f < 4; f++ {
        for c := 0; c < 4; c++ {
            fmt.Print(mat[f][c], " ")
        }
        fmt.Println();
    }
            fmt.Println();

    fmt.Println("Elementos de la diagonal principal.")
    for k := 0; k < 4; k++ {//OBTENER LA DIAGONAL DE LA MATRIZ
        fmt.Print(mat[k][k]," ")
    }        
            
    
}