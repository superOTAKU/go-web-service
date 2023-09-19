package cmd

import (
	"go-web-service/pkg/logging"
	"go-web-service/pkg/repo"
	"go-web-service/pkg/rest"
	"net/http"

	"github.com/gin-gonic/gin"
)

var getUsersCommand = &Command{
	Handler: getUsers,
	Route: RouteInfo{
		Method:             "GET",
		Path:               "/user",
		NeedAuthentication: false,
	},
}

func getUsers(c *gin.Context) {
	users, err := repo.Users.FindAll()
	if err != nil {
		rest.RtnErr(c, http.StatusInternalServerError, rest.Unknown)
		return
	}
	logging.Logger.Info("get users success", "users", &users)
	c.JSON(http.StatusOK, users)
}
