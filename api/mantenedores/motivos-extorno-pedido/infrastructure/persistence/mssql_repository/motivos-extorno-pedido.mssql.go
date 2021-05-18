package mssql_repository

import (
	"gitlab.com/ns-desarrollo-web/erp/ns-core-service"
	db "gitlab.com/ns-desarrollo-web/erp/ns-core-service/services/databases"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/model"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/repository"
	"ns-compras-service/api/utils"
)

type MotivosExtornoPedidoMssqlRepository struct {
	db db.Sql
}

func (p *MotivosExtornoPedidoMssqlRepository) Listar(client *ns.Context, start int, length int, search string, order string) (entidad []model.MotivosExtornoPedidoModel, total int64, filtered int64, err error) {
	rs, err := p.db.Run(client, `
		EXEC [NS_MAINTAINER_LIST]
				@idempresa = '',
				@columns = 'id,nombre,codigo,habilitado,fecha_guardado,transaction_uid',
				@table = 'TMMOTIVOS_EXTORNO_PEDIDO',
				@start = $1,
				@length = $2,
				@search = $3,
				@order = '',
				@advanced_filter = '',
				@default_filter = 'nombre,codigo'
	`, start, length, search)
	if err != nil {
		return
	}

	err = rs.Unmarshal(&entidad)
	if err != nil {
		return
	}

	total, filtered = rs.Total, rs.Filtered

	return
}

func (p *MotivosExtornoPedidoMssqlRepository) Crear(ctx *ns.Context, modelo *model.MotivosExtornoPedidoModel) (id int64, err error) {
	json, err := db.Marshal(modelo)

	id, err = p.db.QueryWithLastID(ctx, "EXEC NS_TMMOTIVOS_EXTORNO_PEDIDO_I  @JSON = $1;", string(json))

	return
}

func (p *MotivosExtornoPedidoMssqlRepository) Actualizar(ctx *ns.Context, id int64, modelo *model.MotivosExtornoPedidoModel) (err error) {
	json, err := db.Marshal(modelo)

	_, err = p.db.Query(ctx, "EXEC NS_TMMOTIVOS_EXTORNO_PEDIDO_U @ID = $1, @JSON = $2;", id, string(json))

	return
}

func (p *MotivosExtornoPedidoMssqlRepository) Eliminar(ctx *ns.Context, id int64) (err error) {
	_, err = p.db.Exec(ctx, `EXEC NS_TMMOTIVOS_EXTORNO_PEDIDO_D @id = $1`, id)

	return
}

func (p *MotivosExtornoPedidoMssqlRepository) BuscarPorId(ctx *ns.Context, id int64) (modelo *model.MotivosExtornoPedidoModel, err error) {
	var m []model.MotivosExtornoPedidoModel

	err = p.db.RunUnmarshal(ctx, &m, `EXEC NS_TMMOTIVOS_EXTORNO_PEDIDO_F @id = $1`, id)
	if err != nil{
		return
	}
	modelo = &m[0]

	return
}

func (p *MotivosExtornoPedidoMssqlRepository) Buscar(ctx *ns.Context, search, advancedSearch string, isColumns bool) (modelo []model.MotivosExtornoPedidoModel, err error) {

	if isColumns {
		var m model.MotivosExtornoPedidoModel
		cols := utils.ObtenerCampos(m.MotivosExtornoPedidoEntity)
		modelo = append(modelo, model.MotivosExtornoPedidoModel{
			Cols: cols,
		})
		return
	}

	err = p.db.Select(ctx, &modelo, `
		EXEC NS_GETDATA_F
			@idempresa = '',
			@columns = 'id,nombre,codigo,habilitado,fecha_guardado,transaction_uid',
			@table = 'TMMOTIVOS_EXTORNO_PEDIDO',
			@search = $1,
			@limit = $2,
 			@id_filter = '',
			@default_filter = 'nombre,codigo',
			@advanced_filter = $3`, search, 10, advancedSearch)

	if err != nil {
		return
	}
	return
}

func (p *MotivosExtornoPedidoMssqlRepository) HabilitarDeshabilitar(ctx *ns.Context, id int64, status bool) (err error) {
	_, err = p.db.Exec(ctx, `EXEC NS_TMMOTIVOS_EXTORNO_PEDIDO_ED @id = $1, @estado = $2`, id, status)

	return
}

func NewMotivosExtornoPedidoRepository(db db.Sql) repository.MotivosExtornoPedidoRepository {
	return &MotivosExtornoPedidoMssqlRepository{db: db}
}
