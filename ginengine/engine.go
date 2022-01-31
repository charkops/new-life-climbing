package ginengine

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine
var engineOnce sync.Once

// Init gin router engine
func Setup() {
	engineOnce.Do(func() {
		// TODO (@Charkops): Add if/else for debug/release mode
		// E.X:
		// if conf.Basic.Debug {
		// 	gin.SetMode(gin.DebugMode)
		// } else {
		// 	gin.SetMode(gin.ReleaseMode)
		// }

		Engine = gin.New()
	})
}
