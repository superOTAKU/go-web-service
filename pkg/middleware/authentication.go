package middleware

import (
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context) {
	// 查询元数据，确定是否鉴权
	metaData := c.MustGet(metaDataKey).(*MetaData)
	log.Info("authentication middleware invoked", "needAuthentication", metaData.NeedAuthentication)
	// TODO: 读取header中的JWT Token并校验
	c.Next()
}
