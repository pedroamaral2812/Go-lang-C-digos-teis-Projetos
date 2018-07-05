package main

import (
	"fmt"
	"time"
)

//RetornaDias : Função que retorna a diferença de dias.
//Caso a segunda data for menor que a primeira data o numero vem negativo
func RetornaDias(Dia1, Mes1, Ano1, Dia2, Mes2, Ano2 int) int {

	//Monta a primeira data
	Data1 := time.Date(Ano1, time.Month(Mes1), Dia1, 0, 0, 0, 0, time.UTC)

	//Monta a segunda data
	Data2 := time.Date(Ano2, time.Month(Mes2), Dia2, 0, 0, 0, 0, time.UTC)

	//Subtrai o numero de dias da data 2 em relação a data 1
	difDias := Data2.YearDay() - Data1.YearDay()

	//Retorna a diferença de dias
	return difDias
}

func main() {

	data := RetornaDias(01, 01, 2001, 20, 01, 2001)
	fmt.Println(data)

}
