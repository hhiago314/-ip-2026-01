package handlers

import (
	"net/http"
	"servidorHTTP/app/utils"
	"strconv"
)

func DeletePatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método inválido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.FormValue("id")
	if idStr == "" {
		http.Error(w, "ID do paciente é obrigatório", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = utils.DeletePatient(id)
	if err != nil {
		http.Error(w, "Erro ao apagar paciente", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/listPatients", http.StatusSeeOther)
}
