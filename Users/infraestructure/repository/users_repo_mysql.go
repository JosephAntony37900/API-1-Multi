package repository

import (
	"database/sql"
	"fmt"

	"github.com/JosephAntony37900/API-1-Multi/Users/domain/entities"
)

type UserRepoMySQL struct {
	db *sql.DB
}

func NewCreateUserRepoMySQL(db *sql.DB) *UserRepoMySQL {
	return &UserRepoMySQL{db: db}
}

func (r *UserRepoMySQL) Save(User entities.Users) error {
	query := "INSERT INTO Usuarios (Nombre, Email, Contraseña) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, User.Nombre, User.Email, User.Contraseña)
	if err != nil {
		return fmt.Errorf("error insertando User: %w", err)
	}
	return nil
}

func (r *UserRepoMySQL) FindByID(id int) (*entities.Users, error) {
	query := "SELECT Id, Nombre, Email FROM Usuarios WHERE Id = ?"
	row := r.db.QueryRow(query, id)

	var User entities.Users
	if err := row.Scan(&User.Id, &User.Nombre, &User.Email); err != nil {
		return nil, fmt.Errorf("error buscando el User: %w", err)
	}
	return &User, nil
}

func (r *UserRepoMySQL) FindAll() ([]entities.Users, error) {
	query := "SELECT Id, Nombre, Email, Contraseña FROM Usuarios"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error buscando los Users: %w", err)
	}
	defer rows.Close()

	var Users []entities.Users
	for rows.Next() {
		var User entities.Users
		if err := rows.Scan(&User.Id, &User.Nombre, &User.Email, &User.Contraseña); err != nil {
			return nil, err
		}
		Users = append(Users, User)
	}
	return Users, nil
}

func (r *UserRepoMySQL) Update(User entities.Users) error {
	query := "UPDATE Usuarios SET Nombre = ?, Email = ?, Contraseña = ? WHERE Id = ?"
	_, err := r.db.Exec(query, User.Nombre, User.Email, User.Contraseña, User.Id)
	if err != nil {
		return fmt.Errorf("error actualizando User: %w", err)
	}
	return nil
}

func (r *UserRepoMySQL) Delete(id int) error {
	query := "DELETE FROM Usuarios WHERE Id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error eliminando User: %w", err)
	}
	return nil
}

func (r *UserRepoMySQL) FindByEmail(email string) (*entities.Users, error) {
	query := "SELECT Id, Nombre, Email, Contraseña FROM Usuarios WHERE Email = ?"
	row := r.db.QueryRow(query, email)

	var user entities.Users
	if err := row.Scan(&user.Id, &user.Nombre, &user.Email, &user.Contraseña); err != nil {
		return nil, fmt.Errorf("error buscando el usuario: %w", err)
	}
	return &user, nil
}
