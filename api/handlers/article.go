package handlers

import (
	"net/http"

	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/api_gateway/genproto/article_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/api_gateway/models"
	"github.com/gin-gonic/gin"
)

// CreateArticle godoc
// @Summary     Create article
// @Description create article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     models.CreateArticleModel true "article body"
// @Success     201     {object} models.JSONResult{data=string} "Success"
// @Failure     400     {object} models.JSONErrorResult "Bad request"
// @Failure     500     {object} models.JSONErrorResult "Server error"
// @Router      /v1/article [post]
func (h Handler) CreateArticle(c *gin.Context) {

	var body models.CreateArticleModel

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}
	articleId, err := h.services.Article.CreateArticle(c.Request.Context(), &article_service.CreateArticleRequest{
		Content: &article_service.Content{
			Title: body.Title,
			Body:  body.Body,
		},
		AuthorId: body.AuthorId,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "Article has been created",
		Data:    articleId,
	})
}

// GetArticleList godoc
// @Summary     List Article
// @Description get article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       limit  query    int    false "10"
// @Param       offset query    int    false "0"
// @Param       search query    string false "string default"
// @Success     200    {object} models.JSONResult{data=[]models.Article} "Success"
// @Failure     400     {object} models.JSONErrorResult "Bad request"
// @Router      /v1/article [get]
func (h Handler) GetArticle(c *gin.Context) {
	search := c.Query("search")

	limit, err := h.getLimitParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}

	offset, err := h.getOffsetParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}
	res, err := h.services.Article.GetArticleList(c.Request.Context(), &article_service.GetArticleListRequest{
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
		Message: "Article List",
		Data:    res,
	})
}

// GetArticleById godoc
// @Summary     Get article by id
// @Description get an article by id
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       id  path     string true "article id"
// @Success     200 {object} models.JSONResult{data=models.GetArticleByIdModel} "Success"
// @Failure     400 {object} models.JSONErrorResult "Bad request"
// @Router      /v1/article/{id} [get]
func (h Handler) GetArticleById(c *gin.Context) {
	id := c.Param("id")

	res, err := h.services.Article.GetArticleById(c.Request.Context(), &article_service.GetArticleByIdRequest{
		Id: id,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    res,
	})
}

// UpdateArticle godoc
// @Summary     Update article
// @Description update article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     models.UpdateArticleModel true "article body"
// @Success     200     {object} models.JSONResult{message=string} "Success"
// @Failure     400     {object} models.JSONErrorResult "Bad request"
// @Failure     500     {object} models.JSONErrorResult "Server error"
// @Router      /v1/article [put]
func (h Handler) UpdateArticle(c *gin.Context) {
	var body models.UpdateArticleModel

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}

	updatedArticleId, err := h.services.Article.UpdateArticle(c.Request.Context(), &article_service.UpdateArticleRequest{
		Id: body.Id,
		Content: &article_service.Content{
			Title: body.Title,
			Body:  body.Body,
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "Article has been  Updated",
		Data: updatedArticleId,
	})

}

// DeleteArticle godoc
// @Summary     Delete article
// @Description delete an article
// @Tags        articles
// @Produce     json
// @Param       id  path     string true "article id"
// @Success     200 {object} models.JSONResult{message=string} "Success"
// @Failure     400 {object} models.JSONErrorResult "Bad Request"
// @Router      /v1/article/{id} [delete]
func (h Handler) DeleteArticle(c *gin.Context) {

	id := c.Param("id")
	deletedArticleid,err := h.services.Article.DeleteArticle(c.Request.Context(),&article_service.DeleteArticleRequest{
		Id:id,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.JSONResult{
		Message: "Article has been Deleted",
		Data: deletedArticleid,
	})
}
