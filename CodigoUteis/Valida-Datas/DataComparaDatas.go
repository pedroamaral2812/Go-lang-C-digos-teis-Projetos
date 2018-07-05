package main

import (
	"fmt"
	"time"
)

//CompararDatas : Função que compara duas Datas e verifica qual a maior retornando:
// -1 : Data 1 é maior que a data 2
//  0 : Data 1 é igual a Data 2
//  1 : Data 2 é maior que a Data 1
func CompararDatas(Dia1, Mes1, Ano1, Dia2, Mes2, Ano2 int) int {

	//Monta a primeira data
	Data1 := time.Date(Ano1, time.Month(Mes1), Dia1, 0, 0, 0, 0, time.UTC)

	//Monta a segunda data
	Data2 := time.Date(Ano2, time.Month(Mes2), Dia2, 0, 0, 0, 0, time.UTC)

	//Subtrai o numero de dias da data 2 em relação a data 1
	//days := Data2.Sub(Data1).Hours() / 24
	days := Data2.Sub(Data1).Hours() / 24

	//Se a data 1 for maior que a data 2
	if days < 0 {
		return -1
		//Se a data 1 for igual que a data 2
	} else if days == 0 {
		return 0
		//Se a data 2 for maior que a data 1
	} else {
		return 1
	}
}

func main() {

	data := CompararDatas(01, 04, 2001, 02, 01, 2001)
	fmt.Println(data)

}
