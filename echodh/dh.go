package echodh

import (
	"math/big"

	"github.com/labstack/echo/v4"
)

type DiffieHellman struct {
	g int64    // generator
	p *big.Int // modulus

	router *echo.Echo
}

func NewEchoDiffieHellman(p *big.Int, g int64) *DiffieHellman {
	return &DiffieHellman{
		p: p,
		g: g,
	}
}

// Use adds the middleware to the Echo engine.
func (p *DiffieHellman) Use(e *echo.Echo) {
	p.router = e
	p.router.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {

	})
}

func (p *DiffieHellman) handleGetGeneratorAndModulus(next echo.HandlerFunc) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		return next(echoCtx)
	}
}
