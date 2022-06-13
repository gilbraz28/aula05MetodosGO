package login

import (
	"fmt"
	"net/http"

	us "github.com/gilbraz28/aula05MetodosGO/server/router/controller/user"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index - Carga ")
	fmt.Println("")
	http.ServeFile(w, r, "./pages/login.html")
}

func ValidateLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ValidateLogin")

	var email = r.PostFormValue("email")
	var password = r.PostFormValue("password")

	fmt.Printf("EMAIL E PASSWORD: %s - %s", email, password)

	if us.SearchUser(email, password) {
		fmt.Fprintf(w,
			`<h1> Iniciar Novo Acesso </h1>
			<p> Método: %s </p>
			<p> Email: %s </p>
			<p> Senha: %s </p>`, r.Method, email, password)

		return
	}

	http.Redirect(w, r, "./login", http.StatusMovedPermanently)

}
