package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	repository repositoryInterface
}

func NewController() *Controller {
	db := ConnectToDb()

	return &Controller{repository: NewRepository(db)}
}

func (ctrl Controller) Redirect(c *gin.Context) {
	link := &Link{}

	if ctrl.repository.Find(link, c.Param("key")) != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Link not found",
		})
		return
	}

	ctrl.repository.CreateLog(link.ID, c.GetHeader("user-agent"), c.ClientIP())

	ctrl.repository.DeleteIfSingleUse(link)

	c.Redirect(http.StatusMovedPermanently, link.Url)
}

func (ctrl Controller) Create(c *gin.Context) {
	request := Link{}

	if !ValidateRequest(c, &request) {
		return
	}

	link := ctrl.repository.Create(request.Url, request.SingleUse)

	c.JSON(http.StatusOK, gin.H{
		"key": link.Key,
		"url": link.Url,
	})
}

func (ctrl Controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
