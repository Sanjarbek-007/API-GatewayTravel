package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"API-Gateway/genproto"
	"API-Gateway/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

// Profile retrieves user profile details.
// @Summary Get user profile
// @Description Retrieve user profile details
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} genproto.ProfileRequest
// @Failure 401 {object} models.Failed
// @Failure 500 {object} models.Failed
// @Router /profile/{user_id} [get]
func (h *Handler) Profile(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	claims, err := h.ValidateToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Failed{Message: "Unauthorized", Error: err.Error()})
		return
	}

	if claims["sub"] != userID {
		ctx.JSON(http.StatusUnauthorized, models.Failed{Message: "Unauthorized access to profile"})
		return
	}

	request := &genproto.ProfileRequest{UserId: userID}
	response, err := h.UsersService.Profile(ctx, request)
	if err != nil {
		h.Log.Error("Failed to fetch profile", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Failed{Message: "Failed to fetch profile", Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// UpdateProfile updates user profile details.
// @Summary Update user profile
// @Description Update user profile details
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param input body models.UpdateProfileRequest true "Update details"
// @Success 200 {object} models.ProfileResponse
// @Failure 400 {object} models.Failed
// @Failure 401 {object} models.Failed
// @Failure 500 {object} models.Failed
// @Router /profile/{user_id} [put]
func (h *Handler) UpdateProfile(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	claims, err := h.ValidateToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Failed{Message: "Unauthorized", Error: err.Error()})
		return
	}

	if claims["sub"] != userID {
		ctx.JSON(http.StatusUnauthorized, models.Failed{Message: "Unauthorized access to profile"})
		return
	}

	var request models.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.Log.Error("Failed to bind JSON", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Failed{Message: "Invalid request payload", Error: err.Error()})
		return
	}

	request.UserID = userID
	response, err := h.UsersService.UpdateProfile(ctx, &genproto.UpdateProfileRequest{
		Id:       request.UserID,
		FullName: request.FullName,
		Bio:      request.Bio,
	})
	if err != nil {
		h.Log.Error("Failed to update profile", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Failed{Message: "Failed to update profile", Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// ValidateToken validates the JWT token from Authorization header.
func (h *Handler) ValidateToken(ctx *gin.Context) (jwt.MapClaims, error) {
	tokenString := ctx.GetHeader("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("salom"), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
