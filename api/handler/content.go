package handlers

import (
	pb "API-Gateway/genproto"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


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
		h.Log.Error("error in CreateStory")
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
		h.Log.Error("error in CreateStory")
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
		h.Log.Error("error")
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
// @Router /content/getAllStories [get]
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

// StoryFullInfoHandler retrieves full information about a specific story.
// @Summary Get Story Full Info
// @Description Retrieve full information about a specific story.
// @Tags Story
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param story_id path string true "Story ID"
// @Success 200 {object} genproto.StoryFullInfoResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/storyFullInfo/{story_id} [get]
func (h *Handler) StoryFullInfo(ctx *gin.Context) {
	storyID := ctx.Param("story_id")

	_, err := uuid.Parse(storyID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("wrong story_id"))
		h.Log.Error("error")
		return
	}

	resp, err := h.ContentService.StoryFullInfo(ctx, &pb.StoryFullInfoRequest{StoryId: storyID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in Gateway StoryFullInfo"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// CommentStoryHandler handles commenting on a story.
// @Summary Comment on Story
// @Description Comment on a specific story.
// @Tags Story
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param story_id path string true "Story ID"
// @Param Create body genproto.CommentStoryRequest true "Comment on Story"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/commentStory/{story_id} [post]
func (h *Handler) CommentStory(ctx *gin.Context) {
	storyID := ctx.Param("story_id")

	_, err := uuid.Parse(storyID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("wrong story_id"))
		h.Log.Error("error")
		return
	}

	request := pb.CommentStoryRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		h.Log.Error("error")
		return
	}
	request.StoryId = storyID

	_, err = h.ContentService.CommentStory(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in Gateway CommentStory"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Comment added to Story",
	})
}

// GetAllCommentsHandler retrieves all comments for a specific story.
// @Summary Get All Comments
// @Description Retrieve all comments for a specific story.
// @Tags Story
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param story_id path string true "Story ID"
// @Param limit query int false "Number of items to return"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} genproto.GetAllCommentResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/getAllComments/{story_id} [get]
func (h *Handler) getGetAllComments(ctx *gin.Context) {
	storyID := ctx.Param("story_id")

	_, err := uuid.Parse(storyID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("wrong story_id"))
		h.Log.Error("error")
		return
	}

	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	request := pb.GetAllCommentRequest{
		StoryId: storyID,
		Limit:   int32(limit),
		Offset:  int32(offset),
	}

	resp, err := h.ContentService.GetAllComments(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in Gateway GetAllComments"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusOK, resp)
}


// CreateLikeHandler handles creating a like for a story.
// @Summary Create Like
// @Description Create a like for a specific story.
// @Tags Story
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param story_id path string true "Story ID"
// @Param Create body genproto.CreateLikeRequest true "Create Like"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/createLike/{story_id} [post]
func (h *Handler) CreateLike(ctx *gin.Context) {
	storyID := ctx.Param("story_id")

	_, err := uuid.Parse(storyID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("wrong story_id"))
		h.Log.Error("error")
		return
	}

	request := pb.CreateLikeRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		h.Log.Error("error")
		return
	}
	request.StoryId = storyID

	_, err = h.ContentService.CreateLike(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in Gateway CreateLike"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Like added to Story",
	})
}

// CreateItinerariesHandler handles creating an itinerary.
// @Summary Create Itinerary
// @Description Create a new itinerary.
// @Tags Itinerary
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body genproto.CreateItinerariesRequest true "Create Itinerary"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/createItineraries [post]
func (h *Handler) CreateItineraries(ctx *gin.Context) {
	request := pb.CreateItinerariesRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		h.Log.Error("error")
		return
	}

	if request.Title == "" || request.Description == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "it is not full information"})
		return
	}

	_, err := h.ContentService.CreateItineraries(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error creating itinerary in handler"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Itinerary Created",
	})
}

