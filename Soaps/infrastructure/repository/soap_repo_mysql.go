package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/JosephAntony37900/API-1-Multi/Soaps/domain/entities"
)

type SoapRepoMySQL struct {
	db *sql.DB
}

// Constructor del repositorio
func NewSoapRepoMySQL(db *sql.DB) *SoapRepoMySQL {
	return &SoapRepoMySQL{db: db}
}

func (r *SoapRepoMySQL) Save(soap entities.Soaps) error {
	query := `
		INSERT INTO Jabon (Nombre, Marca, Tipo, Precio, Densidad, Id_Usuario_Admin )
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.Exec(query, soap.Nombre, soap.Marca, soap.Tipo, soap.Precio, soap.Densidad, soap.Id_Usuario_Admin)
	if err != nil {
		return fmt.Errorf("error al guardar el jabón en la BD: %w", err)
	}

	// Obtener el ID del jabón insertado
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener el ID del jabón: %w", err)
	}
	soap.Id = int(id)

	log.Printf("✅ Jabón guardado en la BD: %+v", soap)
	return nil
}

func (r *SoapRepoMySQL) FindById(id int) (*entities.Soaps, error) {
	query := `
		SELECT Id, Nombre, Marca, Tipo, Precio, Densidad, Id_Usuario_Admin 
		FROM Jabon
		WHERE Id = ?
	`

	row := r.db.QueryRow(query, id)
	var soap entities.Soaps
	if err := row.Scan(&soap.Id, &soap.Nombre, &soap.Marca, &soap.Tipo, &soap.Precio, &soap.Densidad, &soap.Id_Usuario_Admin); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("jabón no encontrado con ID %d", id)
		}
		return nil, fmt.Errorf("error al obtener el jabón con ID %d: %w", id, err)
	}

	return &soap, nil
}

func (r *SoapRepoMySQL) GetAll() ([]entities.Soaps, error) {
	query := `
		SELECT Id, Nombre, Marca, Tipo, Precio, Densidad
		FROM Jabon
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener todos los jabones: %w", err)
	}
	defer rows.Close()

	var soaps []entities.Soaps
	for rows.Next() {
		var soap entities.Soaps
		if err := rows.Scan(&soap.Id, &soap.Nombre, &soap.Marca, &soap.Tipo, &soap.Precio, &soap.Densidad); err != nil {
			return nil, fmt.Errorf("error al escanear jabones: %w", err)
		}
		soaps = append(soaps, soap)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error en la iteración de jabones: %w", err)
	}

	return soaps, nil
}

func (r *SoapRepoMySQL) Update(soap entities.Soaps) error {
	query := `
		UPDATE Jabon
		SET Nombre = ?, Marca = ?, Tipo = ?, Precio = ?, Densidad = ?
		WHERE Id = ?
	`

	result, err := r.db.Exec(query, soap.Nombre, soap.Marca, soap.Tipo, soap.Precio, soap.Densidad, soap.Id)
	if err != nil {
		return fmt.Errorf("error al actualizar el jabón %d: %w", soap.Id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener filas afectadas: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró el jabón con ID %d", soap.Id)
	}

	log.Printf("✅ Jabón actualizado en la BD: %+v", soap)
	return nil
}

func (r *SoapRepoMySQL) Delete(id int) error {
	query := `
		DELETE FROM Jabon
		WHERE Id = ?
	`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar el jabón %d: %w", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener filas afectadas: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró el jabón con ID %d", id)
	}

	log.Printf("✅ Jabón eliminado correctamente: %d", id)
	return nil
}

func (r *SoapRepoMySQL) FindByAdminId(adminId int) ([]entities.Soaps, error) {
	query := `
		SELECT Id, Nombre, Marca, Tipo, Precio, Densidad, Id_Usuario_Admin
		FROM Jabon
		WHERE Id_Usuario_Admin = ?
	`

	rows, err := r.db.Query(query, adminId)
	if err != nil {
		return nil, fmt.Errorf("error al obtener los jabones del administrador con ID %d: %w", adminId, err)
	}
	defer rows.Close()

	var soaps []entities.Soaps
	for rows.Next() {
		var soap entities.Soaps
		if err := rows.Scan(&soap.Id, &soap.Nombre, &soap.Marca, &soap.Tipo, &soap.Precio, &soap.Densidad, &soap.Id_Usuario_Admin); err != nil {
			return nil, fmt.Errorf("error al escanear los jabones: %w", err)
		}
		soaps = append(soaps, soap)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error en la iteración de jabones: %w", err)
	}

	log.Printf("✅ Jabones obtenidos para el administrador con ID %d: %+v", adminId, soaps)
	return soaps, nil
}
