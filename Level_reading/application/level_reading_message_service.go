package application

import (
	"fmt"
	"log"
	"time"

	_ "github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/entities"
	messagingmq "github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/messaging_MQ"
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/repository"
	_"github.com/JosephAntony37900/API-1-Multi/helpers"
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

func (s *LevelReadingMessageService) ProcessMessage(level float64, idJabon int, codigoIdentificador string, tipo bool) error {

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

    // Si no hay un nivel previo o el estado es diferente, creamos un nuevo nivel
    var newLevelId int
    if lastLevel == nil || lastLevel.Nivel_Jabon != nivelEstado {
        log.Println("Creando un nuevo nivel de lectura...")
        newLevelId, err = s.createUseCase.RunWithReturnId(time.Now().Unix(), idJabon, float64(nivelEstado), codigoIdentificador, tipo) 
        if err != nil {
            return fmt.Errorf("error al crear un nuevo nivel de lectura: %w", err)
        }

        createdLevel, err := s.repo.FindById(newLevelId)
        if err != nil || createdLevel == nil {
            return fmt.Errorf("error confirmando el nivel de lectura creado con ID: %d", newLevelId)
        }
        log.Printf("Nivel de lectura creado y confirmado con ID: %d", newLevelId)

        err = s.PublishAlertIfNecessary(nivelEstado, newLevelId, codigoIdentificador, tipo) 
        if err != nil {
            log.Printf("Error al intentar publicar alerta: %v", err)
        }
    } else {
        log.Println("El nivel es igual al último, ignorando el mensaje")
    }

    return nil
}

func (s *LevelReadingMessageService) PublishAlertIfNecessary(nivelEstado int, idLectura int, codigoIdentificador string, tipo bool) error {
    if nivelEstado != 4 && nivelEstado != 5 { // Publicar solo si el nivel es "Bajo" o "Vacío"
        s.lastAlertSent = 0 // Resetear el estado de alerta cuando cambia
        return nil
    }

    if s.lastAlertSent == nivelEstado {
        log.Println("Alerta ya enviada previamente, no se publicará nuevamente.")
        return nil
    }

    estado := "Pendiente" 

    err := s.publisher.Publish(estado, idLectura, codigoIdentificador, tipo, "sensor.alerta")
    if err != nil {
        log.Printf("Error al publicar el mensaje: %v", err)
        return err
    }

    log.Printf("Mensaje de alerta publicado: Estado=%s, IdLectura=%d, CodigoIdentificador=%s, Tipo=%t", estado, idLectura, codigoIdentificador, tipo)
    s.lastAlertSent = nivelEstado
    return nil
}