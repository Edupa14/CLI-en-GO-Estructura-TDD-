package mssql_repository

import (
	"gitlab.com/ns-desarrollo-web/erp/ns-core-service"
	db "gitlab.com/ns-desarrollo-web/erp/ns-core-service/services/databases"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/model"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/repository"
	"ns-compras-service/api/utils"
)

type EjemploDeTextoMssqlRepository struct {
	db db.Sql
}

func (p *EjemploDeTextoMssqlRepository) Listar(client *ns.Context, start int, length int, search string, order string) (entidad []model.EjemploDeTextoModel, total int64, filtered int64, err error) {
	rs, err := p.db.Run(client, `
		EXEC [NS_MAINTAINER_LIST]
				@idempresa = '',
				@columns = 'id,nombre,codigo,habilitado,fecha_guardado,transaction_uid',
				@table = 'TMEJEMPLO_DE_TEXTO',
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

func (p *EjemploDeTextoMssqlRepository) Crear(ctx *ns.Context, modelo *model.EjemploDeTextoModel) (id int64, err error) {
	json, err := db.Marshal(modelo)

	id, err = p.db.QueryWithLastID(ctx, "EXEC NS_TMEJEMPLO_DE_TEXTO_I  @JSON = $1;", string(json))

	return
}

func (p *EjemploDeTextoMssqlRepository) Actualizar(ctx *ns.Context, id int64, modelo *model.EjemploDeTextoModel) (err error) {
	json, err := db.Marshal(modelo)

	_, err = p.db.Query(ctx, "EXEC NS_TMEJEMPLO_DE_TEXTO_U @ID = $1, @JSON = $2;", id, string(json))

	return
}

func (p *EjemploDeTextoMssqlRepository) Eliminar(ctx *ns.Context, id int64) (err error) {
	_, err = p.db.Exec(ctx, `EXEC NS_TMEJEMPLO_DE_TEXTO_D @id = $1`, id)

	return
}

func (p *EjemploDeTextoMssqlRepository) BuscarPorId(ctx *ns.Context, id int64) (modelo *model.EjemploDeTextoModel, err error) {
	var m []model.EjemploDeTextoModel

	err = p.db.RunUnmarshal(ctx, &m, `EXEC NS_EjemploDeTexto_F @id = $1`, id)
	if err != nil{
		return
	}
	modelo = &m[0]

	return
}

func (p *EjemploDeTextoMssqlRepository) Buscar(ctx *ns.Context, search, advancedSearch string, isColumns bool) (modelo []model.EjemploDeTextoModel, err error) {

	if isColumns {
		var m model.EjemploDeTextoModel
		cols := utils.ObtenerCampos(m.EjemploDeTextoEntity)
		modelo = append(modelo, model.EjemploDeTextoModel{
			Cols: cols,
		})
		return
	}

	err = p.db.Select(ctx, &modelo, `
		EXEC NS_GETDATA_F
			@idempresa = '',
			@columns = 'id,nombre,codigo,habilitado,fecha_guardado,transaction_uid',
			@table = 'TMEJEMPLO_DE_TEXTO',
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

func (p *EjemploDeTextoMssqlRepository) HabilitarDeshabilitar(ctx *ns.Context, id int64, status bool) (err error) {
	_, err = p.db.Exec(ctx, `EXEC NSP_TEjemploDeTexto_A @id = $1, @estado = $2`, id, status)

	return
}

func NewEjemploDeTextoRepository(db db.Sql) repository.EjemploDeTextoRepository {
	return &EjemploDeTextoMssqlRepository{db: db}
}
