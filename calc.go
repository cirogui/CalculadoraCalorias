package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)            //ir para o diretório do template
	http.HandleFunc("/process", processor) //tratar a requisição http
	http.ListenAndServe(":8080", nil)      //criar servidor http local

}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil) //executar o template
}

func processor(w http.ResponseWriter, r *http.Request) {
	//se o método utilizado não for post, retorna
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	//cada variável vai receber os dados dos formulários em formato string
	Peso := r.FormValue("Peso")
	Sexo := r.FormValue("Sexo")
	Altura := r.FormValue("Altura")
	Idade := r.FormValue("Idade")
	FatorAtividade := r.FormValue("Fator de Atividade")

	//variáveis para converter string para float
	intPeso, err := strconv.ParseFloat(Peso, 32)
	fmt.Println(err)
	intAltura, err := strconv.ParseFloat(Altura, 32)
	intIdade, err := strconv.ParseFloat(Idade, 32)
	intFatorAtividade, err := strconv.ParseFloat(FatorAtividade, 32)

	//fórmula Método Harris-Benedict já com os dados em float para serem calculados
	FormulaMasc := 66 + (13.7 * intPeso) + (5 * intAltura) - (6.8 * intIdade)
	FormulaFem := 655 + (9.6 * intPeso) + (1.8 * intAltura) - (4.7 * intIdade)

	if Sexo == "Masculino" {

		Result := FormulaMasc * intFatorAtividade
		var intResult int = int(Result)
		Complemento := "Seu gasto calórico diário é de aproximadamente: "
		Complemento2 := "kcal"
		Complemento3 := "<p>Para perda de peso intensa: consumir aprox 500kcal abaixo do total.<p>Para perda de peso moderada: consumir aprox 300kcal abaixo do total.</p><p>Para manutenção do peso: consumir a quantidade total de kcal.<p>Para aumento de peso: consumir aprox de 300 a 500kcal acima do total."

		d := struct {
			Text  int
			Comp  string
			Comp2 string
			Comp3 string
		}{
			Comp:  Complemento,
			Text:  intResult,
			Comp2: Complemento2,
			Comp3: Complemento3,
		}

		tpl.ExecuteTemplate(w, "index.gohtml", d)

	} else {

		Result := FormulaFem * intFatorAtividade
		var intResult int = int(Result)
		Complemento := "Seu gasto calórico diário é de aproximadamente: "
		Complemento2 := "kcal"
		Complemento3 := "<p>Para perda de peso intensa: consumir aprox 500kcal abaixo do total.<p>Para perda de peso moderada: consumir aprox 300kcal abaixo do total.</p><p>Para manutenção do peso: consumir a quantidade total de kcal.<p>Para aumento de peso: consumir aprox de 300 a 500kcal acima do total."

		d := struct {
			Text  int
			Comp  string
			Comp2 string
			Comp3 string
		}{
			Comp:  Complemento,
			Text:  intResult,
			Comp2: Complemento2,
			Comp3: Complemento3,
		}

		tpl.ExecuteTemplate(w, "index.gohtml", d)

	}
}
