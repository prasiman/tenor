package maincontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetServiceDetail(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<html>Tenor App</html>"))
}
