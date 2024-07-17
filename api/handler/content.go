package handler

import (
	"API-Gateway/api/auth"
	pb "API-Gateway/genproto/content"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateStory godoc
// @Security ApiKeyAuth
// @Summary create story
// @Description create new story
// @Tags stories
// @Param info body content.CreateStoriesRequest true "info"
// @Success 200 {object} content.CreateStoriesResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories [post]
func (h *Handler) CreateStory(c *gin.Context) {
	h.Log.Info("CreateStory started")
	req := pb.CreateStoriesRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	accessToken := c.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req.UserId = id
	res, err := h.ContentService.CreateStories(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("CreateStory ended")
}

// UpdateStory godoc
// @Security ApiKeyAuth
// @Summary Update story
// @Description Update new story
// @Tags stories
// @Param story_id path string true "story_id"
// @Param info body content.UpdateStoriesReq true "info"
// @Success 200 {object} content.UpdateStoriesRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id} [put]
func (h *Handler) UpdateStory(c *gin.Context) {
	h.Log.Info("UpdateStory started")
	req := pb.UpdateStoriesReq{}
	if err := c.BindJSON(&req); err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req.Id = c.Param("story_id")
	res, err := h.ContentService.UpdateStories(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("UpdateStory ended")
}

// DeleteStory godoc
// @Security ApiKeyAuth
// @Summary Delete story
// @Description Delete new story
// @Tags stories
// @Param story_id path string true "story_id"
// @Success 200 {object} string "succesfully deleted"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id} [delete]
func (h *Handler) DeleteStory(c *gin.Context) {
	h.Log.Info("DeleteStory started")
	req := pb.StoryId{}
	req.Id = c.Param("story_id")
	_, err := h.ContentService.DeleteStories(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "succesfully deleted"})
	h.Log.Info("DeleteStory ended")
}

// GetAllStories godoc
// @Security ApiKeyAuth
// @Summary get all story
// @Description get all stories
// @Tags stories
// @Param limit query string false "Number of stories to fetch"
// @Param offset query string false "Number of stories to omit"
// @Success 200 {object} content.GetAllStoriesRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories [get]
func (h *Handler) GetAllStories(c *gin.Context) {
	h.Log.Info("GetAllStories started")
	req := pb.GetAllStoriesReq{}
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		req.Limit = int64(limit)
	} else {
		req.Limit = 10
	}

	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		req.Offset = int64(offset)
	} else {
		req.Offset = 0
	}

	res, err := h.ContentService.GetAllStories(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetAllStories ended")
}

// GetStory godoc
// @Security ApiKeyAuth
// @Summary Get story
// @Description Get story by id
// @Tags stories
// @Param story_id path string true "story_id"
// @Success 200 {object} content.GetStoryRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id} [get]
func (h *Handler) GetStory(c *gin.Context) {
	h.Log.Info("GetStory started")
	req := pb.StoryId{}
	req.Id = c.Param("story_id")
	if len(req.Id) <= 0 {

		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})

	}
	res, err := h.ContentService.GetStory(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetStory ended")
}

// CommentStory godoc
// @Security ApiKeyAuth
// @Summary comment story
// @Description comment to story
// @Tags stories
// @Param story_id path string true "story_id"
// @Param info body content.CommentStoryReq true "story_id"
// @Success 200 {object} content.CommentStoryRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id}/comments [post]
func (h *Handler) CommentStory(c *gin.Context) {
	h.Log.Info("CommentStory started")
	accessToken := c.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req := pb.CommentStoryReq{}
	c.BindJSON(&req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req.AuthorId = id
	req.StoryId = c.Param("story_id")
	if len(req.StoryId) <= 0 {

		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})

	}
	res, err := h.ContentService.CommentStory(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("CommentStory ended")
}

// GetCommentsOfStory godoc
// @Security ApiKeyAuth
// @Summary comment of story
// @Description get comment of story
// @Tags stories
// @Param story_id path string true "story_id"
// @Param limit query string false "Number of stories to fetch"
// @Param offset query string false "Number of stories to omit"
// @Success 200 {object} content.GetCommentsOfStoryRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id}/comments [get]
func (h *Handler) GetCommentsOfStory(c *gin.Context) {
	h.Log.Info("GetCommentsOfStory started")
	req := pb.GetCommentsOfStoryReq{}
	req.StoryId = c.Param("story_id")
	if len(req.StoryId) <= 0 {

		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})

	}
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		req.Limit = int64(limit)
	} else {
		req.Limit = 10
	}

	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		req.Offset = int64(offset)
	} else {
		req.Offset = 0
	}

	res, err := h.ContentService.GetCommentsOfStory(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetCommentsOfStory ended")
}

