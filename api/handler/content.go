package handlers

import (
	pb "API-Gateway/genproto"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var logger *zap.Logger

// CreateStory handles the creation of a new Story.
// @Summary Create Story
// @Description Create a new Story
// @Tags Story
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body genproto.CrateStoryRequest true "Create Story"
// @
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/createStory [post]
func (h *Handler) CreateStory(ctx *gin.Context) {
	request := pb.CrateStoryRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		logger.Error("error in CreateStory")
		return
	}

	if request.Location == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "it is not full information"})
		return
	}

	_, err := uuid.Parse(request.AuthorId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("wrong id"))
		h.Log.Error("error")
		return
	}

	_, err = h.ContentService.CrateStory(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "errorncreating story in handler",
		})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Story Created",
	})
}

// UpdateStory handles the update of a story.
// @Summary Update Story
// @Description Update an existing story
// @Tags Story
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Story ID"
// @Param Update body genproto.UpdateStoryRequest true "Update Story"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/updateStory/{id} [put]
func (h *Handler) UpdateStory(ctx *gin.Context) {
	request := pb.UpdateStoryRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		logger.Error("error in CreateStory")
		return
	}
	request.StoryId = ctx.Param("id")

	_, err := uuid.Parse(request.StoryId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("wrong id"))
		h.Log.Error("error")
		return
	}

	if request.Title == "" || request.Content == "" {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("it is not full information"))
		h.Log.Error("error")
		return
	}
	_, err = h.ContentService.UpdateStory(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("error in Gateway UpdateStory"))
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Story Updated",
	})
}

// DeleteStoryHandler handles the deletion of a story.
// @Summary Delete Story
// @Description Delete an existing story
// @Tags Story
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "story ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/deleteStory/{id} [delete]
func (h *Handler) DeleteStory(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("wrong id"))
		h.Log.Error("error")
		return
	}

	_, err = h.ContentService.DeleteStory(ctx, &pb.DeleteStoryRequest{StoryId: id})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error in Gateway DeleteStory"})
		logger.Error("error")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Story Deleted",
	})
}

// GetAllStoriesHandler retrieves a list of stories with optional filtering and pagination.
// @Summary Get All Stories
// @Description Retrieve a list of Stories with optional filtering and pagination.
// @Tags Story
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param story_id query string false "Filter by story_id"
// @Param limit query int false "Number of items to return"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} genproto.MenusResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/get_all [get]
func (h *Handler) GetAllStories(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	request := pb.GetAllStoriesRequest{
		AuthorId: ctx.Query("name"),
		Limit:    int32(limit),
		Offset:   int32(offset),
	}

	resp, err := h.ContentService.GetAllStories(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in Gateway GetAllStories"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
