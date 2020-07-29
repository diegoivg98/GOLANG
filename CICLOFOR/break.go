package main

import "fmt"

func main() {
		/*INGRESAR TANTOS NUMEROS,
		SUMARLOS,
		PERO SI INGRESA UN 0 SE DETIENE EL CICLO
		*/
	var suma int
	var n int

	for i := 1; i > 0; i++ {
		fmt.Print("INGRESE UN NUMERO: ")
		fmt.Scan(&n)
		suma += n //obtener suma de los numeros
		i+=1

		if(n==0){//en caso de ingresar 0
			break //romper ciclo for
		}
	}
	fmt.Print("TOTAL DE LA SUMA: ",suma)
}