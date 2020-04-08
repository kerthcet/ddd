package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"ddd/application/service"
	"ddd/infrastructure/common/response_code/code"
	"ddd/interface/dto/rest"
	"ddd/interface/factory"
)

// Leave 请假单
func Leave(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
			SendResponse(c, code.Code_UNKNOWN, r)
		}
	}()

	var leave rest.CreateLeaveRequestDTO
	if err := c.ShouldBindJSON(&leave); err != nil {
		log.Error(err)
		SendResponse(c, code.Code_INVALID_ARGUMENT, err)
		return
	}

	data, err := service.CreateLeave(factory.ToCreateLeaveDO(&leave))
	if err != nil {
		SendResponse(c, code.Code_UNKNOWN, err)
		return
	}
	SendResponse(c, nil, factory.ToCreateLeaveDTO(data))
}
