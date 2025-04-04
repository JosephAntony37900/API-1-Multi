package service

import (
	"os"

	adapters "github.com/JosephAntony37900/API-1-Multi/Users/infraestructure/adapters"
	"github.com/JosephAntony37900/API-1-Multi/Users/domain/services"
)

// Inicializar el servicio de BCrypt
func InitBcryptService() services.IBcrypService {
	return adapters.NewBcrypt()
}

// Inicializar el Token Manager
func InitTokenManager() services.TokenManager {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic("JWT_SECRET no está configurado en las variables de entorno")
	}
	return &adapters.JWTManager{SecretKey: jwtSecret}
}