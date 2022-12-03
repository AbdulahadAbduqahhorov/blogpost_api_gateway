package handlers

import (
	"net/http"
	"strconv"

	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/api_gateway/genproto/author_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/api_gateway/models"

	"github.com/gin-gonic/gin"
)

// CreateAuthor godoc
// @Summary     Create author
// @Description create author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       author body     models.CreateAuthorModel true "author body"
// @Success     201    {object} models.JSONResult{data=string} "Success"
// @Failure     400    {object} models.JSONErrorResult "Bad request"
// @Failure 	422    {object} models.JSONErrorResult{error=string} "Validation Error"
// @Failure     500    {object} models.JSONErrorResult "Server error"
// @Router      /v1/author [post]
func (h Handler) CreateAuthor(c *gin.Context) {
	var author models.CreateAuthorModel

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}

	authorId, err := h.services.Author.CreateAuthor(c.Request.Context(), &author_service.CreateAuthorRequest{
		FullName: author.FullName,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "Author has been created",
		Data:    authorId,
	})
}

// GetAuthorList godoc
// @Summary     List Author
// @Description get author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       limit  query    int    false "10"
// @Param       offset query    int    false "0"
// @Param       search query    string false "string default"
// @Success     200 {object} models.JSONResult{data=[]models.Author} "Success"
// @Failure     400    {object} models.JSONErrorResult "Bad request"
// @Router      /v1/author [get]
func (h Handler) GetAuthor(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")
	search := c.Query("search")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}
	res, err := h.services.Author.GetAuthor(c.Request.Context(), &author_service.GetAuthorRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
		Search: search,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "Author List",
		Data:    res,
	})
}

// GetAuthorById godoc
// @Summary     Get author by id
// @Description get an author by id
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       id  path     string true "author id"
// @Success     200 {object} models.JSONResult{data=models.Author} "Success"
// @Failure     404 {object} models.JSONErrorResult "Bad request"
// @Router      /v1/author/{id} [get]
func (h Handler) GetAuthorById(c *gin.Context) {
	id := c.Param("id")

	res, err := h.services.Author.GetAuthorById(c.Request.Context(), &author_service.GetAuthorByIdRequest{
		Id: id,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    res,
	})
}

// UpdateAuthor godoc
// @Summary     Update author
// @Description update author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       author body     models.UpdateAuthorModel true "author body"
// @Success     200    {object} models.JSONResult{data=models.Author} "Success"
// @Failure     400    {object} models.JSONErrorResult "Bad request"
// @Router      /v1/author [put]
func (h Handler) UpdateAuthor(c *gin.Context) {
	var author models.UpdateAuthorModel

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}

	status, err := h.services.Author.UpdateAuthor(c.Request.Context(), &author_service.UpdateAuthorRequest{
		Id:author.Id,
		FullName: author.FullName,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Data: status ,
	})
}

// DeleteAuthor godoc
// @Summary     Delete author
// @Description delete an author
// @Tags        authors
// @Produce     json
// @Param       id  path     string true "author id"
// @Success     200 {object} models.JSONResult{} "Success"
// @Failure     400 {object} models.JSONErrorResult "Bad request"
// @Router      /v1/author/{id} [delete]
func (h Handler) DeleteAuthor(c *gin.Context) {

	id := c.Param("id")
	status, err := h.services.Author.DeleteAuthor(c.Request.Context(), &author_service.DeleteAuthorRequest{
		Id:id,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Data: status,
	})
}
