package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	Service
}

func New(serv Service) *gin.Engine {
	h := handler{
		serv,
	}

	r := gin.New()

	r.Static("/css", "templates/css/")
	r.LoadHTMLGlob("templates/index.html")

	r.GET("/order/:orderUid", h.GetOrderInfo)
	r.GET("/order/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	return r
}
