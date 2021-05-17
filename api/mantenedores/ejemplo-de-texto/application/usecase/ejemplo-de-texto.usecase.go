package usecase

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gitlab.com/ns-desarrollo-web/erp/ns-core-service"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/application/service"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/dto"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/mapping"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/repository"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/usecase"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/valueobject"
)

type EjemploDeTextoUseCase struct {
	repo repository.EjemploDeTextoRepository
	serv *service.EjemploDeTextoService
}

func (p EjemploDeTextoUseCase) Listar(ctx *ns.Context, start int, length int, search string, order string) (res *valueobject.EjemploDeTextoPaginacion, err error) {
	res, err = p.serv.EjemploDeTextoPaginacion(ctx, start, length, search, order)
	if err != nil {
		log.Error("Error al Listar [EjemploDeTexto]: ", err)
		return
	}

	return
}

func (p EjemploDeTextoUseCase) Buscar(ctx *ns.Context, search, advancedSearch string, columns bool) (res []*valueobject.EjemploDeTextoSuggest, err error) {
	res, err = p.serv.EjemploDeTextoSuggest(ctx, search, advancedSearch, columns)
	if err != nil {
		log.Error("Error al Buscar [EjemploDeTexto]: ", err)
		return
	}

	return
}

func (p EjemploDeTextoUseCase) BuscarPorId(ctx *ns.Context, id int64) (res *valueobject.EjemploDeTextoResponse, err error) {
	rs, err := p.repo.BuscarPorId(ctx, id)
	if err != nil {
		log.Error("Error al Buscar Por Id [EjemploDeTexto]: ", err)
		return
	}

	res = mapping.MapEjemploDeTextoModelToVO(rs)

	return
}

func (p EjemploDeTextoUseCase) Crear(ctx *ns.Context, request *dto.EjemploDeTextoRequest) (res *valueobject.EjemploDeTextoResponse, err error) {
	mod := mapping.MapEjemploDeTextoDTOTOModel(request)
	id, err := p.repo.Crear(ctx, mod)
	if err != nil {
		log.Error("Error al Crear [EjemploDeTexto]: ", err)
		return
	}

	rs, err := p.repo.BuscarPorId(ctx, id)
	if err != nil {
		log.Error("Error al Crear Buscar Por Id [EjemploDeTexto]: ", err)
		return
	}

	res = mapping.MapEjemploDeTextoModelToVO(rs)

	return
}

func (p EjemploDeTextoUseCase) Actualizar(ctx *ns.Context, id int64, request *dto.EjemploDeTextoRequest) (res *valueobject.EjemploDeTextoResponse, err error) {
	mod := mapping.MapEjemploDeTextoDTOTOModel(request)
	err = p.repo.Actualizar(ctx, id, mod)
	if err != nil {
		log.Error("Error al Actualizar [EjemploDeTexto]: ", err)
		return
	}

	rs, err := p.repo.BuscarPorId(ctx, id)
	if err != nil {
		log.Error("Error al Actualizar Buscar Por Id [EjemploDeTexto]: ", err)
		return
	}

	res = mapping.MapEjemploDeTextoModelToVO(rs)

	return
}

func (p EjemploDeTextoUseCase) Eliminar(ctx *ns.Context, id int64) (err error) {
	err = p.repo.Eliminar(ctx, id)
	if err != nil {
		log.Error("Error al Eliminar [EjemploDeTexto]: ", err)
		return
	}

	return
}

func (p EjemploDeTextoUseCase) Habilitar(ctx *ns.Context, id int64) (err error) {
	rs, err := p.repo.BuscarPorId(ctx, id)
	if err != nil {
		log.Error("Error al Habilitar Buscar Por Id [EjemploDeTexto]: ", err)
		return
	}

	if rs.Habilitado{
		err = errors.New("El registro ya se encuentra habilitado")
		return
	}

	err = p.repo.HabilitarDeshabilitar(ctx, id, true)
	if err != nil {
		log.Error("Error al Habilitar [EjemploDeTexto]: ", err)
		return
	}

	return
}

func (p EjemploDeTextoUseCase) Deshabilitar(ctx *ns.Context, id int64) (err error) {
	rs, err := p.repo.BuscarPorId(ctx, id)
	if err != nil {
		log.Error("Error al Deshabilitar Buscar Por Id [EjemploDeTexto]: ", err)
		return
	}

	if !rs.Habilitado{
		err = errors.New("El registro ya se encuentra deshabilitado")
		return
	}

	err = p.repo.HabilitarDeshabilitar(ctx, id, false)
	if err != nil {
		log.Error("Error al Deshabilitar [EjemploDeTexto]: ", err)
		return
	}

	return
}

func NewEjemploDeTextoUseCase(EjemploDeTextoRepository repository.EjemploDeTextoRepository, EjemploDeTextoService *service.EjemploDeTextoService) usecase.EjemploDeTextoUseCase {
	return EjemploDeTextoUseCase{
		repo: EjemploDeTextoRepository,
		serv: EjemploDeTextoService,
	}
}
