package repository

import (
	"database/sql"
	"fmt"

	"github.com/JosephAntony37900/API-1-Multi/Users/domain/entities"
)

type UserClientRepoMySQL struct {
	db *sql.DB
}

func NewUserClientRepoMySQL(db *sql.DB) *UserClientRepoMySQL {
	return &UserClientRepoMySQL{db: db}
}

func (r *UserClientRepoMySQL) Save(user entities.UserClient) error {
	query := "INSERT INTO Admin_Cliente (Id_Admin, Id_Cliente) VALUES (?, ?)"
	_, err := r.db.Exec(query, user.Id_Admin, user.Id_client)
	if err != nil {
		return fmt.Errorf("error insertando UserClient: %w", err)
	}
	return nil
}

func (r *UserClientRepoMySQL) FindByEmail(email string) (*entities.UserClient, error) {
	// Este método no se usa por ahora, pero está definido en la interfaz
	return nil, fmt.Errorf("FindByEmail no está implementado")
}