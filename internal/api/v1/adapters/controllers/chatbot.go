package controllers

import (
	"example.com/m/internal/api/v1/core/application/dto"
	"example.com/m/internal/api/v1/core/application/services/chat_bot_service"
	"github.com/gin-gonic/gin"
)

type ChatBotController struct {
	chatBotService chat_bot_service.ChatBotService
}

func NewChatBotController(chatBotService chat_bot_service.ChatBotService) *ChatBotController {
	return &ChatBotController{
		chatBotService: chatBotService,
	}
}

func (c *ChatBotController) SendMessage(ctx *gin.Context) {
	var message dto.ChatMessageDto
	if err := ctx.ShouldBindBodyWithJSON(&message); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// whatever..
}

func (c *ChatBotController) GetChat(ctx *gin.Context) {
	var query dto.ChatQueryDto
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// whatever..
}
