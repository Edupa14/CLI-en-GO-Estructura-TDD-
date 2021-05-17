package dto

type EjemploDeTextoRequest struct {
	Id	int64	`json:"id"`
	Nombre	string	`json:"nombre"`
	Codigo	string	`json:"codigo"`
	Habilitado	bool	`json:"habilitado"`
	
}
