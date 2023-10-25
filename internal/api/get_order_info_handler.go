package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	response "wb-tech/internal/pkg/http"
)

func (h *handler) GetOrderInfo(c *gin.Context) {
	orderUid := c.Param("orderUid")

	orderData, err := h.Service.GetFromCache(orderUid)
	if err != nil {
		response.WriteErrorResponse(c, err)
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"order": orderData,
	})
}
