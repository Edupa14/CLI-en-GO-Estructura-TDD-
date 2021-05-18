package mapping

import (
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/dto"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/entity"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/model"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/valueobject"
)

func MapMotivosExtornoPedidoPaginación(Modelo []model.MotivosExtornoPedidoModel, total int64, filtered int64) *valueobject.MotivosExtornoPedidoPaginacion {
	var res []valueobject.MotivosExtornoPedidoResponse

	for i := range Modelo {
		res = append(res, valueobject.MotivosExtornoPedidoResponse{
			Id:               	Modelo[i].Id,
	Nombre:               	Modelo[i].Nombre,
	Codigo:               	Modelo[i].Codigo,
	Habilitado:               	Modelo[i].Habilitado,
	
		})
	}

	return &valueobject.MotivosExtornoPedidoPaginacion{
		Data: res,
		Meta: valueobject.MetaPaginacion{
			RecordsFiltered: filtered,
			RecordsTotal:    total,
		},
	}
}

func MapMotivosExtornoPedidoSuggest(mod []model.MotivosExtornoPedidoModel) (res []*valueobject.MotivosExtornoPedidoSuggest) {
	for i := range mod {

		var allData interface{}

		allData = mod

		if mod[i].Cols != nil {
			allData = mod[i].Cols
		}

		res = append(res, &valueobject.MotivosExtornoPedidoSuggest{
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

func MapMotivosExtornoPedidoModelToVO(Modelo *model.MotivosExtornoPedidoModel) *valueobject.MotivosExtornoPedidoResponse {


	return &valueobject.MotivosExtornoPedidoResponse{
		Id:               	Modelo.Id,
	Nombre:               	Modelo.Nombre,
	Codigo:               	Modelo.Codigo,
	Habilitado:               	Modelo.Habilitado,
	
	}
}

func MapMotivosExtornoPedidoDTOTOModel(Request *dto.MotivosExtornoPedidoRequest) *model.MotivosExtornoPedidoModel {
	return &model.MotivosExtornoPedidoModel{
		MotivosExtornoPedidoEntity: entity.MotivosExtornoPedidoEntity{
			Id:               	Request.Id,
	Nombre:               	Request.Nombre,
	Codigo:               	Request.Codigo,
	Habilitado:               	Request.Habilitado,
	
		},
	}
}
