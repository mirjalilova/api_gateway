// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/comments": {
            "get": {
                "description": "Get all comments for a memory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Get comments",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Memory ID",
                        "name": "memory_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/memory.Comment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update the content of an existing comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Update a comment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Comment ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Comment Update Body",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/memory.CommentUpdateBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Comment updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new comment on a memory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Create a comment",
                "parameters": [
                    {
                        "description": "Comment Create Body",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/memory.CommentCreateBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Comment created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a comment by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Delete a comment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Comment ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Comment deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/media": {
            "get": {
                "description": "Get a media entry by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "media"
                ],
                "summary": "Get media",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Memory ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/memory.Media"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new media entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "media"
                ],
                "summary": "Create media",
                "parameters": [
                    {
                        "description": "Media Create Body",
                        "name": "media",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/memory.MediaCreateBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Media created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a media entry by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "media"
                ],
                "summary": "Delete media",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Media ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Media deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/memories": {
            "get": {
                "description": "Get a memory by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "memories"
                ],
                "summary": "Get a memory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Memory ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/memory.MemoryRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a memory by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "memories"
                ],
                "summary": "Update a memory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Memory ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "MemoryCreateBody",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/memory.MemoryCreateBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Memory updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new memory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "memories"
                ],
                "summary": "Create a memory",
                "parameters": [
                    {
                        "description": "MemoryCreateBody",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/memory.MemoryCreateBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Memory created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a memory by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "memories"
                ],
                "summary": "Delete a memory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Memory ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Memory deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/memories/all": {
            "get": {
                "description": "Get all memories with optional filters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "memories"
                ],
                "summary": "Get all memories",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Start Date",
                        "name": "start_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "End Date",
                        "name": "end_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Tag",
                        "name": "tag",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Type",
                        "name": "type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/memory.GetAllRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/shares": {
            "get": {
                "description": "Get all shares for a memory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shares"
                ],
                "summary": "Get shares",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Memory ID",
                        "name": "memory_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/memory.ShareRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Remove a share for a memory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shares"
                ],
                "summary": "Unshare a memory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Share ID",
                        "name": "share_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Share Delete Body",
                        "name": "share",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/memory.ShareDeleteBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Share updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Share a memory with other users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shares"
                ],
                "summary": "Create a share",
                "parameters": [
                    {
                        "description": "Share Create Body",
                        "name": "share",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/memory.ShareCreateBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Share created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "memory.Comment": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "memory_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "writer_id": {
                    "type": "string"
                }
            }
        },
        "memory.CommentCreateBody": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "memory_id": {
                    "type": "string"
                }
            }
        },
        "memory.CommentUpdateBody": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                }
            }
        },
        "memory.GetAllRes": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "memories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/memory.MemoryRes"
                    }
                }
            }
        },
        "memory.Media": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "memory_id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "memory.MediaCreateBody": {
            "type": "object",
            "properties": {
                "memory_id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "memory.MemoryCreateBody": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "locations": {
                    "$ref": "#/definitions/memory.Point"
                },
                "place_name": {
                    "type": "string"
                },
                "privacy": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "memory.MemoryRes": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "locations": {
                    "$ref": "#/definitions/memory.Point"
                },
                "place_name": {
                    "type": "string"
                },
                "privacy": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "memory.Point": {
            "type": "object",
            "properties": {
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                }
            }
        },
        "memory.ShareCreateBody": {
            "type": "object",
            "properties": {
                "memory_id": {
                    "type": "string"
                },
                "shared_with": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "memory.ShareDeleteBody": {
            "type": "object",
            "properties": {
                "memory_id": {
                    "type": "string"
                },
                "unshared_with": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "memory.ShareRes": {
            "type": "object",
            "properties": {
                "memony_title": {
                    "type": "string"
                },
                "shared_with": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
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
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Time Capsule API",
	Description:      "API for Time Capsule resources",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
