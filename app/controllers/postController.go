package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Rayato159/isekai-shop-api/app/models"
	"github.com/labstack/echo/v4"
)

type postErr string

const (
	findOnePostErr postErr = "post-001"
)

type IPostControllers interface {
	FindOnePost(c echo.Context) error
}

type postControllers struct{}

func NewPostControllers() IPostControllers {
	return &postControllers{}
}

func (h *postControllers) FindOnePost(c echo.Context) error {
	posts := map[int]*models.Post{
		1: {
			Id:      1,
			Title:   "Hello Svelte",
			Content: "This content is just a stupid content using Go Echo + Svelte.",
		},
	}
	postIdInt, err := strconv.Atoi(strings.Trim(c.Param("post_id"), " "))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.Error{
			TraceId: string(findOnePostErr),
			Message: err.Error(),
		})
	}

	if posts[postIdInt] != nil {
		return c.JSON(http.StatusOK, posts[postIdInt])
	}
	return c.JSON(http.StatusBadRequest, &models.Error{
		TraceId: string(findOnePostErr),
		Message: fmt.Sprintf("post_id: %s not found", strings.Trim(c.Param("post_id"), " ")),
	})
}
