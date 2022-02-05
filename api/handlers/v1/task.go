package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	pb "github.com/sulton0011/api-gateway/genproto/task"
	l "github.com/sulton0011/api-gateway/pkg/logger"
	"github.com/sulton0011/api-gateway/pkg/utils"
)

// TaskCreat creates task
// @Summary TaskCreat
// @Description This API for creating a new task
// @Tags task
// @Accept json
// @Produce json
// @Param Task reques body model.TaskCreat true "taskCreateRequest"
// @Success 200 {object} model.Task
// @Failure 400 {object} model.StandardErrorModel
// @Failure 500 {object} model.StandardErrorModel
// @Router /v1/tasks [post]
func (h *handlerV1) CreateTask(c *gin.Context) {
	var (
		body        pb.Task
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaskService().Create(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create Task", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetTask gets Task by id
// @Summary TaskGet
// @Description This API for getting task detail
// @Tags task
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} model.Task
// @Failure 400 {object} model.StandardErrorModel
// @Failure 500 {object} model.StandardErrorModel
// @Router /v1/tasks/{id} [get]
func (h *handlerV1) GetTask(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaskService().Get(
		ctx, &pb.IdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get Task", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListTasks returns list of Tasks
// @Summary TaskList
// @Description This API for getting list of task
// @Tags task
// @Accept json
// @Produce json
// @Param page query string false "Page"
// @Param limit query string false "limit"
// @Success 200 {object} model.ListResp
// @Failure 400 {object} model.StandardErrorModel
// @Failure 500 {object} model.StandardErrorModel
// @Router /v1/tasks/ [get]
func (h *handlerV1) ListTasks(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaskService().List(
		ctx, &pb.ListReq{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list Tasks", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateTask updates Task by id
// @Summary TastUpdate
// @Description This API for updating task
// @Tags task
// @Accept json
// @Produce json
// @Param Task request body model.TaskCreat true "taskUpdateRequest"
// @Success 200 {object} model.Task
// @Failure 400 {object} model.StandardErrorModel
// @Failure 500 {object} model.StandardErrorModel
// @Router /v1/tasks/{id} [put]
func (h *handlerV1) UpdateTask(c *gin.Context) {
	var (
		body        pb.Task
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaskService().Update(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update Task", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteTask deletes Task by id
// @Summary DeleteTask
// @Description This API for deleting task
// @Tags task
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200
// @Failure 400 {object} model.StandardErrorModel
// @Failure 500 {object} model.StandardErrorModel
// @Router /v1/tasks/{id} [delete]
func (h *handlerV1) DeleteTask(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaskService().Delete(
		ctx, &pb.IdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete Task", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