// UpdateItinerariesHandler handles updating an itinerary.
// @Summary Update Itinerary
// @Description Update an existing itinerary.
// @Tags Itinerary
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Itinerary ID"
// @Param Update body genproto.UpdateItinerariesRequest true "Update Itinerary"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/updateItineraries/{id} [put]
func (h *Handler) UpdateItineraries(ctx *gin.Context) {
	request := pb.UpdateItinerariesRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		h.Log.Error("error")
		return
	}
	request.ItineraryId = ctx.Param("id")

	_, err := uuid.Parse(request.ItineraryId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("wrong id"))
		h.Log.Error("error")
		return
	}

	if request.Title == "" || request.Description == "" {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("it is not full information"))
		h.Log.Error("error")
		return
	}

	_, err = h.ContentService.UpdateItineraries(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("error in Gateway UpdateItineraries"))
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Itinerary Updated",
	})
}

// DeleteItinerariesHandler handles deleting an itinerary.
// @Summary Delete Itinerary
// @Description Delete an existing itinerary.
// @Tags Itinerary
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Itinerary ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/deleteItineraries/{id} [delete]
func (h *Handler) DeleteItineraries(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("wrong id"))
		h.Log.Error("error")
		return
	}

	_, err = h.ContentService.DeleteItineraries(ctx, &pb.DeleteItinerariesRequest{ItineraryId: id})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error in Gateway DeleteItineraries"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Itinerary Deleted",
	})
}

// GetAllItinerariesHandler retrieves all itineraries with optional filtering and pagination.
// @Summary Get All Itineraries
// @Description Retrieve all itineraries with optional filtering and pagination.
// @Tags Itinerary
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param itinerary_id query string false "Filter by itinerary_id"
// @Param limit query int false "Number of items to return"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} genproto.GetAllItinerariesResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/getAllItineraries [get]
func (h *Handler) GetAllItineraries(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	request := pb.GetAllItinerariesRequest{
		ItineraryId: ctx.Query("name"),
		Limit:       int32(limit),
		Offset:      int32(offset),
	}

	resp, err := h.ContentService.GetAllItineraries(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in Gateway GetAllItineraries"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// ItinerariesFullInfoHandler retrieves full information about a specific itinerary.
// @Summary Get Itinerary Full Info
// @Description Retrieve full information about a specific itinerary.
// @Tags Itinerary
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param itinerary_id path string true "Itinerary ID"
// @Success 200 {object} genproto.ItinerariesFullInfoResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/itinerariesFullInfo/{itinerary_id} [get]
func (h *Handler) ItinerariesFullInfo(ctx *gin.Context) {
	itineraryID := ctx.Param("itinerary_id")

	_, err := uuid.Parse(itineraryID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("wrong itinerary_id"))
		h.Log.Error("error")
		return
	}

	resp, err := h.ContentService.ItinerariesFullInfo(ctx, &pb.ItinerariesFullInfoRequest{ItineraryId: itineraryID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in Gateway ItinerariesFullInfo"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// CommentItinerariesHandler handles commenting on an itinerary.
// @Summary Comment on Itinerary
// @Description Comment on a specific itinerary.
// @Tags Itinerary
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param itinerary_id path string true "Itinerary ID"
// @Param Create body genproto.CommentItinerariesRequest true "Comment on Itinerary"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/commentItineraries/{itinerary_id} [post]
func (h *Handler) CommentItineraries(ctx *gin.Context) {
	itineraryID := ctx.Param("itinerary_id")

	_, err := uuid.Parse(itineraryID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("wrong itinerary_id"))
		h.Log.Error("error")
		return
	}

	request := pb.CommentItinerariesRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		h.Log.Error("error")
		return
	}
	request.ItineraryId = itineraryID

	_, err = h.ContentService.CommentItineraries(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in Gateway CommentItineraries"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Comment added to Itinerary",
	})
}

// GetDestinationsHandler retrieves all destinations.
// @Summary Get All Destinations
// @Description Retrieve all destinations.
// @Tags Destination
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} genproto.GetDestinationsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/getDestinations [get]
func (h *Handler) GetDestinations(ctx *gin.Context) {
	resp, err := h.ContentService.GetDestinations(ctx, &pb.GetDestinationsRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in Gateway GetDestinations"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetDestinationInfoHandler retrieves information about a specific destination.
// @Summary Get Destination Info
// @Description Retrieve information about a specific destination.
// @Tags Destination
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param destination_id path string true "Destination ID"
// @Success 200 {object} genproto.GetDestinationInfoResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/getDestinationInfo/{destination_id} [get]
func (h *Handler) GetDestinationInfo(ctx *gin.Context) {
	destinationID := ctx.Param("destination_id")

	_, err := uuid.Parse(destinationID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("wrong destination_id"))
		h.Log.Error("error")
		return
	}

	resp, err := h.ContentService.GetDestinationInfo(ctx, &pb.GetDestinationInfoRequest{DestinationId: destinationID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in Gateway GetDestinationInfo"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// SentMessageHandler handles sending a message.
// @Summary Send Message
// @Description Send a message.
// @Tags Message
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body genproto.SentMessageRequest true "Send Message"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/sentMessage [post]
func (h *Handler) SentMessage(ctx *gin.Context) {
	request := pb.SentMessageRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		h.Log.Error("error")
		return
	}

	if request.Content == "" || request.RecipientId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "it is not full information"})
		return
	}

	_, err := h.ContentService.SentMessage(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error sending message in handler"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Message Sent",
	})
}

