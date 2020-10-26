package echo_limiter

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	ContextKeyAvailable = "io.github.zionkit.echo_limiter/available"
)

func GetAvailable(ctx echo.Context) (val int64, ok bool) {
	val, ok = ctx.Get(ContextKeyAvailable).(int64)
	return
}

type Options struct {
	Skipper middleware.Skipper
}

func New(l Limiter, opts Options) echo.MiddlewareFunc {
	if opts.Skipper == nil {
		opts.Skipper = middleware.DefaultSkipper
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if opts.Skipper(ctx) {
				return next(ctx)
			}
			l.Take()
			defer l.Return()
			ctx.Set(ContextKeyAvailable, l.Available())
			return next(ctx)
		}
	}
}
