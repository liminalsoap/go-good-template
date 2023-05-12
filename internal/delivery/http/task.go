package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo/internal/entitiy"
	"todo/internal/usecase"
	"todo/pkg/logger"
)

type taskRoutes struct {
	useCase usecase.Task
	log     logger.Interface
}

type errResponse struct {
	message string `json:"message"`
}

func newTaskRoutes(handler *gin.RouterGroup, log logger.Interface, uc usecase.Task) {
	r := &taskRoutes{uc, log}

	h := handler.Group("/task")
	{
		h.POST("/", r.CreateOne)
		h.PUT("/:id", r.Update)
		h.DELETE("/:id", r.Delete)

		h.GET("/:id", r.GetById)
	}
	handler.GET("/tasks", r.List)
}

type createInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// CreateOne godoc
//
//	@Summary		Create task
//	@Description	Create one task
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			Task	body		createInput	false	"Create Task input"
//	@Success		200		{object}	response
//	@failure		400		{object}	response
//	@Failure		500		{object}	response
//	@Router			/task [post]
func (r taskRoutes) CreateOne(c *gin.Context) {
	var input createInput
	if err := c.ShouldBindJSON(&input); err != nil {
		r.log.Errorf("err: %s", err)
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	task := entitiy.Task{
		Title:       input.Title,
		Description: input.Description,
	}
	if err := r.useCase.Create(c.Request.Context(), task); err != nil {
		r.log.Errorf("err: %s", err)
		errorResponse(c, http.StatusInternalServerError, "failed to create task")

		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// GetById godoc
//
//	@Summary		Get task
//	@Description	Get task by ID
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	false	"Task id"
//	@Success		200	{object}	entitiy.Task
//	@failure		400	{object}	response
//	@failure		404	{object}	response
//	@Router			/task/:id [get]
func (r taskRoutes) GetById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		r.log.Errorf("err: %s", err)
		errorResponse(c, http.StatusBadRequest, "invalid request param")

		return
	}

	task, err := r.useCase.GetById(c.Request.Context(), uint(id))
	if err != nil {
		r.log.Errorf("err: %s", err)
		errorResponse(c, http.StatusNotFound, "Task not found")

		return
	}
	c.JSON(http.StatusOK, task)
}

// Update godoc
//
//	@Summary		Update task
//	@Description	Update task by ID
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			Task	body		createInput	false	"Update Task input"
//	@Param			id		path		int			false	"Task id"
//	@Success		200		{object}	entitiy.Task
//	@failure		400		{object}	response
//	@failure		500		{object}	response
//	@Router			/task/:id [post]
func (r taskRoutes) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		r.log.Errorf("err: %s", err)
		errorResponse(c, http.StatusBadRequest, "invalid request param")

		return
	}

	var input createInput
	if err := c.ShouldBindJSON(&input); err != nil {
		r.log.Errorf("err: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind input"})
		return
	}
	fmt.Println(input)

	task := entitiy.Task{
		Title:       input.Title,
		Description: input.Description,
	}
	updatedTask, err := r.useCase.Update(c.Request.Context(), uint(id), task)
	if err != nil {
		r.log.Errorf("err: %s", err)
		errorResponse(c, http.StatusInternalServerError, "failed to update task")

		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

// Delete godoc
//
//	@Summary		Delete task
//	@Description	Delete task by ID
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	false	"Task id"
//	@Success		200	{object}	entitiy.Task
//	@failure		400	{object}	response
//	@failure		404	{object}	response
//	@Router			/task/:id [delete]
func (r taskRoutes) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		r.log.Errorf("err: %s", err)
		errorResponse(c, http.StatusBadRequest, "invalid request param")

		return
	}

	if err := r.useCase.Delete(c.Request.Context(), uint(id)); err != nil {
		r.log.Errorf("err: %s", err)
		errorResponse(c, http.StatusNotFound, err.Error())

		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// List godoc
//
//	@Summary		Get tasks
//	@Description	Get all tasks
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]entitiy.Task
//	@failure		500	{object}	response
//	@Router			/tasks [get]
func (r taskRoutes) List(c *gin.Context) {
	tasks, err := r.useCase.List(c.Request.Context())
	if err != nil {
		r.log.Errorf("err: %s", err)
		errorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, tasks)
}
