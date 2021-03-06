// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/posts": {
            "get": {
                "description": "get all posts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get posts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.PostListSerializer"
                        }
                    }
                }
            }
        },
        "/api/v1/posts/": {
            "post": {
                "description": "create new post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "create new post",
                "parameters": [
                    {
                        "description": "Post Body",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/form.PostCreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.SinglePostSerializer"
                        }
                    }
                }
            }
        },
        "/api/v1/posts/{post_id}": {
            "get": {
                "description": "find post by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "find post by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post ID",
                        "name": "post_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.SinglePostSerializer"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post ID",
                        "name": "post_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.SuccessSerializer"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post ID",
                        "name": "post_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Post Body",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/form.PostUpdateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.SinglePostSerializer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "form.PostCreateReq": {
            "type": "object",
            "properties": {
                "post": {
                    "type": "object",
                    "properties": {
                        "body": {
                            "type": "string"
                        },
                        "title": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "form.PostUpdateReq": {
            "type": "object",
            "properties": {
                "post": {
                    "type": "object",
                    "properties": {
                        "body": {
                            "type": "string"
                        },
                        "title": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "serializer.PostListSerializer": {
            "type": "object",
            "properties": {
                "posts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/serializer.postSerializer"
                    }
                },
                "postsCount": {
                    "type": "integer"
                }
            }
        },
        "serializer.SinglePostSerializer": {
            "type": "object",
            "properties": {
                "post": {
                    "type": "object",
                    "$ref": "#/definitions/serializer.postSerializer"
                }
            }
        },
        "serializer.SuccessSerializer": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "object",
                    "$ref": "#/definitions/serializer.successEntry"
                }
            }
        },
        "serializer.postSerializer": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "serializer.successEntry": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "petstore.swagger.io",
	BasePath:    "/v2",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server Petstore server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
