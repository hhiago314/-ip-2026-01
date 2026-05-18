package handlers

import (
	"net/http"
	"servidorHTTP/app/utils"
	"strconv"
)

func CreatePatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método inválido", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	ageStr := r.FormValue("age")
	bloodType := r.FormValue("blood_type")

	if name == "" || ageStr == "" || bloodType == "" {
		http.Error(w, "Todos os campos são obrigatórios", http.StatusBadRequest)
		return
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		http.Error(w, "Idade inválida", http.StatusBadRequest)
		return
	}

	err = utils.InsertPatient(name, age, bloodType)
	if err != nil {
		http.Error(w, "Erro ao criar paciente", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/listPatients", http.StatusSeeOther)
}
