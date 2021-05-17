package ejemplo_de_texto

import (
	"gitlab.com/ns-desarrollo-web/erp/ns-core-service"
	db "gitlab.com/ns-desarrollo-web/erp/ns-core-service/services/databases"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/application/service"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/application/usecase"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/infrastructure/delivery/http/routes"
	"ns-compras-service/api/mantenedores/ejemplo-de-texto/infrastructure/persistence/mssql_repository"
)

func Create(router *ns.Router, db db.Sql) {
	// Repository implementation
	EjemploDeTextoRepository := mssql_repository.NewEjemploDeTextoRepository(db)

	// Service Implementation
	EjemploDeTextoService := service.NewEjemploDeTextoService(EjemploDeTextoRepository)

	// UseCase implementation
	EjemploDeTextoUseCase := usecase.NewEjemploDeTextoUseCase(EjemploDeTextoRepository, EjemploDeTextoService)

	// Http router controller register
	routes.NewEjemploDeTextoRoutes(router, EjemploDeTextoUseCase)
}
