package routes

import (
	"gitlab.com/ns-desarrollo-web/erp/ns-core-service"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/usecase"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/infrastructure/delivery/http/controller"
)

func NewMotivosExtornoPedidoRoutes(router *ns.Router, useCase usecase.MotivosExtornoPedidoUseCase) {
	c := controller.NewMotivosExtornoPedidoController(useCase)

	router.GET("/motivos-extorno-pedido", c.Listar)
	router.GET("/motivos-extorno-pedido-buscar", c.Buscar)
	router.GET("/motivos-extorno-pedido/:id", c.BuscarPorId)
	router.POST("/motivos-extorno-pedido", c.Crear)
	router.PUT("/motivos-extorno-pedido/:id", c.Actualizar)
	router.DELETE("/motivos-extorno-pedido/:id", c.Eliminar)
	router.PATCH("/motivos-extorno-pedido/:id/habilitar", c.Habilitar)
	router.PATCH("/motivos-extorno-pedido/:id/deshabilitar", c.Deshabilitar)
}
