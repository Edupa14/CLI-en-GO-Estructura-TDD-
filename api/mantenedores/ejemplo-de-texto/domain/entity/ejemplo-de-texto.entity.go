package entity

import (
	mssql "github.com/denisenkom/go-mssqldb"
	"time"
)

type EjemploDeTextoEntity struct {
	Id	int64	`db:"id"`
	Nombre	string	`db:"nombre"`
	Codigo	string	`db:"codigo"`
	Habilitado	bool	`db:"habilitado"`
	FechaGuardado	 time.Time 	`db:"fecha_guardado"`
	TransactionUid	 mssql.UniqueIdentifier 	`db:"transaction_uid"`
	
}
