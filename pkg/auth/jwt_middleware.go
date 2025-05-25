package auth

import (
	"github.com/bingodfok/freshguard-boot/pkg/model/resp"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// JwtAuthMiddleware 验证请求中的 token
func JwtAuthMiddleware(excludes []string, keyFunc func() string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		for _, exclude := range excludes {
			if fiber.RoutePatternMatch(c.Path(), exclude) {
				return c.Next()
			}
		}
		// 从请求头中获取 token
		token := c.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(resp.EmptyDataResult(resp.UnauthorizedCode))
		}
		jwtToken := strings.Split(token, "Bearer ")[1]
		jwtAuth := JwtAuth{
			SigningKey: keyFunc(),
		}
		claims, err := jwtAuth.ParseToken(jwtToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(resp.EmptyDataResult(resp.UnauthorizedCode))
		}
		// 将用户信息存储在上下文中
		c.Locals("auth_context", claims)
		// 如果验证通过，继续处理请求
		return c.Next()
	}
}

func GetAuthContext(c *fiber.Ctx) *StandardClaims {
	return c.Locals("auth_context").(*StandardClaims)
}
