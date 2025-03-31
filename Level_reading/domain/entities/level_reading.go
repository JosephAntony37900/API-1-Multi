package entities

import "time"


type Level_Reading struct{
	Id int
	Fecha time.Time
	Id_Jabon int
	Nivel_Jabon int
	Jabon_Nombre string
	Nivel_Texto  string
	Codigo_Identificador string
}