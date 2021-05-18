package model

import (
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/entity"
	"ns-compras-service/api/utils"
)

type MotivosExtornoPedidoModel struct {
	entity.MotivosExtornoPedidoEntity `db:",squash"`
	Cols                  []utils.SuggestColumns        `db:"cols"`
}
