// Code generated by protoc-gen-go-gin. DO NOT EDIT.
// versions:
// - protoc-gen-go-gin v1.0.0
// - protoc             v3.14.0
// source: proto/user.api.proto

package api

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	binding "github.com/gin-gonic/gin/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
)

var userServiceHTTP UserServiceHTTPServer

// UserServiceHTTPServer is the server API for UserService service.
type UserServiceHTTPServer interface {
	CreateUser(context.Context, *User) (*User, error)
	DeleteUser(context.Context, *UserId) (*emptypb.Empty, error)
	UpdateUser(context.Context, *UpdateUserReq) (*User, error)
	GetUser(context.Context, *UserId) (*User, error)
	ListUsers(context.Context, *ListUsersReq) (*ListUsersResp, error)
}

func userServiceCreateUser(c *gin.Context) {
	a := &User{}
	err := c.BindWith(a, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type")))
	type Status struct {
		Code    int32       `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	status := &Status{}
	if err != nil {
		status.Code = http.StatusBadRequest
		status.Message = err.Error()
		c.JSON(http.StatusBadRequest, status)
		return
	}
	resp, err := userServiceHTTP.CreateUser(c, a)
	if err != nil {
		status.Code = http.StatusInternalServerError
		status.Message = err.Error()
		c.JSON(http.StatusInternalServerError, status)
		return
	}
	status.Data = resp
	c.JSON(http.StatusOK, status)
}
func userServiceDeleteUser(c *gin.Context) {
	a := &UserId{}
	err := c.BindWith(a, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type")))
	type Status struct {
		Code    int32       `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	status := &Status{}
	if err != nil {
		status.Code = http.StatusBadRequest
		status.Message = err.Error()
		c.JSON(http.StatusBadRequest, status)
		return
	}
	resp, err := userServiceHTTP.DeleteUser(c, a)
	if err != nil {
		status.Code = http.StatusInternalServerError
		status.Message = err.Error()
		c.JSON(http.StatusInternalServerError, status)
		return
	}
	status.Data = resp
	c.JSON(http.StatusOK, status)
}
func userServiceUpdateUser(c *gin.Context) {
	a := &UpdateUserReq{}
	err := c.BindWith(a, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type")))
	type Status struct {
		Code    int32       `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	status := &Status{}
	if err != nil {
		status.Code = http.StatusBadRequest
		status.Message = err.Error()
		c.JSON(http.StatusBadRequest, status)
		return
	}
	resp, err := userServiceHTTP.UpdateUser(c, a)
	if err != nil {
		status.Code = http.StatusInternalServerError
		status.Message = err.Error()
		c.JSON(http.StatusInternalServerError, status)
		return
	}
	status.Data = resp
	c.JSON(http.StatusOK, status)
}
func userServiceGetUser(c *gin.Context) {
	a := &UserId{}
	err := c.BindWith(a, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type")))
	type Status struct {
		Code    int32       `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	status := &Status{}
	if err != nil {
		status.Code = http.StatusBadRequest
		status.Message = err.Error()
		c.JSON(http.StatusBadRequest, status)
		return
	}
	resp, err := userServiceHTTP.GetUser(c, a)
	if err != nil {
		status.Code = http.StatusInternalServerError
		status.Message = err.Error()
		c.JSON(http.StatusInternalServerError, status)
		return
	}
	status.Data = resp
	c.JSON(http.StatusOK, status)
}
func userServiceListUsers(c *gin.Context) {
	a := &ListUsersReq{}
	err := c.BindWith(a, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type")))
	type Status struct {
		Code    int32       `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	status := &Status{}
	if err != nil {
		status.Code = http.StatusBadRequest
		status.Message = err.Error()
		c.JSON(http.StatusBadRequest, status)
		return
	}
	resp, err := userServiceHTTP.ListUsers(c, a)
	if err != nil {
		status.Code = http.StatusInternalServerError
		status.Message = err.Error()
		c.JSON(http.StatusInternalServerError, status)
		return
	}
	status.Data = resp
	c.JSON(http.StatusOK, status)
}

const (
	PathUserServiceCreateUser = "/example.UserService/CreateUser"
	PathUserServiceDeleteUser = "/example.UserService/DeleteUser"
	PathUserServiceUpdateUser = "/example.UserService/UpdateUser"
	PathUserServiceGetUser    = "/example.UserService/GetUser"
	PathUserServiceListUsers  = "/example.UserService/ListUsers"
)

func RegisterUserServiceHTTP(e *gin.Engine, svr UserServiceHTTPServer, middleware map[string][]gin.HandlerFunc) {
	userServiceHTTP = svr
	e.POST(PathUserServiceCreateUser, append(middleware[PathUserServiceCreateUser], userServiceCreateUser)...)
	e.POST(PathUserServiceDeleteUser, append(middleware[PathUserServiceDeleteUser], userServiceDeleteUser)...)
	e.POST(PathUserServiceUpdateUser, append(middleware[PathUserServiceUpdateUser], userServiceUpdateUser)...)
	e.GET(PathUserServiceGetUser, append(middleware[PathUserServiceGetUser], userServiceGetUser)...)
	e.GET(PathUserServiceListUsers, append(middleware[PathUserServiceListUsers], userServiceListUsers)...)
}
