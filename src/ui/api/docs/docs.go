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
            "name": "DIT - IFAL",
            "email": "wmrn1@aluno.ifal.edu.br"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/accounts/profile": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Esta rota retorna todas as informações de todas as contas cadastradas no banco de dados.\nDados como \"professional\" irão somente aparecer caso a role da conta for própria para contenção desses.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Geral"
                ],
                "summary": "Pesquisar dados do perfil de uma conta.",
                "operationId": "Accounts.FindProfile",
                "responses": {
                    "200": {
                        "description": "Requisição realizada com sucesso.",
                        "schema": {
                            "$ref": "#/definitions/response.Account"
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
            },
            "put": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Geral"
                ],
                "summary": "Atualizar dados do perfil de uma conta.",
                "operationId": "Account.UpdateProfile",
                "parameters": [
                    {
                        "description": "JSON com todos os dados necessários para o processo de atualização de dados do perfil.",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateAccountProfile"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
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
        },
        "/accounts/update-password": {
            "put": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Geral"
                ],
                "summary": "Realizar a atualização de senha de uma conta.",
                "operationId": "Account.UpdateAccountPassword",
                "parameters": [
                    {
                        "description": "JSON com todos os dados necessários para a atualização da senha da conta.",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdatePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
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
        },
        "/admin/accounts": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Esta rota retorna todas as informações de todas as contas cadastradas no banco de dados.\nDados como \"professional\" irão somente aparecer caso a role da conta for própria para contenção desses.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Administrador"
                ],
                "summary": "Listar todas as contas existentes do banco de dados.",
                "operationId": "Accounts.List",
                "responses": {
                    "200": {
                        "description": "Requisição realizada com sucesso.",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.Account"
                            }
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
            },
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Ao enviar dados para cadastro de uma nova conta, os dados relacionados à \"Profissional\"\nsão facultativos, tendo somente que enviar os dados que são relacionados à role definida.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Administrador"
                ],
                "summary": "Cadastrar uma nova conta de usuário",
                "operationId": "Accounts.Create",
                "parameters": [
                    {
                        "description": "JSON com todos os dados necessários para o cadastro de uma conta de usuário.",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateAccount"
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
        },
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Anônimo"
                ],
                "summary": "Adquirir autorização de acesso aos recursos da API através de credenciais de uma conta.",
                "operationId": "Auth.Login",
                "parameters": [
                    {
                        "description": "JSON com todos os dados necessários para o processo de autenticação.",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Credentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Requisição realizada com sucesso.",
                        "schema": {
                            "$ref": "#/definitions/response.Authorization"
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
        },
        "/auth/logout": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Geral"
                ],
                "summary": "Remove a sessão do registro de sessões permitidas.",
                "operationId": "Auth.Logout",
                "responses": {
                    "204": {
                        "description": "Requisição realizada com sucesso."
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
        },
        "/auth/reset-password": {
            "post": {
                "description": "cadastra uma nova entrada para a entidade ` + "`" + `password_reset` + "`" + ` vinculada à conta da sessão\ne envia um e-mail para o email dessa.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Anônimo"
                ],
                "summary": "Solicitar email com token para atualização de senha.",
                "operationId": "Auth.PasswordReset",
                "parameters": [
                    {
                        "description": "JSON com todos os dados necessários para resetar a senha por email.",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreatePasswordReset"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
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
        },
        "/auth/reset-password/{token}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Anônimo"
                ],
                "summary": "Verificar a existência de uma solicitação de atualização de senha por token.",
                "operationId": "Auth.FindPasswordResetByToken",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token recebido pelo email da conta do usuário da plataforma.",
                        "name": "token",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
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
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Anônimo"
                ],
                "summary": "Atualizar a senha de uma conta a partir de um token de atualização de senha.",
                "operationId": "Auth.UpdatePasswordByPasswordReset",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token recebido pelo email da conta do usuário da plataforma.",
                        "name": "token",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "JSON com todos os dados necessários para resetar a senha por email.",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdatePasswordByPasswordReset"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
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
        },
        "/res/account-roles": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Pode ser utilizada para visualizar as funções de conta cadastradas no banco de dados e adquirir o\nidentificador da função desejada para a criação de uma nova conta.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recursos"
                ],
                "summary": "Listar todas as funções de conta existentes do banco de dados.",
                "operationId": "Resources.ListAccountRoles",
                "responses": {
                    "200": {
                        "description": "Requisição realizada com sucesso.",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.Role"
                            }
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
        "request.CreateAccount": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "type": "string"
                },
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role_code": {
                    "type": "string"
                }
            }
        },
        "request.CreatePasswordReset": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "request.Credentials": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.UpdateAccountProfile": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "request.UpdatePassword": {
            "type": "object",
            "properties": {
                "current_password": {
                    "type": "string"
                },
                "new_password": {
                    "type": "string"
                }
            }
        },
        "request.UpdatePasswordByPasswordReset": {
            "type": "object",
            "properties": {
                "new_password": {
                    "type": "string"
                }
            }
        },
        "response.Account": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "professional": {
                    "$ref": "#/definitions/response.Professional"
                },
                "profile": {
                    "$ref": "#/definitions/response.Person"
                },
                "role": {
                    "$ref": "#/definitions/response.Role"
                }
            }
        },
        "response.Authorization": {
            "type": "object",
            "properties": {
                "access_token": {
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
        },
        "response.Person": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "type": "string"
                },
                "cpf": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "response.Professional": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "response.Role": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "DIT Backend API",
	Description:      "DIT Backend template for new backend projects",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
