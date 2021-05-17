package usecase

import (
	"gitlab.com/ns-desarrollo-web/erp/ns-core-service"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/dto"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/valueobject"
)

type EjemploDeTextoUseCase interface {
	Listar(ctx *ns.Context, start int, length int, search string, order string) (res *valueobject.EjemploDeTextoPaginacion, err error)
	Buscar(ctx *ns.Context, ssearch, advancedSearch string, columns bool) (res []*valueobject.EjemploDeTextoSuggest, err error)
	BuscarPorId(ctx *ns.Context, id int64) (res * valueobject.EjemploDeTextoResponse, err error)
	Crear(ctx *ns.Context, request *dto.EjemploDeTextoRequest) (res *valueobject.EjemploDeTextoResponse, err error)
	Actualizar(ctx *ns.Context, id int64, request *dto.EjemploDeTextoRequest) (res *valueobject.EjemploDeTextoResponse, err error)
	Eliminar(ctx *ns.Context, id int64) (err error)
	Habilitar(ctx *ns.Context, id int64) (err error)
	Deshabilitar(ctx *ns.Context, id int64) (err error)
}
