package mapping

import (
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/dto"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/entity"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/model"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/valueobject"
)

func MapEjemploDeTextoPaginación(Modelo []model.EjemploDeTextoModel, total int64, filtered int64) *valueobject.EjemploDeTextoPaginacion {
	var res []valueobject.EjemploDeTextoResponse

	for i := range Modelo {
		res = append(res, valueobject.EjemploDeTextoResponse{
			Id:               	Modelo[i].Id,
	Nombre:               	Modelo[i].Nombre,
	Codigo:               	Modelo[i].Codigo,
	Habilitado:               	Modelo[i].Habilitado,
	
		})
	}

	return &valueobject.EjemploDeTextoPaginacion{
		Data: res,
		Meta: valueobject.MetaPaginacion{
			RecordsFiltered: filtered,
			RecordsTotal:    total,
		},
	}
}

func MapEjemploDeTextoSuggest(mod []model.EjemploDeTextoModel) (res []*valueobject.EjemploDeTextoSuggest) {
	for i := range mod {

		var allData interface{}

		allData = mod

		if mod[i].Cols != nil {
			allData = mod[i].Cols
		}

		res = append(res, &valueobject.EjemploDeTextoSuggest{
			Id:          mod[i].Id,
			Code:        &mod[i].Codigo,
			Label:       &mod[i].Nombre,
			Description: &mod[i].Codigo,
			Badge:       &mod[i].Codigo,
			Data:        allData,
		})
	}

	return
}

func MapEjemploDeTextoModelToVO(Modelo *model.EjemploDeTextoModel) *valueobject.EjemploDeTextoResponse {


	return &valueobject.EjemploDeTextoResponse{
		Id:               	Modelo.Id,
	Nombre:               	Modelo.Nombre,
	Codigo:               	Modelo.Codigo,
	Habilitado:               	Modelo.Habilitado,
	
	}
}

func MapEjemploDeTextoDTOTOModel(Request *dto.EjemploDeTextoRequest) *model.EjemploDeTextoModel {
	return &model.EjemploDeTextoModel{
		EjemploDeTextoEntity: entity.EjemploDeTextoEntity{
			Id:               	Request.Id,
	Nombre:               	Request.Nombre,
	Codigo:               	Request.Codigo,
	Habilitado:               	Request.Habilitado,
	
		},
	}
}