// Like godoc
// @Security ApiKeyAuth
// @Summary comment story
// @Description comment to story
// @Tags stories
// @Param story_id path string true "story_id"
// @Success 200 {object} content.LikeRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id}/like [post]
func (h *Handler) Like(c *gin.Context) {
	h.Log.Info("Like started")
	accessToken := c.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req := pb.LikeReq{}
	req.UserId = id
	req.StoryId = c.Param("story_id")
	if len(req.StoryId) <= 0 {
		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
	}
	res, err := h.ContentService.Like(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("Like ended")
}

// Itineraries godoc
// @Security ApiKeyAuth
// @Summary create
// @Description create itineraries
// @Tags itineraries
// @Param info body content.ItinerariesReq true "info"
// @Success 200 {object} content.ItinerariesRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/itineraries [post]
func (h *Handler) Itineraries(c *gin.Context) {
	h.Log.Info("Itineraries started")
	accessToken := c.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req := pb.ItinerariesReq{}
	err = c.BindJSON(&req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req.UserId = id
	res, err := h.ContentService.Itineraries(c, &req)
	if err != nil {
		h.Log.Error("there")
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("Itineraries ended")
}

// UpdateItineraries godoc
// @Security ApiKeyAuth
// @Summary update
// @Description update itineraries
// @Tags itineraries
// @Param itinerary_id path string true "itinerary_id"
// @Param info body content.UpdateItinerariesReq true "info"
// @Success 200 {object} content.ItinerariesRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/itineraries/{itinerary_id} [put]
func (h *Handler) UpdateItineraries(c *gin.Context) {
	h.Log.Info("UpdateItineraries started")
	req := pb.UpdateItinerariesReq{}
	err := c.BindJSON(&req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req.Id = c.Param("itinerary_id")
	if len(req.Id) <= 0 {

		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})

	}
	res, err := h.ContentService.UpdateItineraries(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("UpdateItineraries ended")
}

// DeleteItineraries godoc
// @Security ApiKeyAuth
// @Summary Delete
// @Description Delete itineraries
// @Tags itineraries
// @Param itinerary_id path string true "itinerary_id"
// @Success 200 {object} string "successfully deleted"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/itineraries/{itinerary_id} [delete]
func (h *Handler) DeleteItineraries(c *gin.Context) {
	h.Log.Info("DeleteItineraries started")
	req := pb.StoryId{}
	req.Id = c.Param("itinerary_id")
	if len(req.Id) <= 0 {

		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})

	}
	_, err := h.ContentService.DeleteItineraries(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "successfully deleted"})
	h.Log.Info("DeleteItineraries ended")
}

// GetItineraries godoc
// @Security ApiKeyAuth
// @Summary Get
// @Description Get itineraries
// @Tags itineraries
// @Param limit query string false "Number of itineraries to fetch"
// @Param offset query string false "Number of itineraries to omit"
// @Success 200 {object} content.GetItinerariesRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/itineraries [get]
func (h *Handler) GetItineraries(c *gin.Context) {
	h.Log.Info("GetItineraries started")
	req := pb.GetItinerariesReq{}
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		req.Limit = int64(limit)
	} else {
		req.Limit = 10
	}

	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		req.Offset = int64(offset)
	} else {
		req.Offset = 0
	}

	res, err := h.ContentService.GetItineraries(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetItineraries ended")
}

