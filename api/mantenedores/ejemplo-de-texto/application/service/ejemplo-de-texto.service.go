package service

import (
	"gitlab.com/ns-desarrollo-web/erp/ns-core-service"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/mapping"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/repository"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/valueobject"
)

type EjemploDeTextoService struct {
	repo repository.EjemploDeTextoRepository
}

func (e *EjemploDeTextoService) EjemploDeTextoPaginacion(ctx *ns.Context, start int, length int, search string, order string) (res *valueobject.EjemploDeTextoPaginacion, err error) {
	rs, total, filtered, err := e.repo.Listar(ctx, start, length, search, order)
	if err != nil {
		return
	}

	res = mapping.MapEjemploDeTextoPaginación(rs, total, filtered)

	return
}

func (e *EjemploDeTextoService) EjemploDeTextoSuggest(ctx *ns.Context, search, advancedSearch string, columns bool) (res []*valueobject.EjemploDeTextoSuggest, err error) {
	rs, err := e.repo.Buscar(ctx, search, advancedSearch, columns)
	if err != nil {
		return
	}

	res = mapping.MapEjemploDeTextoSuggest(rs)

	return
}

func NewEjemploDeTextoService(EjemploDeTextoRepository repository.EjemploDeTextoRepository) *EjemploDeTextoService{
	return &EjemploDeTextoService{repo: EjemploDeTextoRepository}
}
