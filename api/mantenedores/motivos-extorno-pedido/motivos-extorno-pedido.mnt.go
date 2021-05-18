package motivos_extorno_pedido

import (
	"gitlab.com/ns-desarrollo-web/erp/ns-core-service"
	db "gitlab.com/ns-desarrollo-web/erp/ns-core-service/services/databases"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/application/service"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/application/usecase"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/infrastructure/delivery/http/routes"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/infrastructure/persistence/mssql_repository"
)

func Create(router *ns.Router, db db.Sql) {
	// Repository implementation
	MotivosExtornoPedidoRepository := mssql_repository.NewMotivosExtornoPedidoRepository(db)

	// Service Implementation
	MotivosExtornoPedidoService := service.NewMotivosExtornoPedidoService(MotivosExtornoPedidoRepository)

	// UseCase implementation
	MotivosExtornoPedidoUseCase := usecase.NewMotivosExtornoPedidoUseCase(MotivosExtornoPedidoRepository, MotivosExtornoPedidoService)

	// Http router controller register
	routes.NewMotivosExtornoPedidoRoutes(router, MotivosExtornoPedidoUseCase)
}
