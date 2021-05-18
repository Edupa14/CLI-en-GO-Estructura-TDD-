package valueobject

type MotivosExtornoPedidoResponse struct {
	Id	int64	`json:"id"`
	Nombre	string	`json:"nombre"`
	Codigo	string	`json:"codigo"`
	Habilitado	bool	`json:"habilitado"`
	
}


type MotivosExtornoPedidoPaginacion struct {
	Data []MotivosExtornoPedidoResponse `json:"data"`
	Meta MetaPaginacion     `json:"meta"`
}

type MetaPaginacion struct {
	RecordsFiltered int64 `json:"recordsFiltered"`
	RecordsTotal    int64 `json:"recordsTotal"`
}

type MotivosExtornoPedidoSuggest struct {
	Id          int64       `json:"id"`
	Code        *string     `json:"code"`
	Label       *string     `json:"label"`
	Description *string     `json:"description"`
	Badge       *string     `json:"badge"`
	Data        interface{} `json:"data"`
}
