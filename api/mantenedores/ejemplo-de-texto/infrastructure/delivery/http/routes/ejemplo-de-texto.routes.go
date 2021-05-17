package routes

import (
	"gitlab.com/ns-desarrollo-web/erp/ns-core-service"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/domain/usecase"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/infrastructure/delivery/http/controller"
)

func NewEjemploDeTextoRoutes(router *ns.Router, useCase usecase.EjemploDeTextoUseCase) {
	c := controller.NewEjemploDeTextoController(useCase)

	router.GET("/ejemplo-de-texto", c.Listar)
	router.GET("/ejemplo-de-texto-buscar", c.Buscar)
	router.GET("/ejemplo-de-texto/:id", c.BuscarPorId)
	router.POST("/ejemplo-de-texto", c.Crear)
	router.PUT("/ejemplo-de-texto/:id", c.Actualizar)
	router.DELETE("/ejemplo-de-texto/:id", c.Eliminar)
	router.PATCH("/ejemplo-de-texto/:id/habilitar", c.Habilitar)
	router.PATCH("/ejemplo-de-texto/:id/deshabilitar", c.Deshabilitar)
}
