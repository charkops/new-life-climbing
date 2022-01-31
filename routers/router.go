package routers

import (
	"climbing/ginengine"
	"climbing/models"
	"climbing/utils"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type routerFunc struct {
	name   string
	weight int
	fn     func(router *gin.Engine)
}

type routerSlice []routerFunc

// Helper functions
func (r routerSlice) Len() int { return len(r) }

func (r routerSlice) Less(i, j int) bool { return r[i].weight > r[j].weight }

func (r routerSlice) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

// This var will hold all routers
var routers routerSlice

var once sync.Once

// register new router with key name
// key name is used to eliminate duplicate routes
// key name not case sensitive
func register(name string, f func(router *gin.Engine)) {
	registerWithWeight(name, 50, f)
}

// register new router with weight
func registerWithWeight(name string, weight int, f func(router *gin.Engine)) {
	if weight > 100 || weight < 0 {
		utils.CheckAndExit(errors.New(fmt.Sprintf("router weight must be >= 0 and <=100")))
	}

	for _, r := range routers {
		if strings.ToLower(name) == strings.ToLower(r.name) {
			utils.CheckAndExit(errors.New(fmt.Sprintf("router [%s] already register", r.name)))
		}
	}

	routers = append(routers, routerFunc{
		name:   name,
		weight: weight,
		fn:     f,
	})
}

func Setup() {
	once.Do(func() {
		sort.Sort(routers)
		for _, r := range routers {
			r.fn(ginengine.Engine)
			fmt.Printf("Loaded router [%s]\n", r.name)
		}
	})
}

// NOTE (@Charkops): Helper functions, not sure if needed
// for the fast return success result
func success() models.CommonResp {
	return models.CommonResp{
		Message:   "success",
		Timestamp: time.Now().Unix(),
	}
}

// for the fast return failed result
func failed(message string, args ...interface{}) models.CommonResp {
	return models.CommonResp{
		Message:   fmt.Sprintf(message, args...),
		Timestamp: time.Now().Unix(),
	}
}

// for the fast return result with custom data
func data(data interface{}) models.CommonResp {
	return models.CommonResp{
		Message:   "success",
		Timestamp: time.Now().Unix(),
		Data:      data,
	}
}
