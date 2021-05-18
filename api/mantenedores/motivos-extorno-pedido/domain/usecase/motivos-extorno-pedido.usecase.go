package usecase

import (
	"gitlab.com/ns-desarrollo-web/erp/ns-core-service"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/dto"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/valueobject"
)

type MotivosExtornoPedidoUseCase interface {
	Listar(ctx *ns.Context, start int, length int, search string, order string) (res *valueobject.MotivosExtornoPedidoPaginacion, err error)
	Buscar(ctx *ns.Context, ssearch, advancedSearch string, columns bool) (res []*valueobject.MotivosExtornoPedidoSuggest, err error)
	BuscarPorId(ctx *ns.Context, id int64) (res * valueobject.MotivosExtornoPedidoResponse, err error)
	Crear(ctx *ns.Context, request *dto.MotivosExtornoPedidoRequest) (res *valueobject.MotivosExtornoPedidoResponse, err error)
	Actualizar(ctx *ns.Context, id int64, request *dto.MotivosExtornoPedidoRequest) (res *valueobject.MotivosExtornoPedidoResponse, err error)
	Eliminar(ctx *ns.Context, id int64) (err error)
	Habilitar(ctx *ns.Context, id int64) (err error)
	Deshabilitar(ctx *ns.Context, id int64) (err error)
}
