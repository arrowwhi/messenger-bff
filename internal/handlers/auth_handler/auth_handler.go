package auth_handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	logger *zap.Logger

	//service
}

func (h *AuthHandler) Register(c *gin.Context) {

}
