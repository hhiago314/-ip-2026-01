package handlers

import (
	"net/http"
	"servidorHTTP/app/utils"
	"strconv"
)

func UpdatePatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método inválido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.FormValue("id")
	name := r.FormValue("name")
	ageStr := r.FormValue("age")
	bloodType := r.FormValue("blood_type")

	if idStr == "" || name == "" || ageStr == "" || bloodType == "" {
		http.Error(w, "Todos os campos são obrigatórios", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		http.Error(w, "Idade inválida", http.StatusBadRequest)
		return
	}

	err = utils.UpdatePatient(id, name, age, bloodType)
	if err != nil {
		http.Error(w, "Erro ao atualizar paciente", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/listPatients", http.StatusSeeOther)
}
