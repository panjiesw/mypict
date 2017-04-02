package model

import (
	"fmt"
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
	CreatedAt     *time.Time  `json:"created,omitempty"`
	UpdatedAt     *time.Time  `json:"updated,omitempty"`
}

func (i *Image) String() string {
	return fmt.Sprintf("Image<ID=%q Title=%q UserID=%q ContentPolicy=%d CreatedAt=%s UpdatedAt=%s>", i.ID, i.Title.String, i.UserID.String, i.ContentPolicy, i.CreatedAt, i.UpdatedAt)
}

// ImageDTO returned when getting an image
type ImageDTO struct {
	*Image
	//Title  string `json:"title,omitempty"`
	//UserID string `json:"uid,omitempty"`
	SID    string `json:"sid,omitempty"`
	GTitle string `json:"gtitle,omitempty"`
	GID    string `json:"gid,omitempty"`
}

func (i *ImageDTO) String() string {
	return fmt.Sprintf("ImageDTO<%s SID=%q GTitle=%q GID=%q>", i.Image, i.SID, i.GTitle, i.GID)
}