// GetItinerariesById godoc
// @Security ApiKeyAuth
// @Summary Get
// @Description Get itineraries by id
// @Tags itineraries
// @Param itinerary_id path string true "itinerary_id"
// @Success 200 {object} content.GetItinerariesByIdRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/itineraries/{itinerary_id} [get]
func (h *Handler) GetItinerariesById(c *gin.Context) {
	h.Log.Info("GetItinerariesById started")
	req := pb.StoryId{}
	req.Id = c.Param("itinerary_id")
	if len(req.Id) <= 0 {
		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
	}
	res, err := h.ContentService.GetItinerariesById(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetItinerariesById ended")
}

// CommentItineraries godoc
// @Security ApiKeyAuth
// @Summary comment
// @Description comment itineraries
// @Tags itineraries
// @Param itinerary_id path string true "itinerary_id"
// @Param info body content.CommentItinerariesReq true "info"
// @Success 200 {object} content.CommentItinerariesRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/itineraries/{itinerary_id}/comments [post]
func (h *Handler) CommentItineraries(c *gin.Context) {
	h.Log.Info("CommentItineraries started")
	accessToken := c.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req := pb.CommentItinerariesReq{}
	err = c.BindJSON(&req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req.ItineraryId = c.Param("itinerary_id")
	if len(req.ItineraryId) <= 0 {
		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
	}
	req.AuthorId = id
	res, err := h.ContentService.CommentItineraries(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("CommentItineraries ended")
}

// GetDestinations godoc
// @Security ApiKeyAuth
// @Summary get
// @Description get destination
// @Tags destinations
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Param name query string false "name"
// @Success 200 {object} content.GetDestinationsRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/destinations [get]
func (h *Handler) GetDestinations(c *gin.Context) {
	h.Log.Info("GetDestinations started")
	req := pb.GetDestinationsReq{}
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	req.Name = c.Query("name")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		req.Limit = int64(limit)
	} else {
		req.Limit = 10
	}

	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		req.Offset = int64(offset)
	} else {
		req.Offset = 0
	}

	res, err := h.ContentService.GetDestinations(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetDestinations ended")
}

// GetDestinationsById godoc
// @Security ApiKeyAuth
// @Summary Get
// @Description Get destination by id
// @Tags destinations
// @Param destination_id path string true "destination_id"
// @Success 200 {object} content.GetDestinationsByIdRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/destinations/{destination_id} [get]
func (h *Handler) GetDestinationsById(c *gin.Context) {
	h.Log.Info("GetDestinationsById started")
	req := pb.GetDestinationsByIdReq{}
	req.Id = c.Param("destination_id")
	if len(req.Id) <= 0 {
		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
	}
	res, err := h.ContentService.GetDestinationsById(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetDestinationsById ended")
}

// SendMessage godoc
// @Security ApiKeyAuth
// @Summary Send Message
// @Description Send Message
// @Tags message
// @Param info body content.SendMessageReq true "info"
// @Success 200 {object} content.SendMessageRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/messages [post]
func (h *Handler) SendMessage(c *gin.Context) {
	h.Log.Info("SendMessage started")
	accessToken := c.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req := pb.SendMessageReq{}
	err = c.BindJSON(&req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req.UserId = id
	res, err := h.ContentService.SendMessage(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, &res)
	h.Log.Info("SendMessage ended")
}

// GetMessages godoc
// @Security ApiKeyAuth
// @Summary get Message
// @Description get Message
// @Tags message
// @Param limit query string false "Number of messages to fetch"
// @Param offset query string false "Number of messages to omit"
// @Success 200 {object} content.GetMessagesRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/messages [get]
func (h *Handler) GetMessages(c *gin.Context) {
	h.Log.Info("GetMessages started")
	req := pb.GetMessagesReq{}
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		req.Limit = int64(limit)
	} else {
		req.Limit = 10
	}

	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		req.Offset = int64(offset)
	} else {
		req.Offset = 0
	}

	res, err := h.ContentService.GetMessages(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, &res)
	h.Log.Info("GetMessages ended")
}

// CreateTips godoc
// @Security ApiKeyAuth
// @Summary create
// @Description create tips
// @Tags tips
// @Param info body content.CreateTipsReq true "destination_id"
// @Success 200 {object} content.CreateTipsRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/travel-tips [post]
func (h *Handler) CreateTips(c *gin.Context) {
	h.Log.Info("CreateTips started")
	accessToken := c.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req := pb.CreateTipsReq{}
	err = c.BindJSON(&req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req.UserId = id
	res, err := h.ContentService.CreateTips(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("CreateTips ended")
}

// GetTips godoc
// @Security ApiKeyAuth
// @Summary get
// @Description get tips
// @Tags tips
// @Param limit query string false "Number of messages to fetch"
// @Param offset query string false "Number of messages to omit"
// @Param category query string false "category"
// @Success 200 {object} content.GetTipsRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/travel-tips [get]
func (h *Handler) GetTips(c *gin.Context) {
	h.Log.Info("GetTips started")

	req := pb.GetTipsReq{}
	req.Category = c.Query("category")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		req.Limit = int64(limit)
	} else {
		req.Limit = 10
	}

	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		req.Offset = int64(offset)
	} else {
		req.Offset = 0
	}

	res, err := h.ContentService.GetTips(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetTips ended")
}

// GetUserStat godoc
// @Security ApiKeyAuth
// @Summary best
// @Description get user
// @Tags users
// @Param user_id path string true "user_id"
// @Success 200 {object} content.GetUserStatRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/users/{user_id}/statistics [get]
func (h *Handler) GetUserStat(c *gin.Context) {
	h.Log.Info("GetUserStat started")
	req := pb.GetUserStatReq{}
	req.UserId = c.Param("user_id")
	if len(req.UserId) <= 0 {
		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
	}
	res, err := h.ContentService.GetUserStat(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetUserStat ended")
}

// TopDestinations godoc
// @Security ApiKeyAuth
// @Summary top places
// @Description get top places
// @Tags top
// @Success 200 {object} content.Answer
// @Failure 500 {object} string "Server error"
// @Router /api/v1/trending-destinations [get]
func (h *Handler) TopDestinations(c *gin.Context) {
	h.Log.Info("TopDestinations started")
	req := pb.Void{}
	res, err := h.ContentService.TopDestinations(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	fmt.Println(res)
	c.JSON(200, &res)
	h.Log.Info("TopDestinations ended")
}
