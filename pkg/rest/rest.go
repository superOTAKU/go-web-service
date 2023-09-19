package rest

import "github.com/gin-gonic/gin"

const (
	Unknown = "unknown"
)

type Err struct {
	code string
	msg  string
}

func RtnErrMsg(c *gin.Context, status int, code string, msg string) {
	c.JSON(status, &Err{code: code, msg: msg})
}

func RtnErr(c *gin.Context, status int, code string) {
	RtnErrMsg(c, status, code, "")
}
