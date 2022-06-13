package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	us "github.com/gilbraz28/aula05MetodosGO/server/router/controller/user"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CARGA ACCOUNT")
	fmt.Println("")

	http.ServeFile(w, r, "./pages/account.html")
}

func ValidateData(w http.ResponseWriter, r *http.Request) {

	fmt.Println("ValidateData ACCOUNT ")
	fmt.Println("")

	var data us.Data

	json.NewDecoder(r.Body).Decode(&data)

	us.NewUser(data)

	//ConsultaDados(w, r)

	http.Redirect(w, r, "./account", http.StatusMovedPermanently)

}

func ConsultaDados(w http.ResponseWriter, r *http.Request) {

	fmt.Println("ConsultaDados ACCOUNT ")
	fmt.Println("")

	grupoUser := us.GetUsers()
	var i int = 1

	fmt.Println("DADOS GET LEN: ", len(grupoUser))

	for _, data := range grupoUser {

		fmt.Printf("\nUser: %d - Nome: %s- Email: %s ", i, data.Name, data.Email)
		i++

		// fmt.Fprintf(w,
		// 	`<h1> Iniciar Novo Acesso </h1>
		// 	<p> Método: %s </p>
		// 	<p> Email: %s </p>
		// 	<p> Senha: %s </p>`, r.Method, email, password)

	}

	//http.Redirect(w, r, "./login", http.StatusOK)
}
