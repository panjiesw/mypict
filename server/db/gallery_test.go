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
		want    *model.GalleryDTO
		wantErr bool
	}{
		{
			name: "ID LkwzNLWkg",
			args: args{id: "LkwzNLWkg"},
			want: &model.GalleryDTO{
				Gallery: &model.Gallery{
					ID:            "LkwzNLWkg",
					ContentPolicy: 1,
					Title:         null.StringFrom("Volutpat Quam Pede"),
					UserID:        null.StringFrom("XpJrI2neRED"),
				},
				Images: []*model.ImageDTO{
					{
						Image: &model.Image{
							ID:            "Sw9rILWzgX",
							ContentPolicy: 1,
							UserID:        null.StringFrom("XpJrI2neRED"),
							Title:         null.StringFrom("Suscipit Nulla Elit"),
						},
						GID: "LkwzNLWkg",
						SID: "sPL_CSMzR2",
					},
					{
						Image: &model.Image{
							ID:            "IQ99SYWkR8",
							ContentPolicy: 1,
							UserID:        null.StringFrom("XpJrI2neRED"),
							Title:         null.StringFrom("Diam Cras Pellentesque Volutpat Dui"),
						},
						GID: "LkwzNLWkg",
						SID: "sELljSMkgR",
					},
					{
						Image: &model.Image{
							ID:            "IQ9rSYWzR1",
							ContentPolicy: 1,
							UserID:        null.StringFrom("XpJrI2neRED"),
							Title:         null.StringFrom("Nibh In Quis Justo"),
						},
						GID: "LkwzNLWkg",
						SID: "yPL_jSMzR_",
					},
					{
						Image: &model.Image{
							ID:            "IQr9SLZkRO",
							ContentPolicy: 1,
							UserID:        null.StringFrom("XpJrI2neRED"),
							Title:         null.StringFrom("Rhoncus Dui Vel"),
						},
						GID: "LkwzNLWkg",
						SID: "sEL_jIGkRY",
					},
					{
						Image: &model.Image{
							ID:     "IQrrILWkgZ",
							UserID: null.StringFrom("XpJrI2neRED"),
							Title:  null.StringFrom("Turpis Donec Posuere Metus"),
						},
						GID: "LkwzNLWkg",
						SID: "sELljIGkR1",
					},
					{
						Image: &model.Image{
							ID:            "4QrrSYWkgh",
							ContentPolicy: 1,
							UserID:        null.StringFrom("XpJrI2neRED"),
							Title:         null.StringFrom("Pulvinar Nulla Pede"),
						},
						GID: "LkwzNLWkg",
						SID: "yEL_CSGkgI",
					},
				},
			},
			wantErr: false,
		},
		{
			name:    "Not found",
			args:    args{id: "somefoo"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got model.GalleryDTO
			err := d.GalleryByID(tt.args.id, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.GalleryByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			got.Gallery.CreatedAt = nil
			for _, img := range got.Images {
				if img.Image.ID == "" {
					t.Errorf("Database.GalleryByID() id not generated = %s", img.Image.Title)
				}
				img.Image.CreatedAt = nil
			}
			if !reflect.DeepEqual(&got, tt.want) {
				t.Errorf("Database.GalleryByID() = \n%v, want \n%v", &got, tt.want)
			}
		})
	}
}

func TestDatabase_GallerySave(t *testing.T) {
	type args struct {
		g *model.GalleryDTO
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "3 no uid no error",
			args: args{
				g: &model.GalleryDTO{
					Gallery: &model.Gallery{
						Title:         null.StringFrom("gallery1"),
						ContentPolicy: 1,
					},
					Images: []*model.ImageDTO{
						{
							Image: &model.Image{
								ID:            "image1",
								Title:         null.StringFrom("image1"),
								ContentPolicy: 0,
							},
						},
						{
							Image: &model.Image{
								ID:            "image2",
								Title:         null.NewString("", false),
								ContentPolicy: 1,
							},
						},
						{
							Image: &model.Image{
								ID:            "image3",
								Title:         null.StringFrom("image3"),
								ContentPolicy: 1,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := d.GallerySave(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.GallerySave() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if tt.args.g.ID == "" {
				t.Error("Database.GallerySave() gid not generated")
				return
			}

			for _, img := range tt.args.g.Images {
				if img.Image.ID == "" {
					t.Error("Database.GallerySave() iid not generated")
					break
				}
				if img.SID == "" {
					t.Error("Database.GallerySave() sid not generated")
					break
				}
			}
		})
	}
}
