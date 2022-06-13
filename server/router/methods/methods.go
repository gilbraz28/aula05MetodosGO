/* ---------------------------------------------------------------*/
/*                                                                */
/* Pacote que terá funções que atendem às solicitações de acordo  */
/* com o método HTTP e/ou parâmetros via URL.                     */
/*                                                                */
/* Veja a implementação no Youtube:                               */
/* URL:https://youtu.be/lEAEUd4n-Yw                               */
/*                                                                */
/* Por: Rodrigo Messias.                                          */
/* Data: 29 de agosto de 2021.                                    */
/* Versão: 1.0                                                    */
/*                                                                */
/* ---------------------------------------------------------------*/

package methods

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Implementa cliente/servidor HTTP.
// Manipula strings codificadas em UTF-8.

// Métodos HTTP
// GET: Apresenta resultados de recursos (Busca/lista dados).
// POST: Cria um novo recurso (Cadastra dados).
// PUT: Atualiza um recurso (Atualiza dados).
// DELETE: Apaga um recurso (Deleta dados).

// Query Params: Parâmetros nomeados na rota após "?" (Filtro, paginação)
// Router Params: Parâmetros utilizados para identificar recursos. (/id , /name)
// Request Body: Corpo da requisição, utilizado para criar ou alterar recursos. ( enviar dados por POST e PUT)

var data struct {
	Method     string `json: method`
	Descrition string `json: descrition`
}

func HandleMethods(w http.ResponseWriter, r *http.Request) {
	fmt.Println("METHODS.GO - HandleMethods")

	// Servir um arquivo específico.
	// Parâmetros: Resposta, requisição e local do arquivo.
	http.ServeFile(w, r, "./pages/methods.html")

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("METHODS.GO - Index")
	data.Method = r.Method
	data.Descrition = " - Apresentar recursos"

	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("METHODS.GO - Create")
	data.Method = r.Method
	data.Descrition = " - Recurso criado"

	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(data)
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("METHODS.GO - Update")
	data.Method = r.Method
	data.Descrition = " - Recurso atualizado"

	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("METHODS.GO - Delete")
	w.WriteHeader(http.StatusNoContent)
}

func Search(w http.ResponseWriter, r *http.Request) {
	fmt.Println("METHODS.GO - Search")
	params := r.URL.Query()

	//var name = params["name"][0]
	//ou
	var name = params.Get("name")

	var date = strings.ReplaceAll(params.Get("date"), "-", "/")

	data.Method = r.Method
	data.Descrition = " - Buscar por " + name + " e data " + date

	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}

func Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println("METHODS.GO - Show")
	var parts []string = strings.Split(r.URL.Path, "/")

	data.Method = r.Method
	data.Descrition = " - Apresentar recurso  com ID - " + parts[3]

	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)

}
