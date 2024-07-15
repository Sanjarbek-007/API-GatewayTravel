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
// @Router /api/content/createStory [post]
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
			"message":"errorncreating story in handler",
		})
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":"Story Created",
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
// @Router /api/menu/update/{id} [put]
func (h *Handler) UpdateStory(ctx *gin.Context) {
	request := pb.UpdateStoryRequest{}

	err := ctx.ShouldBindJSON(&request)
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		logger.Error("error in CreateStory")
		return
	}
	request.StoryId = ctx.Param("id")

	_, err = uuid.Parse(request.StoryId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("wrong id"))
		h.Log.Error("error")
		return
	}

	if request.Name == "" || request.Description == "" {
		BadRequest(ctx, fmt.Errorf("malumot toliq emas"))
		h.Log.Error("error")
		return
	}
	_, err = h.ReservationService.GetByIdRestaurant(ctx, &pb.IdRequest{Id: request.RestaurantId})
	if err != nil {
		h.Log.Error("error")
		BadRequest(ctx, fmt.Errorf("restaurant id yoq %v", err))
		return
	}
	_, err = h.ReservationService.GetByIdMenu(ctx, &pb.IdRequest{Id: request.RestaurantId})
	if err != nil {
		BadRequest(ctx, fmt.Errorf("menu id yoq"))
		return
	}
	_, err = h.ReservationService.GetByIdMenu(ctx, &pb.IdRequest{Id: request.Id})
	if err != nil {
		BadRequest(ctx, fmt.Errorf("error is ->bu id yoq database da yoq"))
		return
	}
	_, err = h.ReservationService.UpdateMenu(ctx, &request)

	if err != nil {
		fmt.Println("++++++++++++++")

		InternalServerError(ctx, err)
		h.Log.Error("error")
		return
	}

	OK(ctx)
}

// DeleteMenuHandler handles the deletion of a menu item.
// @Summary Delete Menu
// @Description Delete an existing menu item
// @Tags Menu
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Menu ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/menu/delete/{id} [delete]
func (h *Handler) DeleteMenuHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if Parse(id) {
		h.Log.Error("error")
		return
	}
	_, err := h.ReservationService.GetByIdMenu(ctx, &pb.IdRequest{Id: id})
	if err != nil {

		h.Log.Error("error")

		BadRequest(ctx, fmt.Errorf("error is  ->bu id yoq database da yoq"))
		return
	}
	_, err = h.ReservationService.GetByIdMenu(ctx, &pb.IdRequest{})
	if err != nil {
		BadRequest(ctx, fmt.Errorf("menu id yoq"))
		return
	}

	_, err = h.ReservationService.DeleteMenu(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		fmt.Println("++++++++++++", err)
		InternalServerError(ctx, err)
		h.Log.Error("error")
		return
	}
	h.Log.Info("ishladi")
	OK(ctx)
}

// GetByIdMenuHandler handles the request to fetch a menu item by its ID.
// @Summary Get Menu by ID
// @Description Get a menu item by its ID
// @Tags Menu
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Menu ID"
// @Success 200 {object} genproto.MenuResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/menu/get_id/{id} [get]
func (h *Handler) GetByIdMenuHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if Parse(id) {
		h.Log.Error("error")
		return
	}

	resp, err := h.ReservationService.GetByIdMenu(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		h.Log.Error("error")
		InternalServerError(ctx, fmt.Errorf(""))
		return
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// GetAllMenuHandler retrieves a list of menu items with optional filtering and pagination.
// @Summary Get All Menu
// @Description Retrieve a list of menu items with optional filtering and pagination.
// @Tags Menu
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param name query string false "Filter by menu item name"
// @Param description query string false "Filter by menu item description"
// @Param restaurant_id query string false "Filter by restaurant ID"
// @Param limit query int false "Number of items to return"
// @Param offset query int false "Offset for pagination"
// @Param price query string false "Filter by menu item price"
// @Success 200 {object} genproto.MenusResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/menu/get_all [get]
func (h *Handler) GetAllMenuHandler(ctx *gin.Context) {
	h.Log.Info("dsndjfjef")
	request := pb.GetAllMenuRequest{
		Name:         ctx.Query("name"),
		Description:  ctx.Query("description"),
		RestaurantId: ctx.Query("restaurant_id"),
		LimitOffset:  &pb.Filter{}, // Ensure LimitOffset is initialized
	}

	limit := ctx.Query("limit")
	limit1, err := IsLimitOffsetValidate(limit)
	if err != nil {
		BadRequest(ctx, err)
		h.Log.Error("error")
		return
	}

	if len(request.RestaurantId) > 0 {
		if Parse(request.RestaurantId) {
			BadRequest(ctx, fmt.Errorf("id hato"))
			h.Log.Error("error")
			return
		}
	}

	offset := ctx.Query("offset")
	offset1, err := IsLimitOffsetValidate(offset)
	if err != nil {
		h.Log.Error("error")
		BadRequest(ctx, err)
		return
	}

	price := ctx.Query("price")
	var price1 int
	if price != "" {
		price1, err = strconv.Atoi(price)
		if err != nil {
			h.Log.Error("error")
			BadRequest(ctx, err)
			return
		}
	}

	request.LimitOffset.Limit = int64(limit1)
	request.LimitOffset.Offset = int64(offset1)
	request.Price = float32(price1)

	if len(request.RestaurantId) != 0 {
		if Parse(request.RestaurantId) {
			BadRequest(ctx, fmt.Errorf("id hato"))
			h.Log.Error("error")
			return
		} else {
			_, err = h.ReservationService.GetByIdRestaurant(ctx, &pb.IdRequest{Id: request.RestaurantId})
			if err != nil {
				h.Log.Error("error")
				BadRequest(ctx, fmt.Errorf("restaurant id yoq"))
				return
			}
		}
	}

	resp, err := h.ReservationService.GetAllMenu(ctx, &request)
	if err != nil {
		fmt.Println("+++++++++", err)
		InternalServerError(ctx, err)
		h.Log.Error("error")
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
