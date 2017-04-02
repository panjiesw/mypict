package model

import (
	"fmt"
	"time"

	"gopkg.in/nullbio/null.v6"
)

type Gallery struct {
	ID            string      `json:"id"`
	Title         null.String `json:"title"`
	UserID        null.String `json:"uid"`
	ContentPolicy uint8       `json:"cp"`
	CreatedAt     *time.Time  `json:"created,omitempty"`
	UpdatedAt     *time.Time  `json:"updated,omitempty"`
}

func (g *Gallery) String() string {
	return fmt.Sprintf("Gallery<ID=%q Title=%q UserID=%q ContentPolicy=%d CreatedAt=%s UpdatedAt=%s>", g.ID, g.Title.String, g.UserID.String, g.ContentPolicy, g.CreatedAt, g.UpdatedAt)
}

type GalleryDTO struct {
	*Gallery
	Images []*ImageDTO `json:"images"`
}

func (g *GalleryDTO) String() string {
	return fmt.Sprintf("GalleryDTO<%s Images=%s>", g.Gallery, g.Images)
}
