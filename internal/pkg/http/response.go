package http

import (
	errors "errors"
	"log"
	"net/http"

	"wb-tech/internal/pkg/app_err"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error `json:"error"`
}

type Error struct {
	Message string `json:"message"`
}

func WriteErrorResponse(c *gin.Context, err error) {
	var bErr app_err.BusinessError

	if errors.As(err, &bErr) {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error": bErr.Error(),
		})
		//	c.JSON(http.StatusBadRequest, errorResponse)
	} else {
		errorResponse := ErrorResponse{
			Error: Error{
				Message: "Something went wrong, try again",
			},
		}

		log.Println(err)

		c.HTML(http.StatusInternalServerError, "index.html", errorResponse)
	}
}
