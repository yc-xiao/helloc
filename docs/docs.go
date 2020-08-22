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
                        "default": "{\"nickname\": \"小明\", \"account\": \"xiaom\", \"password\": \"123456\", \"email\": \"\", \"phone\": \"\", \"isAdmin\": false, \"photoFile\":\"\"}",
                        "description": "account(*账号), passowrd(*密码) ",
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
                        "description": "{\"message\":\"创建用户完成!\",\"results\":{\"id\":1,\"nickname\":\"小明\",\"account\":\"xiaoming\",\"password\":\"123456\",\"email\":\"\",\"phone\":\"\",\"isAdmin\":false,\"photoFile\":\"\",\"createdTime\":\"\"}}}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"参数错误/用户已存在，创建失败!\", \"results\": null}",
                        "schema": {
                            "type": "string"
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
                        "default": "{\"nickname\": \"小明\", \"password\": \"123456\", \"email\": \"\", \"phone\": \"\"}",
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
                        "description": "{\"msg\": \"参数错误/删除失败!\", \"results\": null}",
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