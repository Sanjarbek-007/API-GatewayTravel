package handlers

import (
	"API-Gateway/genproto"
	"go.uber.org/zap"

)

type Handler struct {
	ContentService     genproto.ContentServiceClient
	// UsersService       genproto.UserServiceClient
	Log                *zap.Logger
}

func NewHandler(content genproto.ContentServiceClient, user genproto.UserServiceClient, l *zap.Logger) *Handler {
	return &Handler{
		ContentService:     content,
		// UsersService:       user,
		Log: l,
		
	}
}
