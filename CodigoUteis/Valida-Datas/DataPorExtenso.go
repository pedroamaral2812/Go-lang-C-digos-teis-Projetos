package main

//A função está pressupondo que a data é valida, mas há uma validação de mês invalido
//Está recebendo uma string, se for um tipo data é necessário transforma-la em string no formato dd/mm/aaaa ou dd-mm-aaaa.
func DataExtenso(data string) string {

	dataExtenso := make([]string, 3)

	//Recebe as datas
	dataExtenso[0] = data[0:2]
	dataExtenso[1] = data[3:5]
	dataExtenso[2] = data[6:10]

	if dataExtenso[1] == "01" {
		return dataExtenso[0] + " de Janeiro de " + dataExtenso[2]
	} else if dataExtenso[1] == "02" {
		return dataExtenso[0] + " de Fevereiro de " + dataExtenso[2]
	} else if dataExtenso[1] == "03" {
		return dataExtenso[0] + " de Março de " + dataExtenso[2]
	} else if dataExtenso[1] == "04" {
		return dataExtenso[0] + " de Abril de " + dataExtenso[2]
	} else if dataExtenso[1] == "05" {
		return dataExtenso[0] + " de Maio de " + dataExtenso[2]
	} else if dataExtenso[1] == "06" {
		return dataExtenso[0] + " de Junho de " + dataExtenso[2]
	} else if dataExtenso[1] == "07" {
		return dataExtenso[0] + " de Julho de " + dataExtenso[2]
	} else if dataExtenso[1] == "08" {
		return dataExtenso[0] + " de Agosto de " + dataExtenso[2]
	} else if dataExtenso[1] == "09" {
		return dataExtenso[0] + " de Setembro de " + dataExtenso[2]
	} else if dataExtenso[1] == "10" {
		return dataExtenso[0] + " de Outubro de " + dataExtenso[2]
	} else if dataExtenso[1] == "11" {
		return dataExtenso[0] + " de Novembro de " + dataExtenso[2]
	} else if dataExtenso[1] == "12" {
		return dataExtenso[0] + " de Dezembro de " + dataExtenso[2]
	} else {
		return "Mês invalido"
	}
}

/* Para testar a função
func main() {
	data := DataExtenso("22-00-2010")
	fmt.Println(data)

}
*/
