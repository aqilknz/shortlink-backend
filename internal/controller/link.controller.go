package controller

import (
	"net/http"
	"strconv"

	"github.com/aqilknz/shortlink-backend/internal/dto"
	"github.com/aqilknz/shortlink-backend/internal/service"
	"github.com/aqilknz/shortlink-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type LinkController struct {
	linkService service.LinkService
}

func NewLinkController(linkService service.LinkService) *LinkController {
	return &LinkController{linkService: linkService}
}

// CreateLink
// @Summary      Create a new short link
// @Description  Creates a short URL from a long URL. Users can optionally provide a 'custom_slug' (3-50 characters, alphanumeric, and hyphens). If left empty, the system will auto-generate a random 6-character slug.
// @Tags         Links
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        request body dto.CreateLinkRequest true "URL Payload. Example: {'original_url': 'https://example.com', 'custom_slug': 'my-promo'}"
// @Success      201  {object}  utils.BaseResponse{results=dto.LinkResponse} "Success example: {'status': 201, 'message': 'Short link created successfully', 'results': {'id': 1, 'original_url': 'https://example.com', 'slug': 'my-promo'}}"
// @Failure      400  {object}  utils.BaseResponse "Error example: Validation failed or Invalid slug format"
// @Failure      401  {object}  utils.BaseResponse "Error example: Unauthorized (Invalid or missing token)"
// @Failure      409  {object}  utils.BaseResponse "Error example: Slug already exists"
// @Failure      500  {object}  utils.BaseResponse
// @Router       /links [post]
func (ctrl *LinkController) CreateLink(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.ErrorResponse(c, utils.ErrUnauthorized.Code, utils.ErrUnauthorized.Message)
		return
	}

	var req dto.CreateLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, utils.ErrInvalidInput.Code, utils.ErrInvalidInput.Message)
		return
	}

	uid := userID.(int)
	res, err := ctrl.linkService.CreateLink(c.Request.Context(), uid, req)
	if err != nil {
		if customErr, ok := err.(*utils.AppError); ok {
			utils.ErrorResponse(c, customErr.Code, customErr.Message)
			return
		}
		utils.ErrorResponse(c, utils.ErrInternalServer.Code, utils.ErrInternalServer.Message)
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Short link created successfully", res)
}

// GetUserLinks
// @Summary      Get all user links (Paginated & Searchable)
// @Description  Retrieves a list of all short links owned by the logged-in user. Supports pagination and searching by original URL or slug.
// @Tags         Links
// @Produce      json
// @Security     ApiKeyAuth
// @Param        page   query int    false "Page number (default: 1)" default(1)
// @Param        limit  query int    false "Items per page (default: 10)" default(10)
// @Param        search query string false "Search keyword for original_url or slug"
// @Success      200  {object}  utils.BaseResponse{results=dto.PaginatedLinkResponse} "Success example with pagination metadata"
// @Failure      401  {object}  utils.BaseResponse "Error example: Unauthorized"
// @Failure      500  {object}  utils.BaseResponse
// @Router       /links [get]
func (ctrl *LinkController) GetUserLinks(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.ErrorResponse(c, utils.ErrUnauthorized.Code, utils.ErrUnauthorized.Message)
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	search := c.Query("search")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	uid := userID.(int)
	res, err := ctrl.linkService.GetUserLinks(c.Request.Context(), uid, search, page, limit)
	if err != nil {
		utils.ErrorResponse(c, utils.ErrInternalServer.Code, utils.ErrInternalServer.Message)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Links retrieved successfully", res)
}

// DeleteLink
// @Summary      Delete a short link
// @Description  Deletes a short link by its ID. This endpoint uses a 'Soft Delete' approach (the record remains in the database but is marked as deleted and becomes inaccessible).
// @Tags         Links
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id   path      int  true  "ID of the link to be deleted. Example: 5"
// @Success      200  {object}  utils.BaseResponse "Success example: {'status': 200, 'message': 'Link deleted successfully', 'results': null}"
// @Failure      400  {object}  utils.BaseResponse "Error example: Invalid link ID format"
// @Failure      401  {object}  utils.BaseResponse
// @Failure      403  {object}  utils.BaseResponse "Error example: Forbidden (Not the owner of the link)"
// @Failure      500  {object}  utils.BaseResponse
func (ctrl *LinkController) DeleteLink(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.ErrorResponse(c, utils.ErrUnauthorized.Code, utils.ErrUnauthorized.Message)
		return
	}

	linkIDStr := c.Param("id")
	linkID, err := strconv.Atoi(linkIDStr)
	if err != nil {
		utils.ErrorResponse(c, utils.ErrInvalidInput.Code, "Invalid link ID format")
		return
	}

	uid := userID.(int)
	err = ctrl.linkService.DeleteLink(c.Request.Context(), uid, linkID)
	if err != nil {
		if customErr, ok := err.(*utils.AppError); ok {
			utils.ErrorResponse(c, customErr.Code, customErr.Message)
			return
		}
		utils.ErrorResponse(c, utils.ErrInternalServer.Code, utils.ErrInternalServer.Message)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Link deleted successfully", nil)
}

// CheckSlugAvailability
// @Summary      Check custom slug availability
// @Description  Checks in real-time whether a custom slug is already in use or available. Highly useful for frontend live form validation.
// @Tags         Links
// @Produce      json
// @Security     ApiKeyAuth
// @Param        slug query string true "The slug to check. Example: /links/check-slug?slug=my-promo"
// @Success      200  {object}  utils.BaseResponse "Success example: {'results': {'available': true, 'slug': 'my-promo'}}"
// @Failure      400  {object}  utils.BaseResponse "Error example: Slug parameter is required"
// @Failure      401  {object}  utils.BaseResponse
// @Router       /links/check-slug [get]
func (ctrl *LinkController) CheckSlugAvailability(c *gin.Context) {
	_, exists := c.Get("userID")
	if !exists {
		utils.ErrorResponse(c, utils.ErrUnauthorized.Code, utils.ErrUnauthorized.Message)
		return
	}

	slug := c.Query("slug")
	if slug == "" {
		utils.ErrorResponse(c, utils.ErrInvalidInput.Code, "Slug parameter is required")
		return
	}

	isAvailable, err := ctrl.linkService.CheckSlugAvailability(c.Request.Context(), slug)
	if err != nil {
		utils.ErrorResponse(c, utils.ErrInternalServer.Code, utils.ErrInternalServer.Message)
		return
	}

	data := map[string]interface{}{
		"slug":      slug,
		"available": isAvailable,
	}

	message := "Slug is available"
	if !isAvailable {
		message = "Slug is already in use"
	}

	utils.SuccessResponse(c, http.StatusOK, message, data)
}

// Redirect godoc
// @Summary      Redirect to Original URL
// @Description  A public endpoint that performs an HTTP 301 Permanent Redirect from the short slug to its original URL. Returns a 404 error if the slug is invalid or has been deleted.
// @Tags         Redirect
// @Produce      json
// @Param        slug path string true "The short slug of the URL. Example: aB3x9Z"
// @Success      301  "Redirects the client to the Original URL"
// @Failure      404  {object} utils.BaseResponse "Error example: Short link not found or has been deleted"
// @Router       /{slug} [get]
func (ctrl *LinkController) Redirect(c *gin.Context) {
	slug := c.Param("slug")

	originalURL, err := ctrl.linkService.GetOriginalURL(c.Request.Context(), slug)
	if err != nil {
		utils.ErrorResponse(c, utils.ErrNotFound.Code, "Short link not found or has been deleted")
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}
