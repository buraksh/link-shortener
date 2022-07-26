package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LinkLog struct {
	gorm.Model
	LinkID    uint   `gorm:"not null"`
	UserAgent string `gorm:"not null"`
	ClientIp  string `gorm:"not null"`
}

type Link struct {
	gorm.Model
	Key       string `gorm:"not null;index:unique;size:12"`
	Url       string `gorm:"not null" json:"url" binding:"required,uri"`
	SingleUse uint8  `gorm:"not null;default:0" json:"single_use"`
	LinkLogs  []LinkLog
}

func main() {
	cfg := Config{}
	cfg.Load()

	r := gin.Default()

	ctrl := NewController()

	r.GET("/ping", ctrl.Ping)
	r.GET("/:key", ctrl.Redirect)
	r.POST("/links", ctrl.Create)

	_ = r.Run(cfg.GetPort())
}
