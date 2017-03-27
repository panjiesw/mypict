package model

import (
	"time"

	"gopkg.in/nullbio/null.v6"
)

type Gallery struct {
	ID            string      `json:"id"`
	Title         null.String `json:"title"`
	UserID        null.String `json:"uid"`
	ContentPolicy uint8       `json:"cp"`
	CreatedAt     time.Time   `json:"created"`
	UpdatedAt     null.Time   `json:"updated"`
}

type GalleryS struct {
	Title         null.String `json:"title"`
	ContentPolicy uint8       `json:"cp"`
	Images        []ImageS    `json:"images"`
}

type GalleryR struct {
	ID            string      `json:"id"`
	Title         null.String `json:"title"`
	UserID        null.String `json:"uid"`
	ContentPolicy uint8       `json:"cp"`
	Images        []ImageSR   `json:"images"`
	CreatedAt     string      `json:"created"`
}
