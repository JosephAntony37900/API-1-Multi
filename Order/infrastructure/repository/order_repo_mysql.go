package repository

import (
	"database/sql"
	"fmt"

	"github.com/JosephAntony37900/API-1-Multi/Order/domain/entities"
)

type orderRepoMySQL struct {
	db *sql.DB
}

func NewOrderRepoMySQL(db *sql.DB) *orderRepoMySQL {
	return &orderRepoMySQL{db: db}
}

func (repo *orderRepoMySQL) Save(order entities.Order) error {
	query := `INSERT INTO Orden (Id_Jabon, Cantidad, Estado, Costo, Codigo_Identificador, Tipo) 
			  VALUES (?, ?, ?, ?, ?, ?)`
	_, err := repo.db.Exec(query, order.Id_Jabon, order.Cantidad, order.Estado, order.Costo, order.Codigo_Identificador, order.Tipo)
	if err != nil {
		return fmt.Errorf("error guardando la orden: %w", err)
	}
	return nil
}

func (repo *orderRepoMySQL) Update(order entities.Order) error {
	query := `UPDATE Orden 
			  SET Id_Jabon = ?, Cantidad = ?, Estado = ?, Costo = ?, Tipo = ?
			  WHERE Codigo_Identificador = ?`
	_, err := repo.db.Exec(query, order.Id_Jabon, order.Cantidad, order.Estado, order.Costo, order.Tipo, order.Codigo_Identificador)
	if err != nil {
		return fmt.Errorf("error actualizando la orden: %w", err)
	}
	return nil
}

func (repo *orderRepoMySQL) FindById(Codigo_Identificador string) (*entities.Order, error) {
	query := `SELECT Id, Id_Jabon, Cantidad, Estado, Costo, Codigo_Identificador, Tipo 
			  FROM Orden 
			  WHERE Codigo_Identificador = ?`
	row := repo.db.QueryRow(query, Codigo_Identificador)

	var order entities.Order
	err := row.Scan(&order.Id, &order.Id_Jabon, &order.Cantidad, &order.Estado, &order.Costo, &order.Codigo_Identificador, &order.Tipo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no se encontró ninguna orden con el Código_Identificador: %s", Codigo_Identificador)
		}
		return nil, fmt.Errorf("error buscando la orden por Código_Identificador: %w", err)
	}
	return &order, nil
}