var form = document.getElementsByTagName('form')[0];

form.addEventListener("submit", getData)

function sendData(method, uri, header, data, serverResponse) {
    let httpRequest = new XMLHttpRequest();

    httpRequest.open(method, uri);

    httpRequest.setRequestHeader("X-Content-Type-Options", header)

    httpRequest.send(data);

    httpRequest.onreadystatechange = serverResponse;
}

function getData(event) {
    event.preventDefault();

    //EXEMPLO: Formato normal, apenas envio dos dados
    //let data =  new FormData(form);
    //EXEMPLO: Formato JSON

    let data = {
    	name: form.name.value,
    	lastname: form.lastname.value,
    	email: form.email.value,
    	password: form.password.value,
    	confirm: form.confirm.value,
    	accept: form.accept.checked
    };

    //let json =  JSON.stringify(data);

    //ADICIONANDO NOVO VALOR NO ATRIBUTO DO PROJETO QUANDO NECESSARIO
    //data.append("telefone", form.telefone.value);

    //sendData("POST", "./register", "multipart/form-data", data, response)
    //sendData("POST", "./registerJson", "application/json", json, response)

    //let header = new Headers({"X-Content-Type-Options":"application/json"});

    //header.append("X-Content-Type-Options","application/json");
    // ou Jogando acima e utilizando comando abaixo
    fetch("./account", {
            headers: new Headers({"X-Content-Type-Options":"application/json"}),
            method: "POST",
            body: JSON.stringify(data),
            //body: new FormData(form),
    }).then((response) => {
        const responseStatus = {
            200:() => { alert("Dados enviados com sucesso."); },
            400:() => { alert("Este cadastro já existe."); },
            404:() => { alert("Tente realizar o cadastro mais tarde."); }
        }

        if (responseStatus[response.status]) {
            let responseUser = responseStatus[response.status];

            //responseUser();
        } else {
            alert("Realize o cadastro mais tarde.");
        }
    })

}