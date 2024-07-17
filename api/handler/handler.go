package handler

import (
	"API-Gateway/genproto/content"
	"API-Gateway/genproto/users"

	"go.uber.org/zap"
)

type Handler struct {
	ContentService     content.ContentClient
	Log                *zap.Logger
	UserService        users.UserServiceClient
}

func NewHandler(content content.ContentClient, l *zap.Logger, user users.UserServiceClient) *Handler {
	return &Handler{
		ContentService:     content,
		Log: l,
		UserService: user,
		
	}
}
