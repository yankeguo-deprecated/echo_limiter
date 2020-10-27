package echo_limiter

import (
	"github.com/labstack/echo/v4"
)

const (
	ContextKeyAvailable = "io.github.zionkit.echo_limiter/available"
)

func GetAvailable(ctx echo.Context) (val int64, ok bool) {
	val, ok = ctx.Get(ContextKeyAvailable).(int64)
	return
}

func New(l Limiter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			l.Take()
			defer l.Return()
			ctx.Set(ContextKeyAvailable, l.Available())
			return next(ctx)
		}
	}
}