// GetAllMessagesHandler retrieves all messages with optional filtering and pagination.
// @Summary Get All Messages
// @Description Retrieve all messages with optional filtering and pagination.
// @Tags Message
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param limit query int false "Number of items to return"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} genproto.GetAllMessagesResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/getAllMessages [get]
func (h *Handler) GetAllMessages(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	request := pb.GetAllMessagesRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	resp, err := h.ContentService.GetAllMessages(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in Gateway GetAllMessages"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// CreateTravelTipHandler handles creating a travel tip.
// @Summary Create Travel Tip
// @Description Create a new travel tip.
// @Tags TravelTip
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body genproto.CreateTravelTipRequest true "Create Travel Tip"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/createTravelTip [post]
func (h *Handler) CreateTravelTip(ctx *gin.Context) {
	request := pb.CreateTravelTipRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		h.Log.Error("error")
		return
	}

	if request.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "it is not full information"})
		return
	}

	_, err := h.ContentService.CreateTravelTip(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error creating travel tip in handler"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Travel Tip Created",
	})
}

// GetTravelTipsHandler retrieves all travel tips with optional filtering and pagination.
// @Summary Get All Travel Tips
// @Description Retrieve all travel tips with optional filtering and pagination.
// @Tags TravelTip
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param limit query int false "Number of items to return"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} genproto.GetTravelTipsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/getTravelTips [get]
func (h *Handler) GetTravelTips(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	request := pb.GetTravelTipsRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	resp, err := h.ContentService.GetTravelTips(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in Gateway GetTravelTips"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// UserStatisticsHandler retrieves statistics about user activity.
// @Summary Get User Statistics
// @Description Retrieve statistics about user activity.
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} genproto.UserStatisticsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/userStatistics/{user_id} [get]
func (h *Handler) UserStatistics(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	_, err := uuid.Parse(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("wrong user_id"))
		h.Log.Error("error")
		return
	}

	resp, err := h.ContentService.UserStatistics(ctx, &pb.UserStatisticsRequest{UserId: userID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in Gateway UserStatistics"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetTrendingDestinationsHandler retrieves trending destinations.
// @Summary Get Trending Destinations
// @Description Retrieve trending destinations.
// @Tags Destination
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} genproto.GetTrendingDestinationsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /content/get_trending_destinations [get]
func (h *Handler) GetTrendingDestinations(ctx *gin.Context) {
	resp, err := h.ContentService.GetTrendingDestinations(ctx, &pb.GetTrendingDestinationsRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in Gateway GetTrendingDestinations"})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
