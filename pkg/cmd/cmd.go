package cmd

import (
	"fmt"
	"go-web-service/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type Command struct {
	Handler func(*gin.Context)
	Route   RouteInfo
}

// 路由元数据，类似SpringMVC作用在Controller上的注解
type RouteInfo struct {
	Method             string
	Path               string
	NeedAuthentication bool
	Permission         string
}

var commands = []*Command{}

func registerCommand(c *Command) {
	commands = append(commands, c)
}

func registerAllCommands() {
	registerCommand(getUsersCommand)
}

func PublishCommands(router *gin.Engine) {
	registerAllCommands()
	for _, c := range commands {
		handler, route := c.Handler, c.Route
		handlers := make([]gin.HandlerFunc, 0, 64)
		handlers = append(handlers, middleware.MetaDataMiddleware(&middleware.MetaData{
			NeedAuthentication: route.NeedAuthentication,
			Permission:         route.Permission,
		}))
		handlers = append(handlers, middleware.Authentication)
		handlers = append(handlers, handler)
		router.Handle(route.Method, route.Path, handlers...)
	}
	fmt.Println("publish commands finish")
}
