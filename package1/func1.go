package package1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	loggerop "github.com/kiwi633/go-demo/logger"
	zap_log "github.com/kiwi633/go-demo/zap-log"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func someBusinessLogic() {
	loggerop.Info("开始处理业务逻辑")

	// 模拟一些操作
	time.Sleep(100 * time.Millisecond)

	loggerop.Info("业务逻辑处理成功")
}

// Handler示例
func UserHandler(c *gin.Context) {
	logger111 := zap_log.GetLoggerFromGin(c)
	logger111.Info("收到用户请求", zap.String("method", c.Request.Method), zap.String("path", c.Request.URL.Path))
	// 调用业务逻辑
	someBusinessLogic()

	if err := anotherBusinessLogic(c.Request.Context()); err != nil {
		loggerop.Error("处理用户请求失败", zap.Error(err))
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}

	logger111.Info("用户请求处理成功")
	c.JSON(200, gin.H{"message": "success", "tid": loggerop.GetTid()})
}
func anotherBusinessLogic(ctx context.Context) error {
	loggerop.Debug("进入另一个业务逻辑")

	// 模拟错误
	if false {
		loggerop.Error("业务逻辑处理失败", zap.Error(fmt.Errorf("some error")))
		return fmt.Errorf("处理失败")
	}

	loggerop.Info("另一个业务逻辑处理完成")
	return nil
}

func PersonList(c *gin.Context) {
	loggerop.Info("姓名：", zap.String("name", "suntong"))
	c.JSON(http.StatusOK, gin.H{"aaa": "accc"})
}
