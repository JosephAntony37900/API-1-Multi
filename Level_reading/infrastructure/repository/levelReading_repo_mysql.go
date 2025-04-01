package repository

import (
	"database/sql"
	"errors"
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/repository"
	"fmt"
	"time"
	"log"
)

type levelReadingRepoMySQL struct {
	db *sql.DB
}

func NewLevelReadingRepoMySQL(db *sql.DB) repository.Level_ReadingRepository {
	return &levelReadingRepoMySQL{db: db}
}

func (repo *levelReadingRepoMySQL) Save(levelReading entities.Level_Reading) error {
    
    query := "INSERT INTO Lectura_Nivel (Fecha, Id_Jabon, Nivel_Jabon, Codigo_Identificador, Tipo) VALUES (?, ?, ?, ?, ?)"
    _, err := repo.db.Exec(query, levelReading.Fecha, levelReading.Id_Jabon, levelReading.Nivel_Jabon, levelReading.Codigo_Identificador, levelReading.Tipo)
    if err != nil {
        log.Printf("Error ejecutando la consulta SQL: %v", err)
        return err
    }
    return nil
}

func (repo *levelReadingRepoMySQL) FindById(id int) (*entities.Level_Reading, error) {
	query := `
		SELECT ln.Id, ln.Fecha, ln.Id_Jabon, j.Nombre AS Jabon_Nombre, ln.Nivel_Jabon, nj.Descripcion AS Nivel_Texto
		FROM Lectura_Nivel ln
		JOIN Jabon j ON ln.Id_Jabon = j.Id
		JOIN Nivel_Jabon nj ON ln.Nivel_Jabon = nj.Id
		WHERE ln.Id = ?
	`
	row := repo.db.QueryRow(query, id)

	var levelReading entities.Level_Reading
	var fechaString string // Temporal para almacenar la fecha como string

	if err := row.Scan(&levelReading.Id, &fechaString, &levelReading.Id_Jabon, &levelReading.Jabon_Nombre, &levelReading.Nivel_Jabon, &levelReading.Nivel_Texto); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	fecha, err := time.Parse("2006-01-02 15:04:05", fechaString)
	if err != nil {
		return nil, fmt.Errorf("error al parsear la fecha: %w", err)
	}
	levelReading.Fecha = fecha

	return &levelReading, nil
}

func (repo *levelReadingRepoMySQL) GetAll() ([]entities.Level_Reading, error) {
	query := `
		SELECT ln.Id, ln.Fecha, ln.Id_Jabon, j.Nombre AS Jabon_Nombre, ln.Nivel_Jabon, nj.Descripcion AS Nivel_Texto
		FROM Lectura_Nivel ln
		JOIN Jabon j ON ln.Id_Jabon = j.Id
		JOIN Nivel_Jabon nj ON ln.Nivel_Jabon = nj.Id
	`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var levelReadings []entities.Level_Reading
	for rows.Next() {
		var levelReading entities.Level_Reading
		var fechaString string // Temporal para almacenar la fecha como string

		if err := rows.Scan(&levelReading.Id, &fechaString, &levelReading.Id_Jabon, &levelReading.Jabon_Nombre, &levelReading.Nivel_Jabon, &levelReading.Nivel_Texto); err != nil {
			return nil, err
		}

		fecha, err := time.Parse("2006-01-02 15:04:05", fechaString)
		if err != nil {
			return nil, fmt.Errorf("error al parsear la fecha: %w", err)
		}
		levelReading.Fecha = fecha

		levelReadings = append(levelReadings, levelReading)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return levelReadings, nil
}

func (repo *levelReadingRepoMySQL) GetLast() (*entities.Level_Reading, error) {
	query := `
		SELECT ln.Id, ln.Fecha, ln.Id_Jabon, j.Nombre AS Jabon_Nombre, ln.Nivel_Jabon, nj.Descripcion AS Nivel_Texto
		FROM Lectura_Nivel ln
		JOIN Jabon j ON ln.Id_Jabon = j.Id
		JOIN Nivel_Jabon nj ON ln.Nivel_Jabon = nj.Id
		ORDER BY ln.Id DESC LIMIT 1
	`
	row := repo.db.QueryRow(query)

	var levelReading entities.Level_Reading
	var fechaString string

	if err := row.Scan(&levelReading.Id, &fechaString, &levelReading.Id_Jabon, &levelReading.Jabon_Nombre, &levelReading.Nivel_Jabon, &levelReading.Nivel_Texto); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No hay registros en la base de datos
		}
		return nil, err
	}

	// Convertir la fecha a time.Time
	fecha, err := time.Parse("2006-01-02 15:04:05", fechaString)
	if err != nil {
		return nil, fmt.Errorf("error al parsear la fecha: %w", err)
	}
	levelReading.Fecha = fecha

	return &levelReading, nil
}

func (repo *levelReadingRepoMySQL) SaveWithReturnId(levelReading entities.Level_Reading) (int, error) {
    query := "INSERT INTO Lectura_Nivel (Fecha, Id_Jabon, Nivel_Jabon, Codigo_Identificador, Tipo) VALUES (?, ?, ?, ?, ?)"
    result, err := repo.db.Exec(query, levelReading.Fecha, levelReading.Id_Jabon, levelReading.Nivel_Jabon, levelReading.Codigo_Identificador, levelReading.Tipo)
    if err != nil {
        return 0, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return int(id), nil
}

func (r *levelReadingRepoMySQL) FindUserAdminByJabon(idJabon int) (int, error) {
	query := `SELECT Id_Usuario_Admin FROM Jabon WHERE Id = ?`
	var idUserAdmin int
	err := r.db.QueryRow(query, idJabon).Scan(&idUserAdmin)
	if err != nil {
		return 0, fmt.Errorf("error obteniendo el usuario administrador: %w", err)
	}
	return idUserAdmin, nil
}
