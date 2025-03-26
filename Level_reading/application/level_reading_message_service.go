package application

import (
	"fmt"
	"log"
	"time"

	_ "github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/entities"
	messagingmq "github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/messaging_MQ"
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/repository"
)

type LevelReadingMessageService struct {
	repo          repository.Level_ReadingRepository
	createUseCase *CreateLevelReading
	publisher     messagingmq.MessagePublisher
	lastAlertSent int
}

func NewLevelReadingMessageService(repo repository.Level_ReadingRepository, createUseCase *CreateLevelReading, publisher messagingmq.MessagePublisher) *LevelReadingMessageService {
	return &LevelReadingMessageService{
		repo:          repo,
		createUseCase: createUseCase,
		publisher:     publisher,
		lastAlertSent: 0,
	}
}

func (s *LevelReadingMessageService) ProcessMessage(level float64, idJabon int) error {
	lastLevel, err := s.repo.GetLast()
	if err != nil {
		return fmt.Errorf("error al obtener el último nivel de lectura: %w", err)
	}

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

	// Publicar alerta si es necesario
	err = s.PublishAlertIfNecessary(nivelEstado)
	if err != nil {
		log.Printf("Error al intentar publicar alerta: %v", err)
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

func (s *LevelReadingMessageService) PublishAlertIfNecessary(nivelEstado int) error {
	// Publicar un mensaje solo si el estado es bajo (4) o vacío (5)
	if nivelEstado != 4 && nivelEstado != 5 {
		s.lastAlertSent = 0 // Resetear el estado de alerta cuando cambia
		return nil
	}

	// Evitar duplicados: solo enviar si el estado es diferente al último enviado
	if s.lastAlertSent == nivelEstado {
		log.Println("Alerta ya enviada previamente, no se publicará nuevamente.")
		return nil
	}

	// Preparar mensaje de alerta
	estadoTexto := map[int]string{
		4: "Bajo",
		5: "Vacío",
	}[nivelEstado]

	message := fmt.Sprintf("Alerta: el nivel está %s", estadoTexto)
	err := s.publisher.Publish(message, "sensor.alerta")
	if err != nil {
		return fmt.Errorf("error al publicar la alerta: %w", err)
	}

	log.Printf("Mensaje de alerta publicado: %s", message)
	s.lastAlertSent = nivelEstado // Actualizar el estado enviado
	return nil
}