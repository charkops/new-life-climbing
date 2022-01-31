package middleware

import (
	"climbing/ginengine"
	"climbing/utils"
	"errors"
	"fmt"
	"sort"

	"github.com/gin-gonic/gin"
)

type middleware struct {
	HandlerFunc func() gin.HandlerFunc
	Weight      int
}

type middlewareSlice []middleware

var middlewares middlewareSlice

func (m middlewareSlice) Len() int { return len(m) }

func (m middlewareSlice) Less(i, j int) bool { return m[i].Weight > m[j].Weight }

func (m middlewareSlice) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

// Register new middleware
func register(fn func() gin.HandlerFunc) {
	mw := middleware{
		HandlerFunc: fn,
		Weight:      50,
	}
	middlewares = append(middlewares, mw)
}

// registering new middleware with weight
func registerWithWeight(weight int, handlerFunc func() gin.HandlerFunc) {

	if weight > 100 || weight < 0 {
		utils.CheckAndExit(errors.New(fmt.Sprintf("middleware weight must be >= 0 and <=100")))
	}

	mw := middleware{
		HandlerFunc: handlerFunc,
		Weight:      weight,
	}
	middlewares = append(middlewares, mw)
}

// framework init func
func Setup() {
	sort.Sort(middlewares)
	for _, mw := range middlewares {
		ginengine.Engine.Use(mw.HandlerFunc())
	}
	fmt.Println("Successfully loaded middlewares")
}
