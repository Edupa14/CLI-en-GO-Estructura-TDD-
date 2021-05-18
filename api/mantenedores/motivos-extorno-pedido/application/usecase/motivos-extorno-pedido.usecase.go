package usecase

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gitlab.com/ns-desarrollo-web/erp/ns-core-service"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/application/service"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/dto"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/mapping"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/repository"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/usecase"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/valueobject"
)

type MotivosExtornoPedidoUseCase struct {
	repo repository.MotivosExtornoPedidoRepository
	serv *service.MotivosExtornoPedidoService
}

func (p MotivosExtornoPedidoUseCase) Listar(ctx *ns.Context, start int, length int, search string, order string) (res *valueobject.MotivosExtornoPedidoPaginacion, err error) {
	res, err = p.serv.MotivosExtornoPedidoPaginacion(ctx, start, length, search, order)
	if err != nil {
		log.Error("Error al Listar [MotivosExtornoPedido]: ", err)
		return
	}

	return
}

func (p MotivosExtornoPedidoUseCase) Buscar(ctx *ns.Context, search, advancedSearch string, columns bool) (res []*valueobject.MotivosExtornoPedidoSuggest, err error) {
	res, err = p.serv.MotivosExtornoPedidoSuggest(ctx, search, advancedSearch, columns)
	if err != nil {
		log.Error("Error al Buscar [MotivosExtornoPedido]: ", err)
		return
	}

	return
}

func (p MotivosExtornoPedidoUseCase) BuscarPorId(ctx *ns.Context, id int64) (res *valueobject.MotivosExtornoPedidoResponse, err error) {
	rs, err := p.repo.BuscarPorId(ctx, id)
	if err != nil {
		log.Error("Error al Buscar Por Id [MotivosExtornoPedido]: ", err)
		return
	}

	res = mapping.MapMotivosExtornoPedidoModelToVO(rs)

	return
}

func (p MotivosExtornoPedidoUseCase) Crear(ctx *ns.Context, request *dto.MotivosExtornoPedidoRequest) (res *valueobject.MotivosExtornoPedidoResponse, err error) {
	mod := mapping.MapMotivosExtornoPedidoDTOTOModel(request)
	id, err := p.repo.Crear(ctx, mod)
	if err != nil {
		log.Error("Error al Crear [MotivosExtornoPedido]: ", err)
		return
	}

	rs, err := p.repo.BuscarPorId(ctx, id)
	if err != nil {
		log.Error("Error al Crear Buscar Por Id [MotivosExtornoPedido]: ", err)
		return
	}

	res = mapping.MapMotivosExtornoPedidoModelToVO(rs)

	return
}

func (p MotivosExtornoPedidoUseCase) Actualizar(ctx *ns.Context, id int64, request *dto.MotivosExtornoPedidoRequest) (res *valueobject.MotivosExtornoPedidoResponse, err error) {
	mod := mapping.MapMotivosExtornoPedidoDTOTOModel(request)
	err = p.repo.Actualizar(ctx, id, mod)
	if err != nil {
		log.Error("Error al Actualizar [MotivosExtornoPedido]: ", err)
		return
	}

	rs, err := p.repo.BuscarPorId(ctx, id)
	if err != nil {
		log.Error("Error al Actualizar Buscar Por Id [MotivosExtornoPedido]: ", err)
		return
	}

	res = mapping.MapMotivosExtornoPedidoModelToVO(rs)

	return
}

func (p MotivosExtornoPedidoUseCase) Eliminar(ctx *ns.Context, id int64) (err error) {
	err = p.repo.Eliminar(ctx, id)
	if err != nil {
		log.Error("Error al Eliminar [MotivosExtornoPedido]: ", err)
		return
	}

	return
}

func (p MotivosExtornoPedidoUseCase) Habilitar(ctx *ns.Context, id int64) (err error) {
	rs, err := p.repo.BuscarPorId(ctx, id)
	if err != nil {
		log.Error("Error al Habilitar Buscar Por Id [MotivosExtornoPedido]: ", err)
		return
	}

	if rs.Habilitado{
		err = errors.New("El registro ya se encuentra habilitado")
		return
	}

	err = p.repo.HabilitarDeshabilitar(ctx, id, true)
	if err != nil {
		log.Error("Error al Habilitar [MotivosExtornoPedido]: ", err)
		return
	}

	return
}

func (p MotivosExtornoPedidoUseCase) Deshabilitar(ctx *ns.Context, id int64) (err error) {
	rs, err := p.repo.BuscarPorId(ctx, id)
	if err != nil {
		log.Error("Error al Deshabilitar Buscar Por Id [MotivosExtornoPedido]: ", err)
		return
	}

	if !rs.Habilitado{
		err = errors.New("El registro ya se encuentra deshabilitado")
		return
	}

	err = p.repo.HabilitarDeshabilitar(ctx, id, false)
	if err != nil {
		log.Error("Error al Deshabilitar [MotivosExtornoPedido]: ", err)
		return
	}

	return
}

func NewMotivosExtornoPedidoUseCase(MotivosExtornoPedidoRepository repository.MotivosExtornoPedidoRepository, MotivosExtornoPedidoService *service.MotivosExtornoPedidoService) usecase.MotivosExtornoPedidoUseCase {
	return MotivosExtornoPedidoUseCase{
		repo: MotivosExtornoPedidoRepository,
		serv: MotivosExtornoPedidoService,
	}
}
