package service

import (
	"gitlab.com/ns-desarrollo-web/erp/ns-core-service"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/mapping"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/repository"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/valueobject"
)

type MotivosExtornoPedidoService struct {
	repo repository.MotivosExtornoPedidoRepository
}

func (e *MotivosExtornoPedidoService) MotivosExtornoPedidoPaginacion(ctx *ns.Context, start int, length int, search string, order string) (res *valueobject.MotivosExtornoPedidoPaginacion, err error) {
	rs, total, filtered, err := e.repo.Listar(ctx, start, length, search, order)
	if err != nil {
		return
	}

	res = mapping.MapMotivosExtornoPedidoPaginación(rs, total, filtered)

	return
}

func (e *MotivosExtornoPedidoService) MotivosExtornoPedidoSuggest(ctx *ns.Context, search, advancedSearch string, columns bool) (res []*valueobject.MotivosExtornoPedidoSuggest, err error) {
	rs, err := e.repo.Buscar(ctx, search, advancedSearch, columns)
	if err != nil {
		return
	}

	res = mapping.MapMotivosExtornoPedidoSuggest(rs)

	return
}

func NewMotivosExtornoPedidoService(MotivosExtornoPedidoRepository repository.MotivosExtornoPedidoRepository) *MotivosExtornoPedidoService{
	return &MotivosExtornoPedidoService{repo: MotivosExtornoPedidoRepository}
}
