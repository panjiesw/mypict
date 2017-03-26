package model

import (
	"time"

	"github.com/jackc/pgx"
	"gopkg.in/nullbio/null.v6"
)

var (
	ImageTable    = "image"
	ImageIDTable  = "image_ids"
	ImageTableI   = pgx.Identifier{ImageTable}
	ImageIDTableI = pgx.Identifier{ImageIDTable}
)

// Image represents image data
type Image struct {
	ID            string      `json:"id"`
	Title         null.String `json:"title"`
	UserID        null.String `json:"uid"`
	ContentPolicy uint8       `json:"cp"`
	CreatedAt     time.Time   `json:"created"`
	UpdatedAt     null.Time   `json:"updated"`
}

// ImageS is the model for creating new image
type ImageS struct {
	ID            string      `json:"id"`
	Title         null.String `json:"title"`
	ContentPolicy uint8       `json:"cp"`
}

// ImageSR returned after saving the image
type ImageSR struct {
	ID    string      `json:"id"`
	Title null.String `json:"title"`
	SID   null.String `json:"sid"`
}

// ImageR returned when getting an image
type ImageR struct {
	ID            string      `json:"id"`
	Title         null.String `json:"title"`
	UserID        null.String `json:"uid"`
	SID           null.String `json:"sid"`
	GalleryTitle  null.String `json:"gtitle"`
	GalleryID     null.String `json:"gid"`
	ContentPolicy uint8       `json:"cp"`
	CreatedAt     string      `json:"created"`
}
