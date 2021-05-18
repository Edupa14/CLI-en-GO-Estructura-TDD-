package repository

import (
	"gitlab.com/ns-desarrollo-web/erp/ns-core-service"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/model"
)

type MotivosExtornoPedidoRepository interface {
	Listar(ctx *ns.Context, start int, length int, search string, order string) (res []model.MotivosExtornoPedidoModel, total int64, filtered int64, err error)
	Buscar(ctx *ns.Context, search, advancedSearch string, columns bool) (res []model.MotivosExtornoPedidoModel, err error)
	BuscarPorId(ctx *ns.Context, id int64) (res *model.MotivosExtornoPedidoModel, err error)
	Crear(ctx *ns.Context, productoModel *model.MotivosExtornoPedidoModel) (id int64, err error)
	Actualizar(ctx *ns.Context, id int64, productoModel *model.MotivosExtornoPedidoModel) (err error)
	Eliminar(ctx *ns.Context, id int64) (err error)
	HabilitarDeshabilitar(ctx *ns.Context, id int64, status bool) (err error)
}
