package webserver

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strconv"
)

type FiberServer struct {
	AppName           string
	Port              uint
	ContextPath       string
	EnablePrintRoutes bool
	app               *fiber.App
	ctx               context.Context
}

// InitFiberServer 初始化fiber服务器
func (server *FiberServer) InitFiberServer() {
	server.app = fiber.New(fiber.Config{
		AppName:           server.AppName,
		ServerHeader:      "Fiber/v2",
		CaseSensitive:     true, // 区分请求地址大小写
		EnablePrintRoutes: true,
	})
}

// Route 添加fiber路由
func (server *FiberServer) Route(router func(fiber.Router)) {
	server.app.Route(server.ContextPath, router)
}

func (server *FiberServer) UseMiddleware() {
	cache.New(cache.Config{})
	server.app.Use(cors.New(cors.Config{}))
}

// Run 运行fiber服务器
func (server *FiberServer) Run() error {
	if server.app == nil {
		return fmt.Errorf("fiber server is not initialized")
	}
	return server.app.Listen(":" + strconv.Itoa(int(server.Port)))
}
