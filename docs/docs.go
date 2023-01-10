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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "用户登录",
                "tags": [
                    "公共方法"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{“code”: \"200\", \"msg\":\"\", \"data\": \"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/problem-detail": {
            "get": {
                "description": "获取问题详细信息",
                "tags": [
                    "公共方法"
                ],
                "summary": "问题详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "问题的唯一标识",
                        "name": "identity",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{“code”: \"200\", \"msg\":\"\", \"data\": \"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/problem-list": {
            "get": {
                "description": "获取问题列表",
                "tags": [
                    "公共方法"
                ],
                "summary": "获取问题列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "请输入当前页,默认第一页",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页多少条数据,默认20条",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "查询的关键字",
                        "name": "keyWord",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "分类的唯一标识",
                        "name": "category_identity",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{“code”: \"200\", \"msg\":\"\", \"data\": \"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "添加问题",
                "tags": [
                    "公共方法"
                ],
                "summary": "添加一个问题",
                "responses": {
                    "200": {
                        "description": "{“code”: \"200\", \"msg\":\"\", \"data\": \"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "用户注册",
                "tags": [
                    "公共方法"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "确认密码",
                        "name": "confirm_password",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "mail",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "验证码",
                        "name": "code",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{“code”: \"200\", \"msg\":\"\", \"data\": \"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/send-code": {
            "post": {
                "description": "发送邮箱验证码",
                "tags": [
                    "公共方法"
                ],
                "summary": "发送邮箱验证码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户邮箱",
                        "name": "email",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{“code”: \"200\", \"msg\":\"\", \"data\": \"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/submit-list": {
            "get": {
                "description": "获取问题列表",
                "tags": [
                    "公共方法"
                ],
                "summary": "获取提交记录列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "请输入当前页,默认第一页",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页多少条数据,默认20条",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "问题的唯一标识",
                        "name": "problem_identity",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "用户的唯一标识",
                        "name": "user_identity",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "提交的状态【-1-待判断，1-答案正确，2-答案错误，3-运行超时，4-运行超内存】",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{“code”: \"200\", \"msg\":\"\", \"data\": \"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user-detail": {
            "get": {
                "description": "获取用户详细信息",
                "tags": [
                    "公共方法"
                ],
                "summary": "获取用户详细信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户的唯一标识",
                        "name": "identity",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{“code”: \"200\", \"msg\":\"\", \"data\": \"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user-list": {
            "get": {
                "description": "获取用户列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "获取所有用户",
                "responses": {
                    "200": {
                        "description": "{“code”: \"200\", \"msg\":\"\", \"data\": \"\"}",
                        "schema": {
                            "type": "string"
                        }
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
