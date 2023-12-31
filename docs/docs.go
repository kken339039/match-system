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
        "/api/users": {
            "post": {
                "description": "Add a new user to the matching system and find any possible matches for the new user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Add a new user and find matches",
                "parameters": [
                    {
                        "description": "New user details",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/match-system_internal_user_dtos.AddSinglePersonAndMatchRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/match-system_internal_user_dtos.AddSinglePersonAndMatchResponse"
                        }
                    }
                }
            }
        },
        "/api/users/query_single": {
            "get": {
                "description": "Query single users based on the specified count.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Query single users from the match system.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Number of users to query",
                        "name": "N",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/match-system_internal_user_dtos.QuerySinglePeopleResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/users/{userId}": {
            "delete": {
                "description": "Remove a user from the match system so that the user cannot be matched anymore.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Remove a user from the match system.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID to be removed",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "match-system_internal_user_dtos.AddSinglePersonAndMatchRequest": {
            "type": "object",
            "properties": {
                "gender": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "wanted_dates": {
                    "type": "integer"
                }
            }
        },
        "match-system_internal_user_dtos.AddSinglePersonAndMatchResponse": {
            "type": "object",
            "properties": {
                "gender": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "matches": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/match-system_internal_user_dtos.MatchedUserResponse"
                    }
                },
                "name": {
                    "type": "string"
                },
                "wantedDates": {
                    "type": "integer"
                }
            }
        },
        "match-system_internal_user_dtos.MatchedUserResponse": {
            "type": "object",
            "properties": {
                "gender": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "match-system_internal_user_dtos.PeopleResponse": {
            "type": "object",
            "properties": {
                "gender": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "wantedDates": {
                    "type": "integer"
                }
            }
        },
        "match-system_internal_user_dtos.QuerySinglePeopleResponse": {
            "type": "object",
            "properties": {
                "people": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/match-system_internal_user_dtos.PeopleResponse"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
