package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type repositoryInterface interface {
	Db() *gorm.DB
	Create(url string, singleUse uint8) *Link
	Find(link *Link, key string) error
	Delete(link *Link)
	DeleteIfSingleUse(link *Link)
	CreateLog(linkId uint, userAgent, clientIp string)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r Repository) Db() *gorm.DB {
	return r.db
}

func (r Repository) Find(link *Link, key string) error {
	result := r.db.
		Select("id", "url", "single_use").
		Where("`key` = ?", key).
		Take(&link)

	return result.Error
}

func (r Repository) Create(url string, singleUse uint8) *Link {
	link := &Link{Url: url, Key: generateSecureToken(10), SingleUse: singleUse}
	r.db.Create(link)

	return link
}

func (r Repository) Delete(link *Link) {
	r.db.Delete(&link)
}

func (r Repository) DeleteIfSingleUse(link *Link) {
	if link.SingleUse == 1 {
		r.Delete(link)
	}
}

func (r Repository) CreateLog(linkId uint, userAgent, clientIp string) {
	r.db.Create(&LinkLog{
		LinkID:    linkId,
		UserAgent: userAgent,
		ClientIp:  clientIp,
	})
}

// TODO generate with SHA 256 (unix time + random string)
func generateSecureToken(length int) string {
	b := make([]byte, length)

	if _, err := rand.Read(b); err != nil {
		return ""
	}

	h := sha256.New()
	h.Write([]byte(time.Now().Format(time.RFC3339) + hex.EncodeToString(b)))

	return fmt.Sprintf("%x", h.Sum(nil))[0:8]
}
