// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Thiago gazaroli",
            "email": "tgazaroli@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/search-link": {
            "post": {
                "description": "Ao Prencher os o body da requisição com a url base, email que o usario quer receber os links encontrados e a quantidade de links que deverá ser encontrado\nOBS: Recomendamos  preencher o campo de quantidade de links com até 150, estaremos trabalhando para aumentar exponencialmente esse número",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Buscador de urls a partir da url enviada",
                "operationId": "Crawler.SearchLinks",
                "parameters": [
                    {
                        "description": "JSON com todos os dados necessários para que seja possivel realizar a buscas das urls",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.SearchLink"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Requisição realizada com sucesso.",
                        "schema": {
                            "$ref": "#/definitions/response.ID"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado. Por favor, contate o suporte.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados está temporariamente indisponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.SearchLink": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "number_links": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "response.ErrorMessage": {
            "type": "object",
            "properties": {
                "invalid_fields": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.InvalidField"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "response.ID": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "response.InvalidField": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "field_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Web Crawler Backend API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
