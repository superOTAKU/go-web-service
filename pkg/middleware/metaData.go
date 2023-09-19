package middleware

import (
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

type MetaData struct {
	NeedAuthentication bool
	Permission         string
}

const (
	metaDataKey = "meta"
)

func MetaDataMiddleware(metaData *MetaData) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(metaDataKey, metaData)
		log.Info("meta set", "path", c.Request.URL.Path, "meta", metaData)
		c.Next()
	}
}
