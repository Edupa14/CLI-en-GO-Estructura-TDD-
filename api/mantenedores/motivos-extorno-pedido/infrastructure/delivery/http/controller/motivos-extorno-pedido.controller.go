package controller

import (
	"gitlab.com/ns-desarrollo-web/erp/ns-core-service"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/dto"
	"ns-compras-service/api/mantenedores/motivos-extorno-pedido/domain/usecase"
	"ns-compras-service/api/utils"
	"strconv"
)

type MotivosExtornoPedidoController struct {
	useCase usecase.MotivosExtornoPedidoUseCase
}

func (p *MotivosExtornoPedidoController) Listar(client *ns.Client) {
	pag := client.GetPagination()

	rs, err := p.useCase.Listar(client.GetContext(), pag.Start, pag.Length, pag.Search, pag.Order)
	if err != nil {
		client.Fail(err)
		return
	}

	client.Ok(rs)
}

func (p *MotivosExtornoPedidoController) Buscar(client *ns.Client) {
	search := client.Query("search")
	var advancedSearch string

	advanced, err := utils.ParseBoolean(client.Query("advanced"))
	if err != nil {
		client.Fail(err)
		return
	}

	columns, err := utils.ParseBoolean(client.Query("c"))
	if err != nil {
		client.Fail(err)
		return
	}

	if advanced && !columns { // is Advanced
		advancedSearch, err = utils.StringKey(search)
		if err != nil {
			client.Fail(err)
			return
		}
	}

	rs, err := p.useCase.Buscar(client.GetContext(), search, advancedSearch, columns)
	if err != nil {
		client.Fail(err)
		return
	}

	client.Ok(rs)
}

func (p *MotivosExtornoPedidoController) BuscarPorId(client *ns.Client) {
	id, err := strconv.ParseInt(client.Param("id"), 10, 64)
	if err != nil {
		client.Fail(err)
		return
	}

	rs, err := p.useCase.BuscarPorId(client.GetContext(), id)
	if err != nil {
		client.Fail(err)
		return
	}

	client.Ok(rs)
}

func (p *MotivosExtornoPedidoController) Crear(client *ns.Client) {
	var req dto.MotivosExtornoPedidoRequest
	err := client.UnmarshalDTO(&req)
	if err != nil {
		client.Fail(err)
		return
	}

	rs, err := p.useCase.Crear(client.GetContext(), &req)
	if err != nil {
		client.Fail(err)
		return
	}

	client.Ok(rs)
}

func (p *MotivosExtornoPedidoController) Actualizar(client *ns.Client) {
	var req dto.MotivosExtornoPedidoRequest
	err := client.UnmarshalDTO(&req)
	if err != nil {
		client.Fail(err)
		return
	}

	id, err := strconv.ParseInt(client.Param("id"), 10, 64)
	if err != nil {
		client.Fail(err)
		return
	}

	rs, err := p.useCase.Actualizar(client.GetContext(), id, &req)
	if err != nil {
		client.Fail(err)
		return
	}

	client.Ok(rs)
}

func (p *MotivosExtornoPedidoController) Eliminar(client *ns.Client) {
	id, err := strconv.ParseInt(client.Param("id"), 10, 64)
	if err != nil {
		client.Fail(err)
		return
	}

	err = p.useCase.Eliminar(client.GetContext(), id)
	if err != nil {
		client.Fail(err)
		return
	}

	client.Ok()
}

func (p *MotivosExtornoPedidoController) Habilitar(client *ns.Client) {
	id, err := strconv.ParseInt(client.Param("id"), 10, 64)
	if err != nil {
		client.Fail(err)
		return
	}

	err = p.useCase.Habilitar(client.GetContext(), id)
	if err != nil {
		client.Fail(err)
		return
	}

	client.Ok()
}

func (p *MotivosExtornoPedidoController) Deshabilitar(client *ns.Client) {
	id, err := strconv.ParseInt(client.Param("id"), 10, 64)
	if err != nil {
		client.Fail(err)
		return
	}

	err = p.useCase.Deshabilitar(client.GetContext(), id)
	if err != nil {
		client.Fail(err)
		return
	}

	client.Ok()
}

func NewMotivosExtornoPedidoController(useCase usecase.MotivosExtornoPedidoUseCase) *MotivosExtornoPedidoController {
	return &MotivosExtornoPedidoController{useCase: useCase}
}
