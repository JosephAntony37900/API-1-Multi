package helpers

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword encripta la contraseña
func HashPassword(password string) (string, error) {
	fmt.Println("Contraseña original sin Hash en el hasheo", password)
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println("Contraseña hasheada: ", string(hashed), "Longitud:", len(hashed))
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// ComparePassword verifica si la contraseña es correcta
func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(strings.TrimSpace(password)))
	fmt.Println("Error en contra: ", err)
	return err == nil
}
