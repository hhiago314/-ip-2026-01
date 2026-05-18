package utils

import (
	"log"
)

type Patient struct {
	ID        int
	Name      string
	Age       int
	BloodType string
	CreatedAt string
}

func InsertPatient(name string, age int, bloodType string) error {
	query := `INSERT INTO patients (name, age, blood_type) VALUES ($1, $2, $3)`
	_, err := DB.Exec(query, name, age, bloodType)
	if err != nil {
		log.Printf("Erro ao inserir paciente no banco de dados: %v", err)
		return err
	}
	log.Println("Paciente inserido com sucesso!")
	return nil
}

func GetPatients() ([]Patient, error) {
	query := `SELECT id, name, age, blood_type, created_at FROM patients ORDER BY id`
	rows, err := DB.Query(query)
	if err != nil {
		log.Printf("Erro ao buscar pacientes no banco de dados: %v", err)
		return nil, err
	}
	defer rows.Close()

	patients := []Patient{}
	for rows.Next() {
		var p Patient
		err := rows.Scan(&p.ID, &p.Name, &p.Age, &p.BloodType, &p.CreatedAt)
		if err != nil {
			log.Printf("Erro ao ler paciente do banco de dados: %v", err)
			return nil, err
		}
		patients = append(patients, p)
	}
	return patients, nil
}

func UpdatePatient(id int, name string, age int, bloodType string) error {
	query := `UPDATE patients SET name = $1, age = $2, blood_type = $3 WHERE id = $4`
	_, err := DB.Exec(query, name, age, bloodType, id)
	if err != nil {
		log.Printf("Erro ao atualizar paciente no banco de dados: %v", err)
		return err
	}
	log.Println("Paciente atualizado com sucesso!")
	return nil
}

func DeletePatient(id int) error {
	query := `DELETE FROM patients WHERE id = $1`
	_, err := DB.Exec(query, id)
	if err != nil {
		log.Printf("Erro ao apagar paciente do banco de dados: %v", err)
		return err
	}
	log.Println("Paciente apagado com sucesso!")
	return nil
}
