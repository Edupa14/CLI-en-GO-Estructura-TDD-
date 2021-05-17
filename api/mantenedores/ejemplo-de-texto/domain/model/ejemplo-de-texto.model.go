package model

import (
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/entity"
	"ns-compras-service/api/utils"
)

type EjemploDeTextoModel struct {
	entity.EjemploDeTextoEntity `db:",squash"`
	Cols                  []utils.SuggestColumns        `db:"cols"`
}
