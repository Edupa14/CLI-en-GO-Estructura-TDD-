package repository

import (
	"gitlab.com/ns-desarrollo-web/erp/ns-core-service"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/model"
)

type EjemploDeTextoRepository interface {
	Listar(ctx *ns.Context, start int, length int, search string, order string) (res []model.EjemploDeTextoModel, total int64, filtered int64, err error)
	Buscar(ctx *ns.Context, search, advancedSearch string, columns bool) (res []model.EjemploDeTextoModel, err error)
	BuscarPorId(ctx *ns.Context, id int64) (res *model.EjemploDeTextoModel, err error)
	Crear(ctx *ns.Context, productoModel *model.EjemploDeTextoModel) (id int64, err error)
	Actualizar(ctx *ns.Context, id int64, productoModel *model.EjemploDeTextoModel) (err error)
	Eliminar(ctx *ns.Context, id int64) (err error)
	HabilitarDeshabilitar(ctx *ns.Context, id int64, status bool) (err error)
}
