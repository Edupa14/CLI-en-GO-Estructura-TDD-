// app.go configuración extra
"ns-compras-service/api/mantenedores/motivos-extorno-pedido/"
motivos_extorno_pedido.Create(router, db)

// template kraken config
{
    "endpoint": "/erp/v1/motivos-extorno-pedido",
    "method": "GET",
    "output_encoding": "no-op",
    "querystring_params": ["start", "length", "search", "order"],
    "headers_to_pass": [
        "Authorization"
    ],
    "backend": [{
        "encoding": "no-op",
        "url_pattern": "/motivos-extorno-pedido",
        "host": [
            "{{ .compras_service }}"
        ]
    }]
},
{
    "endpoint": "/erp/v1/motivos-extorno-pedido-buscar",
    "method": "GET",
    "output_encoding": "no-op",
    "querystring_params": ["search","advanced","c"],
    "headers_to_pass": [
        "Authorization"
    ],
    "backend": [{
        "encoding": "no-op",
        "url_pattern": "/motivos-extorno-pedido-buscar",
        "host": [
            "{{ .compras_service }}"
        ]
    }]
},
{
    "endpoint": "/erp/v1/motivos-extorno-pedido/{id}",
    "method": "GET",
    "output_encoding": "no-op",
    "headers_to_pass": [
        "Authorization"
    ],
    "backend": [{
        "encoding": "no-op",
        "url_pattern": "/motivos-extorno-pedido/{id}",
        "host": [
            "{{ .compras_service }}"
        ]
    }]
},
{
    "endpoint": "/erp/v1/motivos-extorno-pedido",
    "method": "POST",
    "output_encoding": "no-op",
    "headers_to_pass": [
        "Authorization"
    ],
    "backend": [{
        "encoding": "no-op",
        "url_pattern": "/motivos-extorno-pedido",
        "host": [
            "{{ .compras_service }}"
        ]
    }]
},
{
    "endpoint": "/erp/v1/motivos-extorno-pedido/{id}",
    "method": "PUT",
    "output_encoding": "no-op",
    "headers_to_pass": [
        "Authorization"
    ],
    "backend": [{
        "encoding": "no-op",
        "url_pattern": "/motivos-extorno-pedido/{id}",
        "host": [
            "{{ .compras_service }}"
        ]
    }]
},
{
    "endpoint": "/erp/v1/motivos-extorno-pedido/{id}",
    "method": "DELETE",
    "output_encoding": "no-op",
    "headers_to_pass": [
        "Authorization"
    ],
    "backend": [{
        "encoding": "no-op",
        "url_pattern": "/motivos-extorno-pedido/{id}",
        "host": [
            "{{ .compras_service }}"
        ]
    }]
},
{
    "endpoint": "/erp/v1/motivos-extorno-pedido/{id}/habilitar",
    "method": "PATCH",
    "output_encoding": "no-op",
    "headers_to_pass": [
        "Authorization"
    ],
    "backend": [{
        "encoding": "no-op",
        "url_pattern": "/motivos-extorno-pedido/{id}/habilitar",
        "host": [
            "{{ .compras_service }}"
        ]
    }]
},
{
    "endpoint": "/erp/v1/motivos-extorno-pedido/{id}/deshabilitar",
    "method": "PATCH",
    "output_encoding": "no-op",
    "headers_to_pass": [
        "Authorization"
    ],
    "backend": [{
        "encoding": "no-op",
        "url_pattern": "/motivos-extorno-pedido/{id}/deshabilitar",
        "host": [
            "{{ .compras_service }}"
        ]
    }]
}
