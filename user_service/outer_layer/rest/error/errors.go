package error

import (
	"net/http"
	domainErrors "shop/user_service/inner_layer/domain/errors"

	"github.com/gin-gonic/gin"
)

type MessageResponse struct {
	Message string `json:"message"`
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	errors := c.Errors
	if len(errors) > 0 {
		err, ok := errors[0].Err.(*domainErrors.AppError)
		if ok {
			message := MessageResponse{err.Error()}
			switch err.Type {
			case domainErrors.Unauthorized:
				c.JSON(http.StatusUnauthorized, message)
			case domainErrors.ValidationError:
				c.JSON(http.StatusBadRequest, message)
			case domainErrors.InternalServerError:
				c.JSON(http.StatusInternalServerError, message)
			case domainErrors.NotFound:
				c.JSON(http.StatusNotFound, message)
			default:
				c.JSON(http.StatusInternalServerError, MessageResponse{"Sorry, this case in development"})
			}

			return
		}

		c.JSON(http.StatusInternalServerError, MessageResponse{"Sorry, this case in development"})
	}
}
