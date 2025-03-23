package repository

import (
	"database/sql"
	"errors"
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/repository"
)

type levelReadingRepoMySQL struct {
	db *sql.DB
}

func NewLevelReadingRepoMySQL(db *sql.DB) repository.Level_ReadingRepository {
	return &levelReadingRepoMySQL{db: db}
}

func (repo *levelReadingRepoMySQL) Save(levelReading entities.Level_Reading) error {
	query := "INSERT INTO level_readings (Fecha, Id_Jabon, Nivel_Jabon) VALUES (?, ?, ?)"
	_, err := repo.db.Exec(query, levelReading.Fecha, levelReading.Id_Jabon, levelReading.Nivel_Jabon)
	if err != nil {
		return err
	}
	return nil
}

func (repo *levelReadingRepoMySQL) FindById(id int) (*entities.Level_Reading, error) {
	query := "SELECT Id, Fecha, Id_Jabon, Nivel_Jabon FROM level_readings WHERE Id = ?"
	row := repo.db.QueryRow(query, id)

	var levelReading entities.Level_Reading
	if err := row.Scan(&levelReading.Id, &levelReading.Fecha, &levelReading.Id_Jabon, &levelReading.Nivel_Jabon); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &levelReading, nil
}

func (repo *levelReadingRepoMySQL) GetAll() ([]entities.Level_Reading, error) {
	query := "SELECT Id, Fecha, Id_Jabon, Nivel_Jabon FROM level_readings"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var levelReadings []entities.Level_Reading
	for rows.Next() {
		var levelReading entities.Level_Reading
		if err := rows.Scan(&levelReading.Id, &levelReading.Fecha, &levelReading.Id_Jabon, &levelReading.Nivel_Jabon); err != nil {
			return nil, err
		}
		levelReadings = append(levelReadings, levelReading)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return levelReadings, nil
}
