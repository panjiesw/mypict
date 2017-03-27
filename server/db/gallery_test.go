package db_test

import (
	"reflect"
	"testing"

	"gopkg.in/nullbio/null.v6"
	"panjiesw.com/mypict/server/model"
)

func TestDatabase_GalleryByID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.GalleryR
		wantErr bool
	}{
		{
			name: "ID LkwzNLWkg",
			args: args{id: "LkwzNLWkg"},
			want: &model.GalleryR{
				ID:            "LkwzNLWkg",
				Title:         null.StringFrom("Volutpat Quam Pede"),
				UserID:        null.StringFrom("XpJrI2neRED"),
				ContentPolicy: 1,
				Images: []model.ImageSR{
					{
						ID:    "Sw9rILWzgX",
						SID:   null.StringFrom("sPL_CSMzR2"),
						Title: null.StringFrom("Suscipit Nulla Elit"),
					},
					{
						ID:    "IQ99SYWkR8",
						SID:   null.StringFrom("sELljSMkgR"),
						Title: null.StringFrom("Diam Cras Pellentesque Volutpat Dui"),
					},
					{
						ID:    "IQ9rSYWzR1",
						SID:   null.StringFrom("yPL_jSMzR_"),
						Title: null.StringFrom("Nibh In Quis Justo"),
					},
					{
						ID:    "IQr9SLZkRO",
						SID:   null.StringFrom("sEL_jIGkRY"),
						Title: null.StringFrom("Rhoncus Dui Vel"),
					},
					{
						ID:    "IQrrILWkgZ",
						SID:   null.StringFrom("sELljIGkR1"),
						Title: null.StringFrom("Turpis Donec Posuere Metus"),
					},
					{
						ID:    "4QrrSYWkgh",
						SID:   null.StringFrom("yEL_CSGkgI"),
						Title: null.StringFrom("Pulvinar Nulla Pede"),
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got model.GalleryR
			err := d.GalleryByID(tt.args.id, &got)
			got.CreatedAt = ""
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.GalleryByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(&got, tt.want) {
				t.Errorf("Database.GalleryByID() = %v, want %v", &got, tt.want)
			}
		})
	}
}

func TestDatabase_GallerySave(t *testing.T) {
	type args struct {
		g   model.GalleryS
		uid null.String
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "3 no uid no error",
			args: args{
				g: model.GalleryS{
					Title:         null.StringFrom("gallery1"),
					ContentPolicy: 1,
					Images: []model.ImageS{
						{ID: "image1", Title: null.StringFrom("image1"), ContentPolicy: 0},
						{ID: "image2", Title: null.NewString("", false), ContentPolicy: 1},
						{ID: "image3", Title: null.StringFrom("image3"), ContentPolicy: 1},
					},
				},
				uid: null.NewString("", false),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := d.GallerySave(tt.args.g, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.GallerySave() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.ID == "" {
				t.Error("Database.GallerySave() gid not generated")
				return
			}

			for _, img := range got.Images {
				if img.ID == "" {
					t.Error("Database.GallerySave() iid not generated")
					break
				}
				if !img.SID.Valid {
					t.Error("Database.GallerySave() sid not generated")
					break
				}
			}
		})
	}
}
