package handler

import (
	"flower-shop/api/models"
	"flower-shop/config"
	"flower-shop/pkg/logger"
	"flower-shop/service"
	"flower-shop/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service.IServiceManager
	Log     logger.ILogger
}

func NewStrg(store storage.IStorage, service service.IServiceManager, log logger.ILogger) Handler {
	return Handler{
		Service: service,
		Log:     log,
	}
}

func handleResponse(c *gin.Context, log logger.ILogger, msg string, statusCode int, data interface{}) {
	resp := models.Response{}

	if statusCode >= 100 && statusCode <= 199 {
		resp.Description = config.ERR_INFORMATION
	} else if statusCode >= 200 && statusCode <= 299 {
		resp.Description = config.SUCCESS
		log.Info("REQUEST SUCCEEDED", logger.Any("msg: ", msg), logger.Int("status: ", statusCode))
	} else if statusCode >= 300 && statusCode <= 399 {
		resp.Description = config.ERR_REDIRECTION
	} else if statusCode >= 400 && statusCode <= 499 {
		resp.Description = config.ERR_BADREQUEST
		log.Error("BAD REQUEST", logger.Any("error: ", msg), logger.Any("data: ", data))
	} else {
		resp.Description = config.ERR_INTERNAL_SERVER
		log.Error("ERR_INTERNAL_SERVER", logger.Any("error: ", msg))
	}

	resp.StatusCode = statusCode
	resp.Data = data

	c.JSON(resp.StatusCode, resp)
}

func ParsePageQueryParam(c *gin.Context) (uint64, error) {
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.ParseUint(pageStr, 10, 30)
	if err != nil {
		return 0, err
	}
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

func ParseLimitQueryParam(c *gin.Context) (uint64, error) {
	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.ParseUint(limitStr, 10, 30)
	if err != nil {
		return 0, err
	}

	if limit == 0 {
		return 10, nil
	}
	return limit, nil
}

func handleResponseGet(c *gin.Context, log logger.ILogger, statusCode int, data interface{}) {
	if statusCode >= 200 && statusCode <= 299 {
		log.Info("REQUEST SUCCEEDED", logger.Int("status: ", statusCode))
	} else if statusCode >= 400 && statusCode <= 499 {
		log.Error("BAD REQUEST", logger.Any("data: ", data))
	} else {
		log.Error("ERR_INTERNAL_SERVER")
	}
	c.JSON(statusCode, data)
}
