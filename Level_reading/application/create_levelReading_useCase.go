package application

import (
	"fmt"
	"time"

	"github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/repository"
)

type CreateLevelReading struct {
	repo repository.Level_ReadingRepository
}

func NewCreateLevelReading(repo repository.Level_ReadingRepository) *CreateLevelReading {
	return &CreateLevelReading{repo: repo}
}

func (clr *CreateLevelReading) Run(Fecha int64, Id_Jabon int, Nivel float64, Codigo_Identificador string, Tipo bool) error {
    // Convertir int64 a time.Time
    fecha := time.Unix(Fecha, 0)

    // Determinar el nivel basado en el valor de Nivel_Jabon
    var Nivel_Jabon int
    switch {
    case Nivel >= 85.00 && Nivel <= 100.00:
        Nivel_Jabon = 1 // Lleno
    case Nivel >= 60.00 && Nivel < 85.00:
        Nivel_Jabon = 2 // Casi lleno
    case Nivel >= 35.00 && Nivel < 60.00:
        Nivel_Jabon = 3 // Medio
    case Nivel >= 5.00 && Nivel < 35.00:
        Nivel_Jabon = 4 // Bajo
    case Nivel >= 0.00 && Nivel < 5.00:
        Nivel_Jabon = 5 // Vacío
    default:
        return fmt.Errorf("valor de nivel inválido: %v", Nivel)
    }

    levelReading := entities.Level_Reading{
        Fecha:              fecha,
        Id_Jabon:           Id_Jabon,
        Nivel_Jabon:        Nivel_Jabon,
        Codigo_Identificador: Codigo_Identificador,
        Tipo:               Tipo, 
    }

    if err := clr.repo.Save(levelReading); err != nil {
        return fmt.Errorf("error guardando el nivel de lectura: %w", err)
    }

    return nil
}

func (clr *CreateLevelReading) RunWithReturnId(Fecha int64, Id_Jabon int, Nivel float64, Codigo_Identificador string, Tipo bool) (int, error) {
    fecha := time.Unix(Fecha, 0)

    levelReading := entities.Level_Reading{
        Fecha:              fecha,
        Id_Jabon:           Id_Jabon,
        Nivel_Jabon:        int(Nivel),
        Codigo_Identificador: Codigo_Identificador,
        Tipo:               Tipo, 
    }

    id, err := clr.repo.SaveWithReturnId(levelReading)
    if err != nil {
        return 0, fmt.Errorf("error guardando el nivel de lectura: %w", err)
    }

    return id, nil
}