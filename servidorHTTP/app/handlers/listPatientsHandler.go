package handlers

import (
	"html/template"
	"net/http"
	"servidorHTTP/app/utils"
)

func ListPatientsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método inválido", http.StatusMethodNotAllowed)
		return
	}

	patients, err := utils.GetPatients()
	if err != nil {
		http.Error(w, "Erro ao buscar pacientes", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("static/forms/listPatients.html")
	if err != nil {
		http.Error(w, "Erro ao carregar a página de pacientes", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, patients)
	if err != nil {
		http.Error(w, "Erro ao renderizar a página de pacientes", http.StatusInternalServerError)
		return
	}
}
