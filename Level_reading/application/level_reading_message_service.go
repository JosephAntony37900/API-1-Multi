package application

import (
	"log"
	"fmt"
	"time"

	_"github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/repository"
)

type LevelReadingMessageService struct {
	repo          repository.Level_ReadingRepository
	createUseCase *CreateLevelReading
}

func NewLevelReadingMessageService(repo repository.Level_ReadingRepository, createUseCase *CreateLevelReading) *LevelReadingMessageService {
	return &LevelReadingMessageService{
		repo:          repo,
		createUseCase: createUseCase,
	}
}

func (s *LevelReadingMessageService) ProcessMessage(level float64, idJabon int) error {
	// Obtener el último nivel de lectura
	lastLevel, err := s.repo.GetLast()
	if err != nil {
		return fmt.Errorf("error al obtener el último nivel de lectura: %w", err)
	}

	// Determinar el estado del nivel del mensaje recibido
	var nivelEstado int
	switch {
	case level >= 85.00 && level <= 100.00:
		nivelEstado = 1 // Lleno
	case level >= 60.00 && level < 85.00:
		nivelEstado = 2 // Casi lleno
	case level >= 35.00 && level < 60.00:
		nivelEstado = 3 // Medio
	case level >= 5.00 && level < 35.00:
		nivelEstado = 4 // Bajo
	case level >= 0.00 && level < 5.00:
		nivelEstado = 5 // Vacío
	default:
		log.Printf("Nivel inválido recibido: %v", level)
		return nil // Ignorar niveles inválidos
	}

	// Si no hay un nivel previo o el estado es diferente, creamos un nuevo nivel
	if lastLevel == nil || lastLevel.Nivel_Jabon != nivelEstado {
		log.Println("Creando un nuevo nivel de lectura...")
		err = s.createUseCase.Run(time.Now().Unix(), idJabon, level)
		if err != nil {
			return fmt.Errorf("error al crear un nuevo nivel de lectura: %w", err)
		}
		log.Println("Nuevo nivel de lectura creado")
	} else {
		log.Println("El nivel es igual al último, ignorando el mensaje")
	}

	return nil
}