package webserver

import (
	"context"
	"errors"
	"fmt"
	bizErrors "github.com/bingodfok/freshguard-boot/pkg/model/errors"
	"github.com/bingodfok/freshguard-boot/pkg/model/resp"
	"github.com/gofiber/fiber/v2"
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
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			var biz *bizErrors.BizError
			if errors.As(err, &biz) {
				return ctx.JSON(resp.NewResultCode(biz.Code, biz.Msg))
			}
			var fiberError *fiber.Error
			if errors.As(err, &fiberError) {
				return ctx.Status(fiberError.Code).JSON(resp.NewResultCode(fiberError.Code, fiberError.Message))
			}
			return ctx.Status(fiber.StatusInternalServerError).JSON(resp.EmptyDataResult(resp.ServerErrorCode))
		},
	})
	if server.ContextPath == "" {
		server.ContextPath = "/"
	}
}

// Route 添加fiber路由
func (server *FiberServer) Route(fun func(router fiber.Router)) {
	server.app.Route(server.ContextPath, fun)
}

func (server *FiberServer) UseMiddleware(handlers ...fiber.Handler) {
	for _, handler := range handlers {
		server.app.Use(handler)
	}
}

// Run 运行fiber服务器
func (server *FiberServer) Run() error {
	if server.app == nil {
		return fmt.Errorf("fiber server is not initialized")
	}
	return server.app.Listen(":" + strconv.Itoa(int(server.Port)))
}
