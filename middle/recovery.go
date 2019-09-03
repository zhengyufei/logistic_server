package middle

import (
	"github.com/gin-gonic/gin"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/log"
	"runtime/debug"
)

// 一个gin框架的中间件，记录panic及其调用栈到日志中然后再原封不动重新抛出panic。
func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// todo
			log.Infof("panic occurred:\n%v\n%s", err, debug.Stack())
			panic(err)
		}
	}()
	c.Next()
}
