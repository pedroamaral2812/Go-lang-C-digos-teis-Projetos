package main

import (
	"fmt"
)

//A função está pressupondo que a data é valida(sem mes 13 ou dias 32 por exemplo)
//A função está recebendo a data em formato de string.
func DataParaPadraoBrasileiro(data string) string {

	//Cria um slice  de tamanho três para receber Ano,Mes e Dia
	dataExtenso := make([]string, 3)

	//Recebe as datas
	dataExtenso[0] = data[0:4]  // Recebendo o ano
	dataExtenso[1] = data[5:7]  // Recebendo o mes
	dataExtenso[2] = data[8:10] // Recebendo o dia

	//Retorna a data formatada no padrão brasileiro separado por '-'.
	return dataExtenso[2] + "-" + dataExtenso[1] + "-" + dataExtenso[0]

	//Caso queira as data separadas por '/' basta descomentar esta linha e comentar a linha do delimitador '-'
	//return dataExtenso[2] + "/" + dataExtenso[1] + "/" + dataExtenso[0]
}

func main() {
	data := DataParaPadraoBrasileiro("2012-11-05")
	fmt.Println(data)
}
