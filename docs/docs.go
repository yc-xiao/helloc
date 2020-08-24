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
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/": {
            "post": {
                "description": "通过JWT验证账号密码，获取token",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "认证"
                ],
                "summary": "登录认证",
                "parameters": [
                    {
                        "default": "{\"name\": \"xiaoming\", \"password\": \"123456\"}",
                        "description": "name(*用户名) password(*密码)",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"token生成成功!\", \"results\": \"tokenString\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"参数错误/token生成失败!/账号不存在或密码错误\", \"results\": null}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth2/": {
            "post": {
                "description": "用于Swagger，只做测试",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "认证"
                ],
                "summary": "登录认证2，用于Swagger，只做测试",
                "parameters": [
                    {
                        "default": "{\"name\": \"xiaoming\", \"password\": \"123456\"}",
                        "description": "name(*用户名) password(*密码)",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"token生成成功!\", \"results\": \"tokenString\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"参数错误/token生成失败!/账号不存在或密码错误\", \"results\": null}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/test/": {
            "post": {
                "description": "描述",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "测试"
                ],
                "summary": "接口测试",
                "parameters": [
                    {
                        "description": "desc",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/v1.TParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.TResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "results": {
                                            "$ref": "#/definitions/v1.TParam"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.TResponse2"
                        }
                    }
                }
            }
        },
        "/users/": {
            "get": {
                "description": "暂时无权限/后需要登录",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "获取用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "页数",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "个数",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\":\"成功获取用户列表!\",\"results\":[{\"id\":1,\"nickname\":\"小明\",\"account\":\"xiaoming\",\"password\":\"123456\",\"email\":\"\",\"phone\":\"\",\"isAdmin\":false,\"photoFile\":\"\",\"createdTime\":\"\"}}]}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"参数错误/获取用户列表失败!\", \"results\": null}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "只有管理员可以添加用户",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "新增用户",
                "parameters": [
                    {
                        "description": "desc",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.AddUserParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "results": {
                                            "$ref": "#/definitions/models.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/users/{id}/": {
            "get": {
                "description": "登陆后才允许获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "获取用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\":\"成功获取用户!\",\"results\":{\"id\":1,\"nickname\":\"小明\",\"account\":\"xiaoming\",\"password\":\"123456\",\"email\":\"\",\"phone\":\"\",\"isAdmin\":false,\"photoFile\":\"\",\"createdTime\":\"\"}}}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"参数错误/用户不存在!\", \"results\": null}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "只有管理员/或用户自己可以修改",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "修改用户",
                "parameters": [
                    {
                        "type": "string",
                        "default": "1",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "desc",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.ModifyUserParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\":\"成功修改用户!\",\"results\":{\"id\":1,\"nickname\":\"小明\",\"account\":\"xiaoming\",\"password\":\"123456\",\"email\":\"\",\"phone\":\"\",\"isAdmin\":false,\"photoFile\":\"\",\"createdTime\":\"\"}}}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"参数错误/用户不存在!/用户修改失败!\", \"results\": null}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "只有管理员可以删除用户",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\":\"删除完成！\",\"results\": null}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"message\": \"参数错误/删除失败!\", \"results\": null}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/{id}/photo/": {
            "put": {
                "description": "只有管理员/或用户自己可以修改",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "上传用户头像",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "头像",
                        "name": "photoFile",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\":\"成功上传用户头像!\",\"results\":[{\"id\":1,\"nickname\":\"小明\",\"account\":\"xiaoming\",\"password\":\"123456\",\"email\":\"\",\"phone\":\"\",\"isAdmin\":false,\"photoFile\":\"\",\"createdTime\":\"\"}}]}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"参数错误/头像上传失败!\", \"results\": null}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/videos/": {
            "get": {
                "description": "无权限",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "视频"
                ],
                "summary": "获取视频信息列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "页数",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "个数",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\":\"成功获取视频列表!\",\"results\":[{\"id\":1,\"nickname\":\"小明\",\"account\":\"xiaoming\",\"password\":\"123456\",\"email\":\"\",\"phone\":\"\",\"isAdmin\":false,\"photoFile\":\"\",\"createdTime\":\"\"}}]}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"参数错误/获取视频列表失败!\", \"results\": null}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "普通用户可上传",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "视频"
                ],
                "summary": "上传视频信息",
                "parameters": [
                    {
                        "default": "{\"title\": \"小明大战小黄\", \"desc\": \"test\", \"category\": \"电影\", \"label\": \"小明，小黄，动作，武打\", \"rm\": \"mp4\"}",
                        "description": "title(*标题)",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\":\"创建视频完成!\",\"results\":{\"id\":1,\"nickname\":\"小明\",\"account\":\"xiaoming\",\"password\":\"123456\",\"email\":\"\",\"phone\":\"\",\"isAdmin\":false,\"photoFile\":\"\",\"createdTime\":\"\"}}}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"参数错误/创建失败!\", \"results\": null}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/videos/{id}/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "视频"
                ],
                "summary": "获取视频信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "视频id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\":\"成功获取视频信息!\",\"results\":{\"id\":1,\"nickname\":\"小明\",\"account\":\"xiaoming\",\"password\":\"123456\",\"email\":\"\",\"phone\":\"\",\"isAdmin\":false,\"photoFile\":\"\",\"createdTime\":\"\"}}}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"参数错误/视频不存在!\", \"results\": null}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "只有管理员/或用户自己可以修改视频信息",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "视频"
                ],
                "summary": "修改视频",
                "parameters": [
                    {
                        "default": "{\"title\": \"小明大战小黄\", \"desc\": \"test\", \"category\": \"电影\", \"label\": \"小明，小黄，动作，武打\", \"rm\": \"mp4\"}",
                        "description": "title(*标题)",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "default": "{\"nickname\": \"小明\", \"account\": \"xiaoming\", \"password\": \"123456\", \"email\": \"\", \"phone\": \"\"}",
                        "description": "用户信息",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\":\"成功修改用户!\",\"results\":{\"id\":1,\"nickname\":\"小明\",\"account\":\"xiaoming\",\"password\":\"123456\",\"email\":\"\",\"phone\":\"\",\"isAdmin\":false,\"photoFile\":\"\",\"createdTime\":\"\"}}}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"参数错误/用户不存在!/用户修改失败!\", \"results\": null}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "只有管理员或作者可以删除视频",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "视频"
                ],
                "summary": "删除视频",
                "parameters": [
                    {
                        "type": "string",
                        "description": "视频id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\":\"删除完成！\",\"results\": null}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"参数错误/删除失败!\", \"results\": null}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{id}/upload/": {
            "post": {
                "description": "管理员",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "视频"
                ],
                "summary": "审核视频",
                "parameters": [
                    {
                        "type": "string",
                        "description": "视频id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "是否通过",
                        "name": "pass",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\":\"视频审核完成\", \"results\":null}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"视频审核失败!\", \"results\": null}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.User": {
            "type": "object",
            "properties": {
                "account": {
                    "description": "账号",
                    "type": "string"
                },
                "createdTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "id": {
                    "description": "用户id",
                    "type": "integer"
                },
                "isAdmin": {
                    "description": "是否是管理员",
                    "type": "boolean"
                },
                "nickname": {
                    "description": "名称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "phone": {
                    "description": "手机",
                    "type": "string"
                },
                "photoFile": {
                    "description": "用户头像",
                    "type": "string"
                }
            }
        },
        "utils.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "返回说明",
                    "type": "string"
                },
                "results": {
                    "type": "object"
                }
            }
        },
        "v1.AddUserParam": {
            "type": "object",
            "required": [
                "account",
                "nickname",
                "password"
            ],
            "properties": {
                "account": {
                    "description": "账号",
                    "type": "string",
                    "example": "xiaom"
                },
                "createdTime": {
                    "description": "创建时间",
                    "type": "string",
                    "example": " "
                },
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": " "
                },
                "isAdmin": {
                    "description": "是否是管理员",
                    "type": "boolean",
                    "example": false
                },
                "nickname": {
                    "description": "名称",
                    "type": "string",
                    "example": "小明"
                },
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "123456"
                },
                "phone": {
                    "description": "手机",
                    "type": "string",
                    "example": " "
                },
                "photoFile": {
                    "description": "用户头像",
                    "type": "string",
                    "example": " "
                }
            }
        },
        "v1.ModifyUserParam": {
            "type": "object",
            "required": [
                "nickname",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": " "
                },
                "nickname": {
                    "description": "名称",
                    "type": "string",
                    "example": "xiaom"
                },
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "123456"
                },
                "phone": {
                    "description": "手机",
                    "type": "string",
                    "example": " "
                }
            }
        },
        "v1.TParam": {
            "type": "object",
            "required": [
                "a"
            ],
            "properties": {
                "a": {
                    "description": "字符串a",
                    "type": "string",
                    "example": "a"
                },
                "b": {
                    "description": "数字b",
                    "type": "integer",
                    "example": 2
                },
                "c": {
                    "description": "布尔c",
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "v1.TResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "成功返回说明",
                    "type": "string"
                },
                "results": {
                    "type": "object"
                }
            }
        },
        "v1.TResponse2": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "错误返回说明",
                    "type": "string"
                },
                "results": {
                    "type": "object"
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
	Host:        "127.0.0.1:9995",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "HelloC API (Swagger Example)",
	Description: "文档描述略",
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
