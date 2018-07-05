package main

import (
	"fmt"
	"time"
)

//DataAdicionarSubtrairDias : Função para adicionar o numero de dias correspondente ao passado pelo usuario
// A saída é uma string com a data no formato americano. Há uma função que converte formato americano em brasileiro neste repositorio
//Para subtrair basta apenas passar o parametro como numero negativo. Ex:data := DataAdicionarSubtrairDias(-30)
func DataAdicionarSubtrairDias(numeroDias int) string {

	//Guarda em uma variavel a data atual
	DataInicial := time.Now()

	//Adiciona o numero de dias passado por parametro
	DataInicialAdicionada := DataInicial.AddDate(0, 0, numeroDias)

	//Formatando a data pro modelo americano
	DataFormatada := DataInicialAdicionada.Format("2006-01-02")

	//Retorna a data
	return DataFormatada
}

//DataAdicionarSubtrairMes : Função para adicionar/subtrair o numero de meses correspondente ao passado pelo usuario
// A saída é uma string com a data no formato americano. Há uma função que converte formato americano em brasileiro neste repositorio
//Para subtrair basta apenas passar o parametro como numero negativo. Ex:data := DataAdicionarSubtrairMes(-30)
func DataAdicionarSubtrairMes(numeroMes int) string {

	//Guarda em uma variavel a data atual
	DataInicial := time.Now()

	//Adiciona o numero de mes passado por parametro
	DataInicialAdicionada := DataInicial.AddDate(0, numeroMes, 0)

	//Formatando a data pro modelo americano
	DataFormatada := DataInicialAdicionada.Format("2006-01-02")

	//Retorna a data
	return DataFormatada
}

//DataAdicionarSubtrairAno : Função para adicionar/subtrair o numero de anos correspondente ao passado pelo usuario
// A saída é uma string com a data no formato americano. Há uma função que converte formato americano em brasileiro neste repositorio
//Para subtrair basta apenas passar o parametro como numero negativo. Ex:data := DataAdicionarSubtrairAno(-2)
func DataAdicionarSubtrairAno(numeroAno int) string {

	//Guarda em uma variavel a data atual
	DataInicial := time.Now()

	//Adiciona o numero de anos passado por parametro
	DataInicialAdicionada := DataInicial.AddDate(numeroAno, 0, 0)

	//Formatando a data pro modelo americano
	DataFormatada := DataInicialAdicionada.Format("2006-01-02")

	//Retorna a data
	return DataFormatada
}

func main() {

	data := DataAdicionarSubtrairDias(30)
	//data := DataAdicionarSubtrairMes(2)
	//data := DataAdicionarSubtrairAno(2)

	fmt.Printf(data)
}
