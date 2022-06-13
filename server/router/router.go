/* ---------------------------------------------------------------*/
/*                                                                */
/* Pacote que verifica a rota de solicitação e executa a função   */
/* correspondente.                                                */
/*                                                                */
/* URL:https://youtu.be/lEAEUd4n-Yw                               */
/*                                                                */
/* Por: Rodrigo Messias.                                          */
/* Data: 29 de agosto de 2021.                                    */
/* Versão: 1.0                                                    */
/*                                                                */
/* ---------------------------------------------------------------*/

package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gilbraz28/aula05MetodosGO/server/router/account"
	"github.com/gilbraz28/aula05MetodosGO/server/router/login"
	"github.com/gilbraz28/aula05MetodosGO/server/router/methods"
)

// Direciona a aplicação para o pacote da
// funcionalidade conforme a requisição.
func HandleRoutes() {

	fmt.Println("INICIO SERVER")
	fmt.Println("")

	// Adiciona manipuladores ao servidor.
	// Parâmetros: Rota e função a ser executada.
	http.HandleFunc("/", handleLogin)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/account", handleAccount)

	http.HandleFunc("/getAccount", account.ConsultaDados)

	//http.HandleFunc("/", frontend.GetHTML)
	http.HandleFunc("/method", handleMethods)
	http.HandleFunc("/method/", handleValue)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleLogin " + r.Method + "\n")
	methods := map[string]func(){
		"GET":  func() { login.Index(w, r) },
		"POST": func() { login.ValidateLogin(w, r) },
	}
	validateRequest(methods, r.Method)
}

func handleAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleAccount " + r.Method)
	methods := map[string]func(){
		"GET":  func() { account.Index(w, r) },
		"POST": func() { account.ValidateData(w, r) },
	}
	validateRequest(methods, r.Method)
}

func validateRequest(methods map[string]func(), method string) {

	var function func()
	var found bool

	function, found = methods[method]

	if found {
		function()
		return
	}

}

func handleMethods(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleMethods: " + r.Method + " - " + r.URL.Path)

	var methods = map[string]func(){
		"GET":    func() { methods.Index(w, r) },
		"POST":   func() { methods.Create(w, r) },
		"PUT":    func() { methods.Update(w, r) },
		"DELETE": func() { methods.Delete(w, r) },
	}

	var method string = r.Method

	var function func()
	var found bool

	function, found = methods[method]

	if found {
		function()
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func handleValue(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleValue: " + r.Method + " - " + r.URL.Path)

	var parts []string = strings.Split(r.URL.Path, "/")

	var numberParts = len(parts)

	if numberParts <= 4 {
		var functions = map[string]func(){
			"id":     func() { methods.Show(w, r) },
			"search": func() { methods.Search(w, r) },
		}

		var function func()
		var found bool

		function, found = functions[parts[2]]

		if found {
			function()
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}
}
